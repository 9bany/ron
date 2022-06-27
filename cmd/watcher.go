package cmd

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/9bany/ron/loger"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	RootPath       string
	IgnorePath     []string
	DispatcherChan chan string
	DoneChan       chan bool
	notifyWatcher  *fsnotify.Watcher
}

func NewWatcher(RootPath string,
	DoneChan chan bool,
	DispatcherChan chan string,
	IgnorePath []string) *Watcher {
	return &Watcher{
		RootPath:       RootPath,
		DoneChan:       DoneChan,
		DispatcherChan: DispatcherChan,
		IgnorePath:     IgnorePath,
	}
}

func (watcher *Watcher) walking(path string, fun func(path string, info fs.FileInfo, err error) error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		loger.Error(err.Error())
		watcher.DoneChan <- true
	} else {
		if err = filepath.Walk(path, fun); err != nil {
			loger.Error(err.Error())
			watcher.DoneChan <- true
		}
	}
}

func (watcher *Watcher) WaitingForChange() {
	notifyWatcher, err := fsnotify.NewWatcher()
	watcher.notifyWatcher = notifyWatcher
	if err != nil {
		log.Println(err)
		watcher.DoneChan <- true
	}

	defer notifyWatcher.Close()
	
	watcher.walking(watcher.RootPath, watcher.fileListening)
	// Ignore listening
	for _, path := range watcher.IgnorePath {
		watcher.walking(path, watcher.fileIgnoreListen)
	}

	for {
		select {
		case event, ok := <-watcher.notifyWatcher.Events:
			if !ok {
				return
			}
			switch event.Op {
			case fsnotify.Write,
				fsnotify.Create,
				fsnotify.Remove,
				fsnotify.Rename:
				watcher.DispatcherChan <- ACT_RESET
			}
		case err, ok := <-watcher.notifyWatcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}

}

func (watcher *Watcher) fileListening(path string, info fs.FileInfo, err error) error {
	if info.Mode().IsDir() {
		return watcher.notifyWatcher.Add(path)
	}
	return nil
}

func (watcher *Watcher) fileIgnoreListen(path string, info fs.FileInfo, err error) error {
	if info.Mode().IsDir() {
		return watcher.notifyWatcher.Remove(path)
	}
	return nil
}
