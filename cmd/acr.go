package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/9bany/ron/console"
)

const (
	ACT_RESET = "ACT_RESET"
	ACT_INIT  = "ACT_INIT"
)

type AppControl struct {
	DispatchChan     <-chan string
	DoneChan         chan bool
	Conf             *Configuration
	selectedLanguage *Language
}

func NewAppcontrol(conf *Configuration, DispatchChan chan string, DoneChan chan bool) *AppControl {
	return &AppControl{
		DispatchChan: DispatchChan,
		Conf:         conf,
		DoneChan:     DoneChan,
	}
}

func (appct *AppControl) initCmd() (*exec.Cmd, io.ReadCloser, error) {
	var cmd *exec.Cmd

	language, err := GetLanguage(appct.Conf.Language)
	if err != nil {
		log.Println(err.Error())
		appct.DoneChan <- true
		return nil, nil, err
	}
	appct.selectedLanguage = language
	if language.ExecCmd == "" {
		cmd = exec.Command(language.ExecPath, appct.Conf.ExecPath)
	} else {
		cmd = exec.Command(language.ExecPath, language.ExecCmd, appct.Conf.ExecPath)
	}

	stdout, err := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	if err != nil {
		log.Println(err.Error())
		appct.DoneChan <- true
		return nil, nil, err
	}
	return cmd, stdout, nil
}

func (appct *AppControl) start() {
	cmd, stdout, err := appct.initCmd()
	if err != nil {
		log.Println(err.Error())
		appct.DoneChan <- true
	}
	// start cmd
	if err = cmd.Start(); err != nil {
		log.Println(err.Error())
		appct.DoneChan <- true
	}

	// listen stdout and print out them
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)

		if !strings.Contains(string(tmp), "signal: killed") {
			fmt.Print(string(tmp))
		}

		if err != nil {
			stdout.Close()
			break

		}
	}

}

func (appct *AppControl) restart() {
	var r *regexp.Regexp
	var command string

	if runtime.GOOS == "linux" {
		command = "ps -u"
	} else if runtime.GOOS == "mac" || runtime.GOOS == "darwin" {
		command = "ps -a"
	}

	command += " | grep " + appct.selectedLanguage.BinPath
	b, _ := exec.Command("/bin/sh", "-c", command).Output()

	r = regexp.MustCompile(fmt.Sprintf(appct.selectedLanguage.ProcessRegexp, appct.Conf.ExecPath))
	match := r.FindStringSubmatch(string(b))

	if len(match) > 1 {
		i, err := strconv.Atoi(match[1])
		if err != nil {
			appct.DoneChan <- true
		}
		p, err := os.FindProcess(i)
		if err != nil {
			appct.DoneChan <- true
		}
		p.Kill()
		go appct.start()
	}
}

func (appct *AppControl) Listening() {
	go appct.start()
	console.Listening()
	for {
		select {
		case action, ok := <-appct.DispatchChan:
			if !ok {
				return
			}
			switch action {
			case ACT_RESET:
				appct.restart()
				console.Restarting()
			case ACT_INIT:
				appct.start()
			}
		case done := <-appct.DoneChan:
			if done {
				return
			}
		}
	}
}
