package ui

import "github.com/charmbracelet/bubbles/key"

type OutputConsoleHelp struct{}

func (keymap OutputConsoleHelp) ShortHelp() []key.Binding {
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
			key.WithKeys("/"),
			key.WithHelp("/", "Search"),
		),
		key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "Run Tests"),
		),
	}
}

func (keymap OutputConsoleHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
