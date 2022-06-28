package main

import (
	"os"

	"github.com/9bany/ron/cmd"
	"github.com/9bany/ron/console"
)

func main() {

	done := make(chan bool)
	dispatch := make(chan string, 10)

	conf, err := cmd.InitConf()
	if err != nil {
		os.Exit(0)
	}
	cmdWatcher := cmd.NewWatcher(conf.RootPath, done, dispatch, conf.Ignore.Files, conf.Watch)
	appcontrol := cmd.NewAppcontrol(conf, dispatch, done)
	console.Intro()
	go appcontrol.Listening()
	go cmdWatcher.WaitingForChange()
	<-done
}
