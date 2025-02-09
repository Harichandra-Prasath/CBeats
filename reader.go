package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Reader struct {
	Offset     int64
	File       *os.File
	NotifyChan chan struct{}
	Active     bool
}

func NewReader(fp string) (*Reader, error) {

	f, err := os.OpenFile(fp, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("creating reader: %s", err)

	}

	return &Reader{Offset: 0, File: f, NotifyChan: make(chan struct{})}, nil
}

func (r *Reader) Read() {
	for {

		select {

		case <-r.NotifyChan:
			// Wait for t seconds and then read all the logs
			// TODO: Maybe done better

			time.Sleep(READ_TIME * time.Second)
			_, err := r.File.Seek(r.Offset, io.SeekStart)
			if err != nil {
				log.Printf("Error in Reading %s", err)
			}

			logs, err := io.ReadAll(r.File)
			if err != nil {
				log.Printf("Error in Reading %s", err)
			}

			r.Offset, _ = r.File.Seek(0, 1)
			log.Printf("Extracted Logs\n%s", string(logs))

		}

	}

}
