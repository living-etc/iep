package ui_test

import (
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/teatest"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func TestNewModel(t *testing.T) {
	testcases := []struct {
		name       string
		keystrokes []tea.KeyMsg
		modelWant  ui.Model
	}{
		{
			name:      "j",
			modelWant: ui.Model{Cursor: 1},
			keystrokes: []tea.KeyMsg{
				{Type: tea.KeyRunes, Runes: []rune("j")},
			},
		},
		{
			name:      "j-k",
			modelWant: ui.Model{Cursor: 0},
			keystrokes: []tea.KeyMsg{
				{Type: tea.KeyRunes, Runes: []rune("j")},
				{Type: tea.KeyRunes, Runes: []rune("k")},
			},
		},
	}

	config := ui.Config{
		ExerciseDatabase: ":memory:",
		LogFile:          "/Users/chris/Code/personal/infrastructure-exercism-prototype/log/test.log",
	}

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	logger := ui.NewLogger(log.DebugLevel, logfile)

	db.RunMigrations(config, logger)

	testModel := teatest.NewTestModel(t, ui.NewModel(config, logger))

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			for _, keystroke := range testcase.keystrokes {
				testModel.Send(keystroke)
			}

			testModel.Send(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("q"),
			})

			modelGot := testModel.FinalModel(t).(ui.Model)
			modelWant := testcase.modelWant

			if modelGot.Cursor != modelWant.Cursor {
				t.Fatalf("want: %v, got: %v", modelWant.Cursor, modelGot.Cursor)
			}
		})
	}
}
