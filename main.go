package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

var (
	listStyle = lipgloss.NewStyle().
			Margin(0, 0).
			Border(lipgloss.RoundedBorder())

	viewportStyle = lipgloss.NewStyle().
			Margin(0, 0).
			Border(lipgloss.RoundedBorder())

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()

	logfile, _ = os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	logger     = log.New(logfile)
)

type Exercise struct {
	title, description, content string
}

func (i Exercise) Title() string       { return i.title }
func (i Exercise) FilterValue() string { return i.title }
func (i Exercise) Description() string { return i.description }

type model struct {
	list     list.Model
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "k":
		}
	case tea.WindowSizeMsg:
		listMarginWidth, listMarginHeight := listStyle.GetFrameSize()
		viewportMarginWidth, viewportMarginHeight := viewportStyle.GetFrameSize()

		listWidth := lipgloss.Width(m.list.View()) + listMarginWidth
		listHeight := msg.Height - listMarginHeight
		viewportWidth := msg.Width - viewportMarginWidth - listWidth - 1
		viewportHeight := msg.Height - viewportMarginHeight

		m.list.SetSize(listWidth, listHeight)
		m.viewport.Width = viewportWidth
		m.viewport.Height = viewportHeight

		selectedItem := m.list.SelectedItem()
		selectedExercise := selectedItem.(Exercise)
		m.viewport.SetContent(selectedExercise.content)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	m.list.Title = "Exercises"

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		listStyle.Render(m.list.View()),
		viewportStyle.Render(m.viewport.View()),
	)
}

func (m model) headerView() string {
	title := titleStyle.Render("Header")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat(
		"─",
		max(0, m.viewport.Width-lipgloss.Width(info)-lipgloss.Width(m.list.View())),
	)
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

type T struct {
	Title       string
	Description string
	Content     string
}

func initializeModel() model {
	files, err := os.ReadDir("./exercises/")
	if err != nil {
		log.Fatal(err)
	}

	var items []list.Item
	for _, file := range files {
		exercise, err := os.ReadFile("exercises/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		var t T
		err = yaml.Unmarshal([]byte(exercise), &t)
		if err != nil {
			logger.Debug(err)
		}

		items = append(
			items,
			Exercise{title: t.Title, description: t.Description, content: t.Content},
		)
	}

	return model{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		viewport: viewport.New(0, 0),
	}
}

func main() {
	if _, err := tea.NewProgram(
		initializeModel(),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
