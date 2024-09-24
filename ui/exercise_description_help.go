package ui

import "github.com/charmbracelet/bubbles/key"

type ExerciseDescriptionHelp struct{}

func (keymap ExerciseDescriptionHelp) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↑/k", "Select up"),
		),
		key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "move down"),
		),
		key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "Focus right"),
		),
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Deploy Exercise"),
		),
	}
}

func (keymap ExerciseDescriptionHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
