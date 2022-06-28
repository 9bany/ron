package cmd

import (
	"errors"
	"os/exec"
)

type Language struct {
	BinPath       string
	ExecPath      string
	ExecCmd       string
	ProcessRegexp string
}

var supportedLanguages = map[string]*Language{
	"go": {
		BinPath:       "go",
		ExecCmd:       "run",
		ProcessRegexp: `(\d+).* %s`,
	},
	"node": {
		BinPath:       "node",
		ProcessRegexp: `(\d+).* %s`,
	},
	"ts-node": {
		BinPath:       "npx",
		ExecCmd:       "ts-node",
		ProcessRegexp: `(\d+).* %s`,
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
