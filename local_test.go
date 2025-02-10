package main

import (
	"log"
	"os"
	"testing"
	"time"
	// "time"
)

func TestDumper(t *testing.T) {

	dumper, err := NewDumper("8989")
	if err != nil {
		t.Fatal(err)
	}

	sample_data := []byte("Foo Completed\nBar Completed\n")
	sample_logs := Logs{
		data: &sample_data,
	}

	dumper.dumpLogs(&sample_logs)
	dumper.TcpConn.Close()
}

func TestWatcher(t *testing.T) {

	tmp := t.TempDir()

	sample_fp := tmp + "/sample_logs.log"

	// Make a Initial Write
	fp, err := os.OpenFile(sample_fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)

	}

	teststrChan := make(chan string)
	watcher, err := NewWatcher(tmp, teststrChan)
	if err != nil {
		t.Fatal(err)

	}

	go watcher.Listen()
	go func(tchan chan string) {
		for {
			select {

			case event := <-tchan:
				log.Print("Event Recieved for " + event)
			}

		}

	}(teststrChan)
	for i := 0; i < 5; i++ {
		_, err := fp.WriteString("New Write\n")
		if err != nil {
			log.Fatal(err)

		}
		time.Sleep(2 * time.Second)
	}

	fp.Close()

}

func TestReader(t *testing.T) {

	tmp := t.TempDir()

	sample_fp := tmp + "/sample_logs.log"

	fp, err := os.OpenFile(sample_fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		_, err := fp.WriteString("New Write\n")
		if err != nil {
			log.Fatal(err)

		}
	}
	fp.Close()

	testLogChan := make(chan *Logs)

	reader, err := NewReader(sample_fp, 5, testLogChan)
	if err != nil {
		t.Fatal(err)
	}

	go reader.Read()
	go func(tchan chan *Logs) {
		for {
			select {

			case logs := <-tchan:
				log.Print(string(*logs.data))
			}

		}

	}(testLogChan)
	reader.NotifyChan <- struct{}{}
	time.Sleep(6 * time.Second)
}

func TestHarvester(t *testing.T) {
	tmp := t.TempDir()

	sample_fp := tmp + "/sample_logs.log"

	// Make a Initial Write
	fp, err := os.OpenFile(sample_fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)

	}

	harvester, err := NewHarvester(HarvesterConfig{
		ReadDir:  tmp + "/",
		ReadTime: 5,
		SnkPort:  "8989",
	})

	if err != nil {
		t.Fatal(err)

	}

	go harvester.Start()
	for i := 0; i < 5; i++ {
		_, err := fp.WriteString("New Write\n")
		if err != nil {
			log.Fatal(err)

		}
		time.Sleep(2 * time.Second)
	}

	fp.Close()

	for {
	}
}
