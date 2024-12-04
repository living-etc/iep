package ui

import (
	"os"
	"path/filepath"
)

type Config struct {
	ExerciseDatabase string
}

// TODO need to handle situations where XDG_CONFIG_HOME isn't set
// Needs to work on Linux and Mac OS
func NewConfig() Config {
	return Config{
		ExerciseDatabase: filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "db/iep/exercises.db"),
	}
}
