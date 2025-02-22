package ui_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/living-etc/iep/ui"
)

type EnvVar struct {
	name  string
	value string
}

func TestNewConfig(t *testing.T) {
	test_cases := []struct {
		name                 string
		configJson           []byte
		configFileJson       []byte
		environmentVariables []EnvVar
		configWant           *ui.Config
	}{
		{
			name:           "XDG_*_HOME_default_value",
			configJson:     []byte(""),
			configFileJson: []byte(""),
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/devops"},
				{name: "XDG_STATE_HOME", value: ""},
				{name: "XDG_DATA_HOME", value: ""},
				{name: "XDG_CONFIG_HOME", value: ""},
			},
			configWant: &ui.Config{
				ExerciseDatabase: "/home/devops/.local/share/iep/exercises.db",
				LogFile:          "/home/devops/.local/state/iep/iep.log",
				MigrationsPath:   "",
				ConfigFile:       "/home/devops/.config/iep/config.json",
			},
		},
		{
			name:           "XDG_*_HOME_custom_value",
			configJson:     []byte(""),
			configFileJson: []byte(""),
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/devops"},
				{name: "XDG_STATE_HOME", value: "/home/devops/.state"},
				{name: "XDG_DATA_HOME", value: "/home/devops/.share"},
				{name: "XDG_CONFIG_HOME", value: "/home/devops/.settings"},
			},
			configWant: &ui.Config{
				ExerciseDatabase: "/home/devops/.share/iep/exercises.db",
				LogFile:          "/home/devops/.state/iep/iep.log",
				MigrationsPath:   "",
				ConfigFile:       "/home/devops/.settings/iep/config.json",
			},
		},
		{
			name: "config literal specified",
			configJson: []byte(
				"{\"exercises-db-file\": \"/home/opsdev/.local/share/iep/exercises.db\", \"log-file\": \"/home/opsdev/.local/state/iep/iep.log\",\"migrations-path\":\"/home/opsdev/migrations\"}",
			),
			configFileJson: []byte(""),
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/opsdev"},
				{name: "XDG_STATE_HOME", value: "/home/opsdev/.local/share"},
				{name: "XDG_DATA_HOME", value: "/home/opsdev/.local/share"},
				{name: "XDG_CONFIG_HOME", value: "/home/opsdev/.config"},
			},
			configWant: &ui.Config{
				ExerciseDatabase: "/home/opsdev/.local/share/iep/exercises.db",
				LogFile:          "/home/opsdev/.local/state/iep/iep.log",
				MigrationsPath:   "/home/opsdev/migrations",
			},
		},
		{
			name:       "config file at XDG_CONFIG_HOME",
			configJson: []byte{},
			configFileJson: []byte(
				"{\"exercises-db-file\": \"/home/blah/.local/share/iep/exercises.db\", \"log-file\": \"/home/blah/.local/state/iep/iep.log\",\"migrations-path\":\"/home/blah/migrations\"}",
			),
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/blah"},
				{name: "XDG_STATE_HOME", value: ""},
				{name: "XDG_DATA_HOME", value: ""},
				{name: "XDG_CONFIG_HOME", value: "/tmp"},
			},
			configWant: &ui.Config{
				ExerciseDatabase: "/home/blah/.local/share/iep/exercises.db",
				LogFile:          "/home/blah/.local/state/iep/iep.log",
				MigrationsPath:   "",
				ConfigFile:       "/tmp/iep/config.json",
			},
		},
	}

	for _, tt := range test_cases {
		for _, envvar := range tt.environmentVariables {
			t.Setenv(envvar.name, envvar.value)
		}

		if len(tt.configFileJson) > 0 {
			err := os.WriteFile(
				filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "config.json"),
				tt.configFileJson,
				0777,
			)
			if err != nil {
				panic(err)
			}
		}

		config, err := ui.NewConfig(tt.configJson)
		if err != nil {
			panic(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.configWant, config); diff != "" {
				t.Error(diff)
			}
		})
	}
}
