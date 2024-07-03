package main

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
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
)

type Model struct {
	list     list.Model
	viewport viewport.Model
	focused  string
	cursor   int
}

type T struct {
	Title       string
	Description string
	Content     string
}

func NewModel() Model {
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

	model := Model{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		viewport: viewport.New(0, 0),
		focused:  "list",
	}

	model.viewport.KeyMap.Down.SetEnabled(false)
	model.viewport.KeyMap.Up.SetEnabled(false)

	model.list.KeyMap.CursorDown.SetEnabled(true)
	model.list.KeyMap.CursorUp.SetEnabled(true)

	model.list.Title = "Exercises"

	selectedItem := model.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	model.viewport.SetContent(glamouriseContent)

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down", "j":
			if m.focused == "list" && m.cursor < len(m.list.Items())-1 {
				m.cursor++

				m.list.Select(m.cursor)

				selectedItem := m.list.SelectedItem()
				selectedExercise := selectedItem.(Exercise)

				glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
				if err != nil {
					log.Fatal(err)
				}
				m.viewport.SetContent(glamouriseContent)
			}
		case "up", "k":
			if m.focused == "list" && m.cursor > 0 {
				m.cursor--

				m.list.Select(m.cursor)

				selectedItem := m.list.SelectedItem()
				selectedExercise := selectedItem.(Exercise)

				glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
				if err != nil {
					log.Fatal(err)
				}
				m.viewport.SetContent(glamouriseContent)
			}
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
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
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
