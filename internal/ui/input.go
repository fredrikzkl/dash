package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputView struct {
	description   string
	placeholder   string
	confirmAction func(m *Model, input string) tea.Cmd
}

var newEntryInputView = inputView{
	description: "Add new directory",
	placeholder: "Path",
	confirmAction: func(m *Model, input string) tea.Cmd {
		entry, cmd := addNewEntry(m.input.Value())
		m.choices = append(m.choices, entry)
		return cmd
	},
}

var modifyCmdInputView = inputView{
	description: "Set cmd",
	placeholder: "e.g nvim",
}

func newTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Path"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 25
	return ti
}

func presetInput(m *Model, val string) {
	m.input.SetValue(val)
	m.input.SetCursor(len(val))
}

func (m *Model) getInputView(inputView inputView) string {
	m.input.Placeholder = inputView.placeholder
	return fmt.Sprintf(
		"%s \n\n%s\n\n%s",
		inputView.description,
		m.input.View(),
		"(esc to cancel)",
	) + "\n"
}
