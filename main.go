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

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os/exec"
// )

// const ShellToUse = "bash"



// func main() {
// 	err, out, errout := Shellout("exec -a ron::golang::1656388862 /usr/local/bin/go run cmd/server.go")
// 	if err != nil {
// 		log.Printf("error: %v\n", err)
// 	}
// 	fmt.Println("--- stdout ---")
// 	fmt.Println(out)
// 	fmt.Println("--- stderr ---")
// 	fmt.Println(errout)
// }
