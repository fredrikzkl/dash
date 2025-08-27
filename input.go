package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

func initalInputModel() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Path"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 25
	return ti
}

func (m model) getDefaultInputView() string {
	return fmt.Sprintf(
		"Directory: \n\n%s\n\n%s",
		m.input.View(),
		"(esc to quit)",
	) + "\n"
}
