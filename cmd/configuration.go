package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/9bany/ron/console"
	"gopkg.in/yaml.v2"
)

// Configuration is the main file config for the ron app
// All the informations and config in the ron.yml
// And we convert it to the Configuration struct
type Configuration struct {
	RootPath        string   `yaml:"root_path"`
	ExecPath        string   `yaml:"exec_path"`
	Language        string   `yaml:"language"`
	WatchExtensions []string `yaml:"watch_extension"`
	IgnorePath      []string `yaml:"ignore_path"`
}

func (conf *Configuration) validate() error {

	if conf.RootPath == "" {
		return errors.New(ERORR_ROOT_PATH_EMPTY)
	}

	if conf.ExecPath == "" {
		return errors.New(ERROR_EXEC_PATH_EMPTY)
	}

	if conf.Language == "" {
		return errors.New(ERROR_LANGUAGE_EMPTY)
	}

	if len(conf.WatchExtensions) == 0 {
		return errors.New(ERROR_EXTENSIONS_EMPTY)
	}

	return nil
}

// getConf: read the main file for `ron`
// and decode it into a new Configuration instance
func getConf(fileName string) (*Configuration, error) {
	var config *Configuration
	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	if err = config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Check the main file if exist return name file
// or not, will return error
func checkFileIfExist() (string, error) {
	yml := fmt.Sprintf("%s.%s", FILE_NAME, EXTENSION)

	if _, err := os.Stat(yml); os.IsNotExist(err) {
		console.Error(err.Error())
		return "", err
	}

	return yml, nil
}

// InitConf is the configuration function
// with responsibility to read and defining the Configuration instance
func InitConf() (*Configuration, error) {
	yml, err := checkFileIfExist()

	if err != nil {
		return nil, err
	}

	return getConf(yml)
}
