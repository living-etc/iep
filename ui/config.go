package ui

import (
	"os"
	"path/filepath"
)

type Config struct {
	ExerciseDatabase string
	LogFile          string
}

const (
	APP_NAME = "iep"
)

// TODO need to handle situations where XDG_CONFIG_HOME isn't set
// Needs to work on Linux and Mac OS
func NewConfig() Config {
	return Config{
		ExerciseDatabase: filepath.Join(os.Getenv("XDG_STATE_HOME"), APP_NAME, "db/exercises.db"),
		LogFile:          filepath.Join(os.Getenv("XDG_STATE_HOME"), APP_NAME, "log/iep.log"),
	}
}
