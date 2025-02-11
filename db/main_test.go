package db_test

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

	exitCode := m.Run()

	os.Exit(exitCode)
}
