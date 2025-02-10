package main

import (
	"log"
	"os"
	"testing"
	"time"
	// "time"
)

func TestDumper(t *testing.T) {

	dumper, err := NewDumper()
	if err != nil {
		t.Fatal(err)
	}

	sample_data := []byte("Foo Completed\nBar Completed\n")
	sample_logs := Logs{
		batch: 0,
		data:  &sample_data,
	}

	dumper.dumpLogs(&sample_logs)

}

// func TestWatcher(t *testing.T) {
//
// 	tmp := t.TempDir()
//
// 	sample_fp := tmp + "/sample_logs.log"
//
// 	// Make a Initial Write
// 	fp, err := os.OpenFile(sample_fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// 	if err != nil {
// 		t.Fatal(err)
//
// 	}
//
// 	watcher, err := NewWatcher(WatcherConfig{Dir: tmp})
// 	if err != nil {
// 		t.Fatal(err)
//
// 	}
//
// 	go watcher.Listen()
// 	for i := 0; i < 5; i++ {
// 		_, err := fp.WriteString("New Write\n")
// 		if err != nil {
// 			log.Fatal(err)
//
// 		}
// 		time.Sleep(2 * time.Second)
// 	}
//
// 	fp.Close()
//
// }

func TestReader(t *testing.T) {

	tmp := t.TempDir()

	sample_fp := tmp + "/sample_logs.log"

	fp, err := os.OpenFile(sample_fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		_, err := fp.WriteString("New Write\n")
		if err != nil {
			log.Fatal(err)

		}
	}
	fp.Close()

	reader, err := NewReader(sample_fp)
	if err != nil {
		log.Fatal(err)
	}

	go reader.Read()
	reader.NotifyChan <- struct{}{}
	time.Sleep(6 * time.Second)
	fp, err = os.OpenFile(sample_fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		_, err := fp.WriteString("New New Write\n")
		if err != nil {
			log.Fatal(err)

		}
	}
	fp.Close()
	reader.NotifyChan <- struct{}{}
	time.Sleep(6 * time.Second)
}
