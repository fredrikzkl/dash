package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	keys keyMap
	help help.Model

	viewport *viewport.Model
	choices  []entry
	cursor   int
	selected map[int]struct{}
	header   string

	input textinput.Model

	state programState
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
	var s string
	switch m.state {
	case MAIN_STATE:
		s = mainView(m)
	case ADD_STATE:
		s += m.getDefaultAddInput()
	}

	render(m.viewport, s)

	helpView := m.help.View(m.keys)
	height := 1 - strings.Count(helpView, "\n")
	spacing := strings.Repeat("\n", height)

	return m.viewport.View() + spacing + helpView
}

func mainView(m model) string {
	headerStyle := lipgloss.NewStyle().
		MarginBottom(1)

	s := headerStyle.Render(m.header) + "\n"

	// Iterate over choices
	for i, entry := range m.choices {
		cursor := " " // nor cursor
		// Cursor at point
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// render the row
		num := i + 1
		s += fmt.Sprintf("%s %d. %s\n", cursor, num, entry.Name)
	}

	if len(m.choices) == 0 {
		s += "No entries"
	}
	return s
}

func (m *model) moveCursor(down bool) {
	maxIndex := len(m.choices) - 1
	if down {
		m.cursor++
		if m.cursor > maxIndex {
			m.cursor = 0
		}
	} else {
		m.cursor--
		if m.cursor < 0 {
			m.cursor = maxIndex
		}
	}
}

func (m *model) setState(state programState) {
	m.state = state
}

func main() {
	m, err := initialModel()
	if err != nil {
		fmt.Println("Error running Bash: %w", err)
		os.Exit(1)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		fmt.Println("Error running Dash")
		os.Exit(1)
	}
}

func initialModel() (*model, error) {
	vp, err := newViewport()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	// Read header.txt
	header := "DASH"
	if b, err := os.ReadFile("header.txt"); err == nil {
		header = string(b)
	}

	loadedEntries, err := loadEntries()
	if err != nil {
		return nil, fmt.Errorf("loading data failed: %w", err)
	}

	return &model{
		keys:     keys,
		help:     help.New(),
		viewport: &vp,
		choices:  loadedEntries,
		selected: make(map[int]struct{}),
		header:   header,
		input:    initalInputModel(),
		state:    MAIN_STATE,
	}, nil
}
