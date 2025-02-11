package ui_test

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	cwd := path.Join(path.Dir(filename), "..")

	os.Setenv("XDG_STATE_HOME", cwd+"/.local/state")
	os.Setenv("XDG_DATA_HOME", cwd+"/.local/share")

	exitCode := m.Run()

	os.Exit(exitCode)
}
