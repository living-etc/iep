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
	APP_NAME               = "iep"
	EXERCISE_DB_FILE       = "exercises.db"
	LOG_FILE               = "iep.log"
	XDG_STATE_HOME_DEFAULT = ".local/state"
)

func NewConfig() Config {
	stateHome := os.Getenv("XDG_STATE_HOME")
	if stateHome == "" {
		stateHome = filepath.Join(os.Getenv("HOME"), XDG_STATE_HOME_DEFAULT)
	}

	return Config{
		ExerciseDatabase: filepath.Join(stateHome, APP_NAME, EXERCISE_DB_FILE),
		LogFile:          filepath.Join(stateHome, APP_NAME, LOG_FILE),
	}
}
