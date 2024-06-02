package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

var (
	unfocusedStyle = lipgloss.NewStyle().
			Margin(0, 0).
			Border(lipgloss.RoundedBorder())

	focusedStyle = lipgloss.NewStyle().
			Margin(0, 0).
			BorderForeground(lipgloss.Color("63")).
			Border(lipgloss.RoundedBorder())

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
	focused  string
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
		case "tab":
			var enableList, enableViewport bool
			if m.focused == "list" {
				m.focused = "viewport"

				enableList = false
				enableViewport = true
			} else {
				m.focused = "list"

				enableList = true
				enableViewport = false
			}

			m.viewport.KeyMap.Down.SetEnabled(enableViewport)
			m.viewport.KeyMap.Up.SetEnabled(enableViewport)

			m.list.KeyMap.CursorDown.SetEnabled(enableList)
			m.list.KeyMap.CursorUp.SetEnabled(enableList)
		}
	case tea.WindowSizeMsg:
		listMarginWidth, listMarginHeight := unfocusedStyle.GetFrameSize()
		viewportMarginWidth, viewportMarginHeight := unfocusedStyle.GetFrameSize()

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

	if m.focused == "list" {
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			focusedStyle.Render(m.list.View()),
			unfocusedStyle.Render(m.viewport.View()),
		)
	} else {
		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			unfocusedStyle.Render(m.list.View()),
			focusedStyle.Render(m.viewport.View()),
		)
	}
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

	model := model{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		viewport: viewport.New(0, 0),
		focused:  "list",
	}

	model.viewport.KeyMap.Down.SetEnabled(false)
	model.viewport.KeyMap.Up.SetEnabled(false)

	model.list.KeyMap.CursorDown.SetEnabled(true)
	model.list.KeyMap.CursorUp.SetEnabled(true)

	return model
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
