package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
)

func TestNewModel(t *testing.T) {
	//tests := []struct {
	//	keystrokes []tea.Key
	//	model      Model
	//}{
	//	{
	//		keystrokes: []tea.Key{
	//			{Type: tea.KeyRunes, Runes: []rune("j")},
	//		},
	//	},
	//}

	tm := teatest.NewTestModel(t, NewModel())

	tm.Send(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune("j"),
	})

	tm.Send(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune("q"),
	})

	fm := tm.FinalModel(t)
	m, _ := fm.(Model)

	if m.cursor != 1 {
		t.Fatalf("want: %v, got: %v", 1, m.cursor)
	}
}
