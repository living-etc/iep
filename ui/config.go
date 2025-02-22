package ui

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ExerciseDatabase string `json:"exercises-db-file"`
	LogFile          string `json:"log-file"`
	MigrationsPath   string `json:"migrations-path"`
	ConfigFile       string
}

const (
	APP_NAME         = "iep"
	EXERCISE_DB_FILE = "exercises.db"
	LOG_FILE         = "iep.log"
	CONFIG_FILE      = "config.json"
)

func NewConfig(configJson []byte) (*Config, error) {
	config := defaultConfig()

	if len(configJson) > 0 {
		config, err := parseJsonConfig(configJson)
		if err != nil {
			return nil, err
		}
		return config, nil
	}

	configFileContent, err := os.ReadFile(config.ConfigFile)
	if err == nil {
		config, err = parseJsonConfig(configFileContent)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func defaultConfig() *Config {
	xdgHomePaths := map[string]string{
		"stateHome":  getXdgEnv("XDG_STATE_HOME"),
		"dataHome":   getXdgEnv("XDG_DATA_HOME"),
		"configHome": getXdgEnv("XDG_CONFIG_HOME"),
	}

	configDefaultValues := map[string]string{
		"exercises-db-file": filepath.Join(xdgHomePaths["dataHome"], APP_NAME, EXERCISE_DB_FILE),
		"log-file":          filepath.Join(xdgHomePaths["stateHome"], APP_NAME, LOG_FILE),
		"config-file":       filepath.Join(xdgHomePaths["configHome"], APP_NAME, CONFIG_FILE),
		"migrations-path":   "",
	}

	return &Config{
		ExerciseDatabase: configDefaultValues["exercises-db-file"],
		LogFile:          configDefaultValues["log-file"],
		MigrationsPath:   configDefaultValues["migrations-path"],
		ConfigFile:       configDefaultValues["config-file"],
	}
}

func parseJsonConfig(configJson []byte) (*Config, error) {
	var config Config

	json.Unmarshal(configJson, &config)

	return &config, nil
}

func getXdgEnv(envvar string) string {
	xdgEnvDefaultValues := map[string]string{
		"XDG_CONFIG_HOME": ".config",
		"XDG_DATA_HOME":   ".local/share",
		"XDG_STATE_HOME":  ".local/state",
	}

	var xdgEnv string
	if xdgEnv = os.Getenv(envvar); xdgEnv == "" {
		xdgEnv = filepath.Join(os.Getenv("HOME"), xdgEnvDefaultValues[envvar])
	}
	return xdgEnv
}
