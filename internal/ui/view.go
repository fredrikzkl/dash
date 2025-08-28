package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var s string
	switch m.state {
	case MAIN_STATE:
		s = mainView(m)
	case ADD_STATE:
		s += m.getDefaultAddInput()
	}

	lineCount := strings.Count(s, "\n")
	m.viewport.Height = lineCount + 4

	render(m.viewport, s)

	helpView := m.help.View(m.keys)
	height := 1 - strings.Count(helpView, "\n")
	spacing := strings.Repeat("\n", height)

	return m.viewport.View() + spacing + helpView
}

func mainView(m Model) string {
	headerStyle := lipgloss.NewStyle().
		Italic(true).
		MarginBottom(1)

	hoverStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("4")) // 4

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
