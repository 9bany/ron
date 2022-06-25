package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Before() error {
	content := []byte(`root_path: "./"
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

	filePath := fmt.Sprintf("%s.%s", FILE_NAME, EXTENSION)

	if err := ioutil.WriteFile(filePath, content, 0644); err != nil {
		return err
	}
	return nil
}

func After() error {
	return os.Remove(fmt.Sprintf("%s.%s", FILE_NAME, EXTENSION))
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
				err := Before()
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

				require.Equal(t, len(conf.Watch.Files), 2)
				require.Equal(t, len(conf.Watch.Extensions), 2)

				require.Equal(t, conf.Watch.Files[0], "index")
				require.Equal(t, conf.Watch.Files[1], "server")

				require.Equal(t, conf.Watch.Extensions[0], "js")
				require.Equal(t, conf.Watch.Extensions[1], "ts")

				require.Equal(t, len(conf.Ignore.Files), 1)
				require.Equal(t, len(conf.Ignore.Extensions), 2)

				require.Equal(t, conf.Ignore.Files[0], "config")
				
				require.Equal(t, conf.Ignore.Extensions[0], "js")
				require.Equal(t, conf.Ignore.Extensions[1], "ts")

				err = After()
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
