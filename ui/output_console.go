package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type OutputConsole struct {
	viewport  viewport.Model
	outputLog string
}

func NewOutputConsole() OutputConsole {
	return OutputConsole{
		viewport: viewport.New(0, 0),
	}
}

func (oc OutputConsole) Update(msg tea.Msg) (OutputConsole, tea.Cmd) {
	var cmd tea.Cmd

	oc.viewport, cmd = oc.viewport.Update(msg)

	return oc, cmd
}

func (oc OutputConsole) View() string {
	return oc.viewport.View()
}

func (el *OutputConsole) EnableScroll(enable bool) {
	el.viewport.KeyMap.Down.SetEnabled(enable)
	el.viewport.KeyMap.Up.SetEnabled(enable)
}

func (ed OutputConsole) Help() OutputConsoleHelp {
	return OutputConsoleHelp{}
}

func (od *OutputConsole) LogEvent(event string) {
	od.outputLog += fmt.Sprintf("[%v] %v\n", time.Now().Format("15:04:05"), event)
	od.viewport.SetContent(od.outputLog)
	od.viewport.GotoBottom()
}
