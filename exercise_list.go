package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ExerciseList struct {
	list list.Model
}

func NewExerciseList(items []list.Item) ExerciseList {
	return ExerciseList{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

func (el ExerciseList) Update(msg tea.Msg) (ExerciseList, tea.Cmd) {
	var cmd tea.Cmd

	el.list, cmd = el.list.Update(msg)

	return el, cmd
}

func (el ExerciseList) View() string {
	return el.list.View()
}

func (el *ExerciseList) EnableScroll(enable bool) {
	el.list.KeyMap.CursorDown.SetEnabled(enable)
	el.list.KeyMap.CursorUp.SetEnabled(enable)
}

func (el ExerciseList) Help() ExerciseDescriptionHelp {
	return ExerciseDescriptionHelp{}
}
