package main

type Exercise struct {
	title, description, content string
}

func (i Exercise) Title() string       { return i.title }
func (i Exercise) FilterValue() string { return i.title }
func (i Exercise) Description() string { return i.description }
