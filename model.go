package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
	exerciseList        ExerciseList
	exerciseDescription ExerciseDescription
	outputConsole       viewport.Model
	help                help.Model
	focused             string
	cursor              int
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
		exerciseList:        NewExerciseList(items),
		exerciseDescription: NewExerciseDescription(),
		outputConsole:       viewport.New(0, 0),
		help:                help.New(),
		focused:             "list",
	}

	model.exerciseList.list.KeyMap.CursorDown.SetEnabled(true)
	model.exerciseList.list.KeyMap.CursorUp.SetEnabled(true)

	model.exerciseDescription.viewport.KeyMap.Down.SetEnabled(false)
	model.exerciseDescription.viewport.KeyMap.Up.SetEnabled(false)

	model.outputConsole.KeyMap.Down.SetEnabled(false)
	model.outputConsole.KeyMap.Up.SetEnabled(false)

	model.exerciseList.list.Title = "Exercises"
	model.exerciseList.list.SetShowHelp(false)

	selectedItem := model.exerciseList.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	model.exerciseDescription.viewport.SetContent(glamouriseContent)

	model.outputLog = "Output Log"
	model.outputConsole.SetContent(model.outputLog)

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) updateSelectedExercise() {
	m.exerciseList.list.Select(m.cursor)

	selectedItem := m.exerciseList.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	m.exerciseDescription.viewport.SetContent(glamouriseContent)
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
			if m.focused == "list" && m.cursor < len(m.exerciseList.list.Items())-1 {
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

			m.exerciseDescription.viewport.KeyMap.Down.SetEnabled(enableViewport)
			m.exerciseDescription.viewport.KeyMap.Up.SetEnabled(enableViewport)

			m.exerciseList.list.KeyMap.CursorDown.SetEnabled(enableList)
			m.exerciseList.list.KeyMap.CursorUp.SetEnabled(enableList)

			m.outputConsole.KeyMap.Down.SetEnabled(enableOutputConsole)
			m.outputConsole.KeyMap.Up.SetEnabled(enableOutputConsole)
		case "enter":
			m.logEvent("Enter pressed")
		}
	case tea.WindowSizeMsg:
		styles := getStyles()

		_, frameHeight := styles.unfocused.GetFrameSize()

		scalingFactor := msg.Width / 100

		helpHeight := lipgloss.Height(m.help.View(HelpKeyMap{}))

		m.exerciseList.list.SetWidth(scalingFactor * 53)
		m.exerciseList.list.SetHeight(msg.Height - frameHeight - helpHeight)

		m.exerciseDescription.viewport.Width = scalingFactor * 80
		m.exerciseDescription.viewport.Height = msg.Height - frameHeight - helpHeight

		m.outputConsole.Width = scalingFactor * 52
		m.outputConsole.Height = msg.Height - frameHeight - helpHeight
	}

	m.exerciseDescription, cmd = m.exerciseDescription.Update(msg)
	cmds = append(cmds, cmd)

	m.outputConsole, cmd = m.outputConsole.Update(msg)
	cmds = append(cmds, cmd)

	m.exerciseList, cmd = m.exerciseList.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) logEvent(event string) {
	m.outputLog += fmt.Sprintf("\n[%v] %v", time.Now().Format("15:04:05"), event)
	m.outputConsole.SetContent(m.outputLog)
	m.outputConsole.GotoBottom()
}

type HelpKeyMap struct{}

func (keymap HelpKeyMap) ShortHelp() []key.Binding {
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
			key.WithKeys("enter"),
			key.WithHelp("enter", "Deploy Exercise"),
		),
		key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "Focus right"),
		),
	}
}

func (keymap HelpKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (m Model) View() string {
	var listRendered, exerciseDescriptionRendered, outputConsoleRendered string

	styles := getStyles()

	if m.focused == "list" {
		listRendered = styles.focused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())
	} else if m.focused == "viewport" {
		listRendered = styles.unfocused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.focused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())
	} else {
		listRendered = styles.unfocused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.focused.Render(m.outputConsole.View())
	}

	helpRendered := m.help.View(HelpKeyMap{})

	horizontal := lipgloss.JoinHorizontal(
		lipgloss.Top,
		listRendered,
		exerciseDescriptionRendered,
		outputConsoleRendered,
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		horizontal,
		helpRendered,
	)
}
