package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func main() {
	config := ui.NewConfig()
	config.ExerciseDatabase = "/Users/chris/Code/personal/infrastructure-exercism-prototype/bin/exercises.db"
	config.LogFile = "/Users/chris/Code/personal/infrastructure-exercism-prototype/log/iep.log"

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger := ui.NewLogger(log.DebugLevel, logfile)

	err = db.RunMigrations(config, logger)
	if err != nil {
		logger.Fatal(err)
	}

	if _, err := tea.NewProgram(
		ui.NewModel(config, logger),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
