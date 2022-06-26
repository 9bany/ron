package main

import (
	"fmt"

	"github.com/9bany/ron/cmd"
)

func main() {
	done := make(chan bool, 2)
	dispatch := make(chan string, 2)
	cmdWatcher := cmd.NewWatcher("./examples/nodejs", done, dispatch, []string{
		"./examples/nodejs/tests/",
	})
	go cmdWatcher.Listening()
	fmt.Println("Ron reloader server")
	<-done
}
