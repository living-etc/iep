package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

var (
	logfile, _ = os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	logger     = log.New(logfile)
)

func main() {
	logger.SetLevel(log.DebugLevel)

	if _, err := tea.NewProgram(
		NewModel(),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
