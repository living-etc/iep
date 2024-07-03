package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
)

func TestNewModel(t *testing.T) {
	tests := []struct {
		keystrokes []tea.KeyMsg
		model      Model
	}{
		{
			model: Model{
				cursor: 1,
			},
			keystrokes: []tea.KeyMsg{
				{Type: tea.KeyRunes, Runes: []rune{'j'}},
			},
		},
	}

	testModel := teatest.NewTestModel(t, NewModel())

	for _, tt := range tests {
		for _, keystroke := range tt.keystrokes {
			testModel.Send(keystroke)
		}

		testModel.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("q"),
		})

		modelGot := testModel.FinalModel(t).(Model)
		modelWant := tt.model

		if modelGot.cursor != modelWant.cursor {
			t.Fatalf("want: %v, got: %v", modelWant.cursor, modelGot.cursor)
		}
	}
}
