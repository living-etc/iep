package ui

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ExerciseDatabase string `json:"exercises-db-file"`
	LogFile          string `json:"log-file"`
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

func NewConfig(configJson []byte) (*Config, error) {
	var config *Config

	if len(configJson) > 0 {
		config, err := parseJsonConfig(configJson)
		if err != nil {
			return nil, err
		}
		return config, nil
	}

	stateHome := getXdgEnv("XDG_STATE_HOME")
	dataHome := getXdgEnv("XDG_DATA_HOME")
	configHome := getXdgEnv("XDG_CONFIG_HOME")

	config = &Config{
		ExerciseDatabase: filepath.Join(dataHome, APP_NAME, EXERCISE_DB_FILE),
		LogFile:          filepath.Join(stateHome, APP_NAME, LOG_FILE),
		ConfigFile:       filepath.Join(configHome, APP_NAME, CONFIG_FILE),
	}

	return config, nil
}

func parseJsonConfig(configJson []byte) (*Config, error) {
	var config Config

	json.Unmarshal(configJson, &config)

	return &config, nil
}

func getXdgEnv(envvar string) string {
	var xdgEnv string
	if xdgEnv = os.Getenv(envvar); xdgEnv == "" {
		xdgEnv = filepath.Join(os.Getenv("HOME"), xdgEnvDefaultValues[envvar])
	}
	return xdgEnv
}
