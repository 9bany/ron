package cmd

import (
	"os"

	"github.com/9bany/ron/console"
)

func Run() {
	done := make(chan bool)
	dispatch := make(chan string, 10)

	conf, err := InitConf()
	if err != nil {
		os.Exit(0)
	}
	cmdWatcher := NewWatcher(conf.RootPath, done, dispatch, conf.IgnorePath, conf.WatchExtensions)
	appcontrol := NewAppcontrol(conf, dispatch, done)
	console.Intro()
	go appcontrol.Listening()
	go cmdWatcher.WaitingForChange()
	<-done

}
