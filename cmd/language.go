package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type Language struct {
	ProcessName   string
	BinPath       string
	ExecPath      string
	ExecCmd       string
	ProcessRegexp string
}

var supportedLanguages = map[string]*Language{
	"go": {
		ProcessName:   fmt.Sprintf("ron::golang::%d", time.Now().Unix()),
		BinPath:       "go",
		ExecCmd:       "run",
		ProcessRegexp: `(\d+).*\/go-build.*\/%s.*`,
	},
	"node": {
		ProcessName:   fmt.Sprintf("ron::ts-node::%d", time.Now().Unix()),
		BinPath:       "node",
		ProcessRegexp: `(\d+).* .*%s.*`,
	},
	"ts-node": {
		ProcessName:   fmt.Sprintf("ron::node::%d", time.Now().Unix()),
		BinPath:       "npx",
		ExecCmd:       "ts-node",
		ProcessRegexp: `(\d+).* .*%s.*`,
	},
}

func GetLanguage(name string) (*Language, error) {
	selectedLanguage := supportedLanguages[name]

	if selectedLanguage == nil {
		return nil, errors.New(ERROR_LANGUAGE_DOSE_NOT_SUPPORT)
	}

	// check language on machine
	execPath, err := exec.LookPath(selectedLanguage.BinPath)

	if err != nil {
		return nil, err
	}

	selectedLanguage.ExecPath = execPath

	return selectedLanguage, nil
}
