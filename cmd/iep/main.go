package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func main() {
	config := ui.NewConfig()

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger := ui.NewLogger(log.DebugLevel, logfile)

	ctx := context.Background()

	conn, err := db.InitDb(ctx, config.ExerciseDatabase)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	err = db.RunMigrations(config, logger, conn)
	if err != nil {
		logger.Fatal(err)
	}

	if _, err := tea.NewProgram(
		ui.NewModel(config, logger, conn),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
