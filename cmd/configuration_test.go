package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var ronyml = []byte(`root_path: "./"
exec_path: "index.js"
language: "node"
watch:
  files:
    - index
    - server
  extensions: 
    - js
    - ts
ignore:
  files:
    - config
  extensions: 
    - js
    - ts`)

var ronNotyml = []byte(`fuck_you_file`)

const nameFile = FILE_NAME + "." + EXTENSION

func Before(nameFile string, content []byte) error {

	if err := ioutil.WriteFile(nameFile, content, 0644); err != nil {
		return err
	}
	return nil
}

func After(nameFile string) error {
	return os.Remove(nameFile)
}

func TestInitConfig(t *testing.T) {

	testCases := []struct {
		Name       string
		buildStubs func(t *testing.T) (*Configuration, error)
		check      func(t *testing.T, conf *Configuration, err error)
	}{
		{
			Name: "OK",
			buildStubs: func(t *testing.T) (*Configuration, error) {
				err := Before(nameFile, ronyml)
				if err != nil {
					log.Panic("Can not init unit test")
				}
				return InitConf()
			},
			check: func(t *testing.T, conf *Configuration, err error) {
				require.Equal(t, err, nil)
				require.Equal(t, conf.RootPath, "./")
				require.Equal(t, conf.ExecPath, "index.js")
				require.Equal(t, conf.Language, "node")

				require.Equal(t, len(conf.Watch.Extensions), 2)

				require.Equal(t, conf.Watch.Extensions[0], "js")
				require.Equal(t, conf.Watch.Extensions[1], "ts")

				require.Equal(t, len(conf.Ignore.Files), 1)
				require.Equal(t, len(conf.Ignore.Extensions), 2)

				require.Equal(t, conf.Ignore.Files[0], "config")

				require.Equal(t, conf.Ignore.Extensions[0], "js")
				require.Equal(t, conf.Ignore.Extensions[1], "ts")

				err = After(nameFile)
				if err != nil {
					log.Panic("Can not clear after run unit test")
				}
			},
		},
		{
			Name: "File not found",
			buildStubs: func(t *testing.T) (*Configuration, error) {

				return InitConf()
			},
			check: func(t *testing.T, conf *Configuration, err error) {
				require.NotEqual(t, err, nil)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			conf, err := tc.buildStubs(t)
			tc.check(t, conf, err)
		})
	}
}

func TestGetConfiguration(t *testing.T) {
	testCases := []struct {
		Name       string
		buildStubs func(t *testing.T) error
		check      func(t *testing.T, err error)
	}{
		{
			Name: "File not Found",
			buildStubs: func(t *testing.T) error {
				_, err := getConf("")
				return err
			},
			check: func(t *testing.T, err error) {
				require.NotEqual(t, err, nil)
			},
		},

		{
			Name: "Unmarshal error",
			buildStubs: func(t *testing.T) error {
				err := Before(nameFile, ronNotyml)
				if err != nil {
					log.Panic("Can not init unit test")
				}
				_, err = getConf(nameFile)
				return err
			},
			check: func(t *testing.T, err error) {
				require.NotEqual(t, err, nil)
				err = After(nameFile)
				if err != nil {
					log.Panic("Can not clear after run unit test")
				}
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			err := tc.buildStubs(t)
			tc.check(t, err)
		})
	}
}
