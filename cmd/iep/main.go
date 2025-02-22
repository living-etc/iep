package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/living-etc/iep/db"
	"github.com/living-etc/iep/ui"
)

func main() {
	cwd, _ := os.Getwd()
	os.Setenv("HOME", cwd)
	configFilePath := flag.String(
		"config-file",
		"",
		"Path to a configuration file. This will override all other configuration.",
	)
	flag.Parse()

	var configFileContents []byte
	if _, err := os.Stat(*configFilePath); err == nil {
		configFileContents, err = os.ReadFile(*configFilePath)
		if err != nil {
			panic(err)
		}
	}

	config, err := ui.NewConfig(configFileContents)
	if err != nil {
		panic(err)
	}

	logfile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger := ui.NewLogger(log.DebugLevel, logfile)

	ctx := context.Background()

	conn, err := db.InitDb(ctx, config.ExerciseDatabase)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = db.RunMigrations(config, logger, conn)
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(
		ui.NewModel(config, logger, conn),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
