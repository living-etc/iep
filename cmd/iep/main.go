package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	ui "github.com/living-etc/iep/ui"
)

func main() {
	config := ui.NewConfig()
	config.ExerciseDatabase = "db/exercises.db"

	logfile, err := os.OpenFile("./log/iep", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger := ui.NewLogger(log.DebugLevel, logfile)

	if _, err := tea.NewProgram(
		ui.NewModel(config, logger),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
