package ui

import (
	"os"
	"path/filepath"
)

type Config struct {
	ExerciseDatabase string
	LogFile          string
	ConfigFile       string
}

const (
	APP_NAME         = "iep"
	EXERCISE_DB_FILE = "exercises.db"
	LOG_FILE         = "iep.log"
	CONFIG_FILE      = "config.toml"
)

var xdgEnvDefaultValues = map[string]string{
	"XDG_CONFIG_HOME": ".config",
	"XDG_DATA_HOME":   ".local/share",
	"XDG_STATE_HOME":  ".local/state",
}

func NewConfig() Config {
	stateHome := getXdgEnv("XDG_STATE_HOME")
	dataHome := getXdgEnv("XDG_DATA_HOME")
	configHome := getXdgEnv("XDG_CONFIG_HOME")

	return Config{
		ExerciseDatabase: filepath.Join(dataHome, APP_NAME, EXERCISE_DB_FILE),
		LogFile:          filepath.Join(stateHome, APP_NAME, LOG_FILE),
		ConfigFile:       filepath.Join(configHome, APP_NAME, CONFIG_FILE),
	}
}

func getXdgEnv(envvar string) string {
	var xdgEnv string
	if xdgEnv = os.Getenv(envvar); xdgEnv == "" {
		xdgEnv = filepath.Join(os.Getenv("HOME"), xdgEnvDefaultValues[envvar])
	}
	return xdgEnv
}
