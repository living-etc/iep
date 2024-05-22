package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, description string
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }
func (i item) Description() string { return i.description }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	items := []list.Item{
		item{
			title:       "Deploy a web server",
			description: "Deploy a web application to a Linux environment",
		},
		item{
			title:       "Set up a subdomain",
			description: "Create a subdomain on a DNS zone",
		},
	}

	model := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	model.list.Title = "Exercises"

	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
