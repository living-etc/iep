package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ExerciseDescription struct {
	viewport viewport.Model
}

func NewExerciseDescription() ExerciseDescription {
	return ExerciseDescription{
		viewport: viewport.New(0, 0),
	}
}

func (ed ExerciseDescription) Update(msg tea.Msg) (ExerciseDescription, tea.Cmd) {
	var cmd tea.Cmd

	ed.viewport, cmd = ed.viewport.Update(msg)

	return ed, cmd
}

func (ed ExerciseDescription) View() string {
	return ed.viewport.View()
}

func (ed *ExerciseDescription) EnableScroll(enable bool) {
	ed.viewport.KeyMap.Down.SetEnabled(enable)
	ed.viewport.KeyMap.Up.SetEnabled(enable)
}

func (ed ExerciseDescription) Help() ExerciseDescriptionHelp {
	return ExerciseDescriptionHelp{}
}
