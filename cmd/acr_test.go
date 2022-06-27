package cmd

import (
	"io"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestAppcontrol(t *testing.T, conf *Configuration, dispatchChan chan string, doneChan chan bool) *AppControl {
	return NewAppcontrol(conf, dispatchChan, doneChan)
}

func TestInitCmd(t *testing.T) {
	doneChan := make(chan bool, 10)
	dispatchChan := make(chan string)
	testcases := []struct {
		name       string
		buildstubs func(*testing.T) (*exec.Cmd, io.ReadCloser, error)
		check      func(*testing.T, *exec.Cmd, io.ReadCloser, error)
	}{
		{
			name: "OK",
			buildstubs: func(t *testing.T) (*exec.Cmd, io.ReadCloser, error) {
				conf := randomConfigForNode()
				appct := newTestAppcontrol(t, conf, dispatchChan, doneChan)
				return appct.initCmd()
			},
			check: func(t *testing.T, cmd *exec.Cmd, rc io.ReadCloser, err error) {
				require.NotEqual(t, nil, cmd)
				require.NotEqual(t, nil, rc)
				require.Equal(t, nil, err)
			},
		},
		{
			name: "Unsupport languaguage",
			buildstubs: func(t *testing.T) (*exec.Cmd, io.ReadCloser, error) {
				conf := randomConfigForNode()
				conf.Language = "nodejs"
				appct := newTestAppcontrol(t, conf, dispatchChan, doneChan)
				return appct.initCmd()
			},
			check: func(t *testing.T, cmd *exec.Cmd, rc io.ReadCloser, err error) {
				require.NotEqual(t, nil, err)
				require.Equal(t, nil, rc)
			},
		},
		{
			name: "Have exec path languaguage",
			buildstubs: func(t *testing.T) (*exec.Cmd, io.ReadCloser, error) {
				conf := randomConfigForGo()
				appct := newTestAppcontrol(t, conf, dispatchChan, doneChan)
				return appct.initCmd()
			},
			check: func(t *testing.T, cmd *exec.Cmd, rc io.ReadCloser, err error) {
				require.NotEqual(t, nil, cmd)
				require.NotEqual(t, nil, rc)
				require.Equal(t, nil, err)
			},
		},
	}

	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			cmd, stdout, err := tc.buildstubs(t)
			tc.check(t, cmd, stdout, err)
		})
	}

	<-doneChan
}
