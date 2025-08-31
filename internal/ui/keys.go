package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Down          key.Binding
	Up            key.Binding
	Confirm       key.Binding
	Add           key.Binding
	Command       key.Binding
	ToggleCommand key.Binding
	Delete        key.Binding
	Help          key.Binding
	Back          key.Binding
	Quit          key.Binding
}

var keys = keyMap{
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
	),
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add entry"),
	),
	Command: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "edit cmd"),
	),
	ToggleCommand: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "toggle cmd"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete entry"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctr+c"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Add, k.Command, k.ToggleCommand, k.Delete,
	}
}

// TODO: Its just the same as short help
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Add},
		{},
	}
}
