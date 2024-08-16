package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type OutputConsole struct {
	viewport viewport.Model
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
