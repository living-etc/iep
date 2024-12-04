package ui_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/living-etc/iep/ui"
)

func TestFromToml(t *testing.T) {
	configWant := ui.Config{
		ExerciseDatabase: "/home/devops/.local/state/iep/db/exercises.db",
		LogFile:          "/home/devops/.local/state/iep/log/iep.log",
	}

	t.Setenv("HOME", "/home/devops")
	t.Setenv("XDG_STATE_HOME", os.Getenv("HOME")+"/.local/state")

	config := ui.NewConfig()

	t.Run("Read config from Toml", func(t *testing.T) {
		if diff := cmp.Diff(configWant, config); diff != "" {
			t.Error(diff)
		}
	})
}
