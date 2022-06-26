package cmd

import "log"

const (
	ACT_RESET = "ACT_RESET"
)

type AppControl struct {
	DispatchChan chan string
}

func (appct *AppControl) start() {
	log.Println("Start server")
}

func (appct *AppControl) restart() {
	log.Println("restart server")
}

func (appct *AppControl) stop() {
	log.Println("Stop server")
} 

func (appct *AppControl) Listening() {
	for {
		select {
		case action:= <- appct.DispatchChan:
			switch action {
			case ACT_RESET:
				appct.restart()
			}
		}
	}
}