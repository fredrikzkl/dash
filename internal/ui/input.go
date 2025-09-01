package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputView struct {
	description   string
	placeholder   string
	confirmAction func(m *Model) tea.Cmd
}

var newEntryInputView = inputView{
	description: "Add new directory",
	placeholder: "Path",
	confirmAction: func(m *Model) tea.Cmd {
		entry, cmd := addNewEntry(m.input.Value())
		m.choices = append(m.choices, entry)
		return cmd
	},
}

var editCmdInputView = inputView{
	description: "Set cmd",
	placeholder: "e.g nvim",
	confirmAction: func(m *Model) tea.Cmd {
		if err := editCommand(m); err != nil {
			return func() tea.Msg { return err }
		}

		return nil
	},
}

func newTextInput() textinput.Model {
	ti := textinput.New()
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
