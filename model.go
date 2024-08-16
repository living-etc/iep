package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
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
	outputConsole       OutputConsole
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

	m := Model{
		exerciseList:        NewExerciseList(items),
		exerciseDescription: NewExerciseDescription(),
		outputConsole:       NewOutputConsole(),
		help:                help.New(),
		focused:             "list",
	}

	m.exerciseList.EnableScroll(true)
	m.exerciseDescription.EnableScroll(false)
	m.outputConsole.EnableScroll(false)

	m.exerciseList.list.Title = "Exercises"
	m.exerciseList.list.SetShowHelp(false)

	selectedItem := m.exerciseList.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		log.Fatal(err)
	}
	m.exerciseDescription.viewport.SetContent(glamouriseContent)

	m.outputLog = "Output Log"
	m.outputConsole.viewport.SetContent(m.outputLog)

	return m
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
			if m.focused == "list" {
				m.focused = "viewport"

				m.exerciseList.EnableScroll(false)
				m.exerciseDescription.EnableScroll(true)
				m.outputConsole.EnableScroll(false)
			} else if m.focused == "viewport" {
				m.focused = "output"

				m.exerciseList.EnableScroll(false)
				m.exerciseDescription.EnableScroll(false)
				m.outputConsole.EnableScroll(true)
			} else {
				m.focused = "list"

				m.exerciseList.EnableScroll(true)
				m.exerciseDescription.EnableScroll(false)
				m.outputConsole.EnableScroll(false)
			}
		case "enter":
			m.logEvent("Enter pressed")
		}
	case tea.WindowSizeMsg:
		styles := getStyles()

		_, frameHeight := styles.unfocused.GetFrameSize()

		scalingFactor := msg.Width / 100

		helpHeight := lipgloss.Height(m.help.View(ExerciseDescriptionHelp{}))

		m.exerciseList.list.SetWidth(scalingFactor * 53)
		m.exerciseList.list.SetHeight(msg.Height - frameHeight - helpHeight)

		m.exerciseDescription.viewport.Width = scalingFactor * 80
		m.exerciseDescription.viewport.Height = msg.Height - frameHeight - helpHeight

		m.outputConsole.viewport.Width = scalingFactor * 52
		m.outputConsole.viewport.Height = msg.Height - frameHeight - helpHeight
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
	m.outputConsole.viewport.SetContent(m.outputLog)
	m.outputConsole.viewport.GotoBottom()
}

func (m Model) View() string {
	var listRendered, exerciseDescriptionRendered, outputConsoleRendered, helpRendered string

	styles := getStyles()

	if m.focused == "list" {
		listRendered = styles.focused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())

		helpRendered = m.help.View(m.exerciseList.Help())
	} else if m.focused == "viewport" {
		listRendered = styles.unfocused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.focused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.unfocused.Render(m.outputConsole.View())

		helpRendered = m.help.View(m.exerciseDescription.Help())
	} else {
		listRendered = styles.unfocused.Render(m.exerciseList.View())
		exerciseDescriptionRendered = styles.unfocused.Render(m.exerciseDescription.View())
		outputConsoleRendered = styles.focused.Render(m.outputConsole.View())

		helpRendered = m.help.View(m.outputConsole.Help())
	}

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
