package main

import (
	"log"

	"github.com/9bany/ron/cmd"
)

func main() {

	done := make(chan bool)
	dispatch := make(chan string, 10)

	conf, err := cmd.InitConf()
	if err != nil {
		log.Panic(err.Error())
	}
	cmdWatcher := cmd.NewWatcher(conf.RootPath, done, dispatch, conf.Ignore.Files)
	appcontrol := cmd.NewAppcontrol(conf, dispatch, done)

	go appcontrol.Listening()
	go cmdWatcher.WaitingForChange()
	<-done
}
