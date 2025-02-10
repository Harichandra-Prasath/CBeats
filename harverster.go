package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Logs struct {
	data *[]byte
}

type HarvesterConfig struct {
	SnkPort  string
	ReadTime int
	ReadDir  string
}

type Harvester struct {
	Dumper    *Dumper
	ReaderMap map[string]*Reader
	Listener  *Watcher
	EventChan chan string
	DumpChan  chan *Logs
}

func NewHarvester(cfg HarvesterConfig) (*Harvester, error) {

	// Check the validity of the dir
	files, err := os.ReadDir(cfg.ReadDir)
	if err != nil {
		return nil, fmt.Errorf("creating harvester: %s", err)

	}

	harvester := Harvester{}
	harvester.EventChan = make(chan string)
	harvester.DumpChan = make(chan *Logs)
	harvester.ReaderMap = make(map[string]*Reader)

	for _, entry := range files {
		if !entry.IsDir() {

			// Look for the log files
			parts := strings.Split(entry.Name(), ".")
			if parts[len(parts)-1] == "log" {

				reader, err := NewReader(cfg.ReadDir+entry.Name(), cfg.ReadTime, harvester.DumpChan)
				if err != nil {
					log.Println("Error creating Reader for " + entry.Name() + ": " + err.Error())
				}

				harvester.ReaderMap[entry.Name()] = reader
			}
		}
	}

	harvester.Listener, err = NewWatcher(cfg.ReadDir, harvester.EventChan)
	if err != nil {
		return nil, fmt.Errorf("creating harvester: %s", err)

	}

	harvester.Dumper, err = NewDumper(cfg.SnkPort)
	if err != nil {
		return nil, fmt.Errorf("creating harvester: %s", err)

	}

	return &harvester, nil
}

func (h *Harvester) Start() {

	// Run the Reader and Watcher
	go h.Listener.Listen()
	for _, reader := range h.ReaderMap {
		go reader.Read()
	}

	for {

		select {

		case f := <-h.EventChan:
			reader, ok := h.ReaderMap[f]
			if ok && !reader.Active {
				reader.NotifyChan <- struct{}{}
			}

		case logs := <-h.DumpChan:
			err := h.Dumper.dumpLogs(logs)
			if err != nil {
				log.Printf("Error in dumping logs: %s", err)
			}

		}

	}

}
