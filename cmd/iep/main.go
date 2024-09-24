package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	ui "github.com/living-etc/iep/ui"
)

func main() {
	ui.Logger.SetLevel(log.DebugLevel)

	if _, err := tea.NewProgram(
		ui.NewModel(),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
