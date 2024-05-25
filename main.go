package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
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
	ready    bool
	list     list.Model
	viewport viewport.Model
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
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			h, v := docStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v)

			content := "hi"
			m.viewport.SetContent(content)

			m.ready = true
		} else {
			h, v := docStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v)
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if !m.ready {
		return "\n Initialising...."
	}

	renderedViewport := fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		docStyle.Render(m.list.View()),
		renderedViewport,
	)
}

func (m model) headerView() string {
	return "header"
}

func (m model) footerView() string {
	return "footer"
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
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
