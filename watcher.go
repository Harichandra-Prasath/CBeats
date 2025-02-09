// File Watcher that listens for writes on a particular file

package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type WatcherConfig struct {
	Dir string
}

type Watcher struct {
	cfg       WatcherConfig
	FSWatcher *fsnotify.Watcher
}

// Creates a new watcher with given config
func NewWatcher(cfg WatcherConfig) (*Watcher, error) {
	watcher := Watcher{}

	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %s", err)
	}

	err = fswatcher.Add(cfg.Dir)
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %s", err)
	}

	watcher.FSWatcher = fswatcher
	watcher.cfg = cfg

	return &watcher, nil
}

func (w *Watcher) Listen() {
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

		case err, ok := <-w.FSWatcher.Errors:
			if !ok {

				continue
			}

			log.Println("Error on Listening: ", err)

		}

	}

}
