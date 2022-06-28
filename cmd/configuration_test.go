package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

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
	const nameFile = FILE_NAME + "." + EXTENSION
	var ronyml = []byte(`root_path: "./"
exec_path: "index.js"
language: "node"
watch_extension:
  - js
  - ts
ignore_path:
  - config`)

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
	const nameFile = FILE_NAME + "." + EXTENSION
	var ronNotyml = []byte(`fuck_you_file`)
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

func getTestcaseForValid(name string, ronyml []byte, errstring string) struct {
	name       string
	buildStubs func(t *testing.T) error
	check      func(t *testing.T, err error)
} {
	const nameFile = FILE_NAME + "." + EXTENSION
	return struct {
		name       string
		buildStubs func(t *testing.T) error
		check      func(t *testing.T, err error)
	}{
		name: name,
		buildStubs: func(t *testing.T) error {
			err := Before(nameFile, ronyml)
			if err != nil {
				log.Panic("Can not init unit test")
			}
			_, err = InitConf()
			return err
		},
		check: func(t *testing.T, err error) {
			require.Equal(t, err.Error(), errstring)
			err = After(nameFile)
			if err != nil {
				log.Panic("Can not clear after run unit test")
			}
		},
	}
}

func TestValidateConf(t *testing.T) {
	var ronymlEmptyRootPath = []byte(`
root_path: ""
exec_path: "index.js"
language: "node"
watch_extension:
  - js
  - ts
ignore_path:
  - config`)
	var ronymlEmptyExecPath = []byte(`
root_path: "./"
exec_path: ""
language: "node"
watch_extension:
  - js
  - ts
ignore_path:
  - config`)
	var ronymlEmptyLanguage = []byte(`
root_path: "./"
exec_path: "index.js"
language: ""
watch_extension:
  - js
  - ts
ignore_path:
  - config`)
	var ronymlEmptyWatchExtensions = []byte(`
root_path: "./"
exec_path: "index.js"
language: "node"
watch_extension:`)
	testcases := []struct {
		name       string
		buildStubs func(t *testing.T) error
		check      func(t *testing.T, err error)
	}{
		getTestcaseForValid("Rootpath empty", ronymlEmptyRootPath, ERORR_ROOT_PATH_EMPTY),
		getTestcaseForValid("Exexpath empty", ronymlEmptyExecPath, ERROR_EXEC_PATH_EMPTY),
		getTestcaseForValid("Language empty", ronymlEmptyLanguage, ERROR_LANGUAGE_EMPTY),
		getTestcaseForValid("Watch extesions empty", ronymlEmptyWatchExtensions, ERROR_EXTENSIONS_EMPTY),
	}

	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			err := tc.buildStubs(t)
			tc.check(t, err)
		})
	}
}
