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
	XDG_DATA_HOME_DEFAULT  = ".local/share"
)

func NewConfig() Config {
	stateHome := os.Getenv("XDG_STATE_HOME")
	if stateHome == "" {
		stateHome = filepath.Join(os.Getenv("HOME"), XDG_STATE_HOME_DEFAULT)
	}

	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = filepath.Join(os.Getenv("HOME"), XDG_DATA_HOME_DEFAULT)
	}

	return Config{
		ExerciseDatabase: filepath.Join(dataHome, APP_NAME, EXERCISE_DB_FILE),
		LogFile:          filepath.Join(stateHome, APP_NAME, LOG_FILE),
	}
}
