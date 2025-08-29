package ui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var s string
	switch m.state {
	case MAIN_STATE:
		s = mainView(m)
	case ADD_STATE:
		s += m.getInputView(newEntryInputView)
	case COMMAND_STATE:
		s += m.getInputView(modifyCmdInputView)
	}

	lineCount := strings.Count(s, "\n")
	m.viewport.Height = lineCount + 4

	m.viewport.SetContent(s)

	helpView := m.help.View(m.keys)
	height := 1 - strings.Count(helpView, "\n")
	spacing := strings.Repeat("\n", height)

	return m.viewport.View() + spacing + helpView
}

func mainView(m Model) string {
	s := headerStyle.Render(m.header) + "\n"

	// Iterate over choices
	for i, entry := range m.choices {
		num := i + 1
		cursor := " " // no cursor

		// Cursor at point
		if m.cursor == i {
			cursor = ">" // cursor!
			s += hoverStyle.Render(fmt.Sprintf("%s %d. %s", cursor, num, entry.Name))
			s += "\n"
		} else {
			s += fmt.Sprintf("%s %d. %s\n", cursor, num, entry.Name)
		}
	}

	if len(m.choices) == 0 {
		s += "No entries"
	}
	return s
}

// func inputView(m Model, iv inputView) string {
// 	s += m.getDefaultAddInput()
// }
