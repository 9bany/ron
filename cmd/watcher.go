package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/9bany/ron/console"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	RootPath       string
	IgnorePath     []string
	Watch          Watch
	DispatcherChan chan string
	DoneChan       chan bool
	notifyWatcher  *fsnotify.Watcher
}

func NewWatcher(RootPath string,
	DoneChan chan bool,
	DispatcherChan chan string,
	IgnorePath []string, Watch Watch) *Watcher {
	return &Watcher{
		RootPath:       RootPath,
		DoneChan:       DoneChan,
		DispatcherChan: DispatcherChan,
		IgnorePath:     IgnorePath,
		Watch:          Watch,
	}
}

func (watcher *Watcher) walking(path string, fun func(path string, info fs.FileInfo, err error) error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		console.Error(err.Error())
		watcher.DoneChan <- true
	} else {
		if err = filepath.Walk(path, fun); err != nil {
			console.Error(err.Error())
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
			case fsnotify.Create:
				watcher.handleCreatenew(event.Name)
				watcher.DispatcherChan <- ACT_RESET
			case fsnotify.Write,
				fsnotify.Remove,
				fsnotify.Rename:
				if !strings.Contains(event.Name, fmt.Sprintf("%s.%s", FILE_NAME, EXTENSION)) {
					watcher.DispatcherChan <- ACT_RESET
				}

			}
		case err, ok := <-watcher.notifyWatcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}

}

func (watcher *Watcher) handleCreatenew(name string) {
	watcher.walking(name, watcher.fileListening)
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
