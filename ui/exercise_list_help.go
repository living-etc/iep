package ui

import "github.com/charmbracelet/bubbles/key"

type ExerciseListHelp struct{}

func (keymap ExerciseListHelp) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↑/k", "Select Up"),
		),
		key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "Move Down"),
		),
		key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "Focus Right"),
		),
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Deploy Exercise"),
		),
		key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "Run Tests"),
		),
	}
}

func (keymap ExerciseListHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
