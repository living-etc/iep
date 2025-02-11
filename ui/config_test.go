package ui_test

import (
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
		environmentVariables []EnvVar
		configWant           ui.Config
	}{
		{
			name: "XDG_*_HOME_default_value",
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/devops"},
				{name: "XDG_STATE_HOME", value: "/home/devops/.local/state"},
				{name: "XDG_DATA_HOME", value: "/home/devops/.local/share"},
				{name: "XDG_CONFIG_HOME", value: "/home/devops/.config"},
			},
			configWant: ui.Config{
				ExerciseDatabase: "/home/devops/.local/share/iep/exercises.db",
				LogFile:          "/home/devops/.local/state/iep/iep.log",
				ConfigFile:       "/home/devops/.config/iep/config.toml",
			},
		},
		{
			name: "XDG_*_HOME_custom_value",
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/devops"},
				{name: "XDG_STATE_HOME", value: "/home/devops/.state"},
				{name: "XDG_DATA_HOME", value: "/home/devops/.share"},
				{name: "XDG_CONFIG_HOME", value: "/home/devops/.settings"},
			},
			configWant: ui.Config{
				ExerciseDatabase: "/home/devops/.share/iep/exercises.db",
				LogFile:          "/home/devops/.state/iep/iep.log",
				ConfigFile:       "/home/devops/.settings/iep/config.toml",
			},
		},
		{
			name: "XDG_*_HOME_undefined",
			environmentVariables: []EnvVar{
				{name: "HOME", value: "/home/devops"},
				{name: "XDG_STATE_HOME", value: ""},
				{name: "XDG_DATA_HOME", value: ""},
				{name: "XDG_CONFIG_HOME", value: ""},
			},
			configWant: ui.Config{
				ExerciseDatabase: "/home/devops/.local/share/iep/exercises.db",
				LogFile:          "/home/devops/.local/state/iep/iep.log",
				ConfigFile:       "/home/devops/.config/iep/config.toml",
			},
		},
	}

	for _, tt := range test_cases {
		for _, envvar := range tt.environmentVariables {
			t.Setenv(envvar.name, envvar.value)
		}

		config := ui.NewConfig()

		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.configWant, config); diff != "" {
				t.Error(diff)
			}
		})
	}
}
