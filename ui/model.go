package ui

import (
	"context"
	"database/sql"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

const (
	getAllExercisesQuery = "SELECT * FROM exercises;"
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
	Cursor              int
	logger              *log.Logger
}

func NewModel(config Config, logger *log.Logger, conn *sql.DB) Model {
	ctx := context.Background()

	rows, err := conn.QueryContext(ctx, getAllExercisesQuery)
	if err != nil {
		logger.Fatal(err)
	}
	defer rows.Close()

	var items []list.Item
	for rows.Next() {
		var e Exercise
		var ignoreInt int
		var ignoreString string
		if err := rows.Scan(&ignoreInt, &ignoreString, &e.title, &e.description, &e.content); err != nil {
			logger.Fatal(err)
		}
		items = append(items, e)
	}

	m := Model{
		exerciseList:        NewExerciseList(items),
		exerciseDescription: NewExerciseDescription(),
		outputConsole:       NewOutputConsole(),
		help:                help.New(),
		focused:             "list",
		logger:              logger,
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
		logger.Fatal(err)
	}
	m.exerciseDescription.viewport.SetContent(glamouriseContent)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) updateSelectedExercise() {
	m.exerciseList.list.Select(m.Cursor)

	selectedItem := m.exerciseList.list.SelectedItem()
	selectedExercise := selectedItem.(Exercise)

	glamouriseContent, err := glamour.Render(selectedExercise.content, "dark")
	if err != nil {
		m.logger.Fatal(err)
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
			m.logger.Debug("[Model::Update] keystroke [j]")

			if m.focused == "list" && m.Cursor < len(m.exerciseList.list.Items())-1 {
				m.Cursor++
				m.updateSelectedExercise()

				m.logger.Debug("[Model::Update] incrementing cursor", "cursor", m.Cursor)
			}
		case "up", "k":
			m.logger.Debug("[Model::Update] keystroke [k]")

			if m.focused == "list" && m.Cursor > 0 {
				m.Cursor--
				m.updateSelectedExercise()

				m.logger.Debug("[Model::Update] decrementing cursor", "cursor", m.Cursor)
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
			m.outputConsole.LogEvent("Enter pressed")
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
