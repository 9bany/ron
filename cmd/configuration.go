package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration is the main file config for the ron app
// All the informations and config in the ron.yml
// And we convert it to the Configuration struct
type Configuration struct {
	RootPath string `yaml:"root_path"`
	ExecPath string `yaml:"exec_path"`
	Language string `yaml:"language"`
	Watch    Watch  `yaml:"watch"`
	Ignore   Ignore `yaml:"ignore"`
}

// getConf: read the main file for `ron`
// and decode it into a new Configuration instance
func getConf(fileName string) (*Configuration, error) {
	var config *Configuration
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Check the main file if exist return name file
// or not, will return error
func checkFileIfExist() (string, error) {
	yml := fmt.Sprintf("%s.%s", FILE_NAME, EXTENSION)

	if _, err := os.Stat(yml); !os.IsNotExist(err) {
		return yml, nil
	}

	return "", errors.New("ron.yml file not found")
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
