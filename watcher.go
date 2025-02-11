// File Watcher that listens for writes on a particular file

package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	FSWatcher *fsnotify.Watcher
	EventChan chan string
	WatchDir  string
}

// Creates a new watcher with given config
func NewWatcher(dir string, eChan chan string) (*Watcher, error) {
	watcher := Watcher{}

	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %s", err)
	}

	err = fswatcher.Add(dir)
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %s", err)
	}

	watcher.FSWatcher = fswatcher
	watcher.EventChan = eChan
	watcher.WatchDir = dir

	return &watcher, nil
}

func (w *Watcher) Listen() {

	log.Println("Watcher Started for " + w.WatchDir)

	for {

		select {

		case event, ok := <-w.FSWatcher.Events:
			if !ok {
				continue
			}

			// Only look for write events
			if event.Has(fsnotify.Write) {
				log.Println("New Write on " + event.Name)
			}

			w.EventChan <- event.Name

		case err, ok := <-w.FSWatcher.Errors:
			if !ok {

				continue
			}

			log.Println("Error on Listening: ", err)

		}

	}

}
