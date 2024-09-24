package ui

import "github.com/charmbracelet/bubbles/key"

type OutputConsoleHelp struct{}

func (keymap OutputConsoleHelp) ShortHelp() []key.Binding {
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
			key.WithKeys("/"),
			key.WithHelp("/", "Search"),
		),
	}
}

func (keymap OutputConsoleHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
