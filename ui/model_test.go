package ui

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/teatest"
)

func TestNewModel(t *testing.T) {
	testcases := []struct {
		name       string
		keystrokes []tea.KeyMsg
		modelWant  Model
	}{
		{
			name:      "j",
			modelWant: Model{cursor: 1},
			keystrokes: []tea.KeyMsg{
				{Type: tea.KeyRunes, Runes: []rune("j")},
			},
		},
		{
			name:      "j-k",
			modelWant: Model{cursor: 0},
			keystrokes: []tea.KeyMsg{
				{Type: tea.KeyRunes, Runes: []rune("j")},
				{Type: tea.KeyRunes, Runes: []rune("k")},
			},
		},
	}

	config := Config{
		ExerciseDatabase: ":memory:",
	}

	var buf bytes.Buffer
	logger := NewLogger(log.DebugLevel, &buf)
	testModel := teatest.NewTestModel(t, NewModel(config, logger))

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			for _, keystroke := range testcase.keystrokes {
				testModel.Send(keystroke)
			}

			testModel.Send(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("q"),
			})

			modelGot := testModel.FinalModel(t).(Model)
			modelWant := testcase.modelWant

			if modelGot.cursor != modelWant.cursor {
				t.Fatalf("want: %v, got: %v", modelWant.cursor, modelGot.cursor)
			}
		})
	}
}
