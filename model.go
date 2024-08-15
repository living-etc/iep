package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type styles struct {
	focused   lipgloss.Style
	unfocused lipgloss.Style
}

func getStyles() styles {
	return styles{
		unfocused: lipgloss.NewStyle().Margin(0, 0).Border(lipgloss.RoundedBorder()),
		focused: lipgloss.NewStyle().
			Margin(0, 0).
			BorderForeground(lipgloss.Color("63")).
			Border(lipgloss.RoundedBorder()),
	}
}

type Model struct {
	list                list.Model
	exerciseDescription viewport.Model
	focused             string
	cursor              int
	outputConsole       viewport.Model
	outputLog           string
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
		list:                list.New(items, list.NewDefaultDelegate(), 0, 0),
		exerciseDescription: viewport.New(0, 0),
		outputConsole:       viewport.New(0, 0),
		focused:             "list",
	}

	model.list.KeyMap.CursorDown.SetEnabled(true)
	model.list.KeyMap.CursorUp.SetEnabled(true)

	model.exerciseDescription.KeyMap.Down.SetEnabled(false)
	model.exerciseDescription.KeyMap.Up.SetEnabled(false)

	model.outputConsole.KeyMap.Down.SetEnabled(false)
	model.outputConsole.KeyMap.Up.SetEnabled(false)

	model.list.Title = "Exercises"

	selectedItem := model.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	model.exerciseDescription.SetContent(glamouriseContent)

	model.outputLog = "Output Log"
	model.outputConsole.SetContent(model.outputLog)

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) updateSelectedExercise() {
	m.list.Select(m.cursor)

	selectedItem := m.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	m.exerciseDescription.SetContent(glamouriseContent)
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
				m.updateSelectedExercise()
			}
		case "up", "k":
			if m.focused == "list" && m.cursor > 0 {
				m.cursor--
				m.updateSelectedExercise()
			}
		case "tab":
			var enableList, enableViewport, enableOutputConsole bool
			if m.focused == "list" {
				m.focused = "viewport"

				enableList = false
				enableViewport = true
				enableOutputConsole = false
			} else if m.focused == "viewport" {
				m.focused = "output"

				enableList = false
				enableViewport = false
				enableOutputConsole = true
			} else {
				m.focused = "list"

				enableList = true
				enableViewport = false
				enableOutputConsole = false
			}

			m.exerciseDescription.KeyMap.Down.SetEnabled(enableViewport)
			m.exerciseDescription.KeyMap.Up.SetEnabled(enableViewport)

			m.list.KeyMap.CursorDown.SetEnabled(enableList)
			m.list.KeyMap.CursorUp.SetEnabled(enableList)

			m.outputConsole.KeyMap.Down.SetEnabled(enableOutputConsole)
			m.outputConsole.KeyMap.Up.SetEnabled(enableOutputConsole)
		case "enter":
			m.logEvent("Enter pressed")
		}
	case tea.WindowSizeMsg:
		styles := getStyles()

		_, frameHeight := styles.unfocused.GetFrameSize()

		scalingFactor := msg.Width / 100

		m.list.SetWidth(scalingFactor * 53)
		m.list.SetHeight(msg.Height - frameHeight)

		m.exerciseDescription.Width = scalingFactor * 80
		m.exerciseDescription.Height = msg.Height - frameHeight

		m.outputConsole.Width = scalingFactor * 52
		m.outputConsole.Height = msg.Height - frameHeight
	}

	m.exerciseDescription, cmd = m.exerciseDescription.Update(msg)
	cmds = append(cmds, cmd)

	m.outputConsole, cmd = m.outputConsole.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m *Model) logEvent(event string) {
	m.outputLog += fmt.Sprintf("\n[%v] %v", time.Now().Format("15:04:05"), event)
	m.outputConsole.SetContent(m.outputLog)
	m.outputConsole.GotoBottom()
}

func (m Model) View() string {
	var listRendered, exerciseDescriptionRendered, outputConsoleRendered string

	styles := getStyles()

	if m.focused == "list" {
		listRendered = styles.focused.Render(m.list.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())
	} else if m.focused == "viewport" {
		listRendered = styles.unfocused.Render(m.list.View())
		exerciseDescriptionRendered = styles.focused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())
	} else {
		listRendered = styles.unfocused.Render(m.list.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.focused.Render(m.outputConsole.View())
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		listRendered,
		exerciseDescriptionRendered,
		outputConsoleRendered,
	)
}
