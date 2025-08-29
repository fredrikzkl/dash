package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

type inputView struct {
	description string
	placeholder string
}

var newEntryInputView = inputView{
	description: "Add new directory",
	placeholder: "Path",
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

func (m *Model) getDefaultAddInput() string {
	return fmt.Sprintf(
		"Add new directory: \n\n%s\n\n%s",
		m.input.View(),
		"(esc to cancel)",
	) + "\n"
}
