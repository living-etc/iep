package ui_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/living-etc/iep/ui"
)

func TestFromToml(t *testing.T) {
	configWant := ui.Config{
		ExerciseDatabase: "/home/devops/.config/db/iep/exercises.db",
	}

	t.Setenv("XDG_CONFIG_HOME", "/home/devops/.config")

	config := ui.NewConfig()

	t.Run("Read config from Toml", func(t *testing.T) {
		if diff := cmp.Diff(configWant, config); diff != "" {
			t.Error(diff)
		}
	})
}
