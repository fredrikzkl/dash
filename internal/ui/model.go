package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fredrikzkl/dash/internal/storage"
)

// Model represents the application state
type Model struct {
	keys     keyMap
	help     help.Model
	viewport *viewport.Model
	choices  []storage.Entry
	cursor   int
	selected map[int]struct{}
	header   string
	input    textinput.Model
	state    State
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) View() string {
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

func mainView(m Model) string {
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

// InitialModel creates and returns the initial model
func InitialModel() (*Model, error) {
	vp, err := newViewport()
	if err != nil {
		return nil, fmt.Errorf("failed to create viewport: %w", err)
	}

	// Read header from assets
	header := "DASH"
	if b, err := os.ReadFile("assets/header.txt"); err == nil {
		header = string(b)
	}

	entries, err := storage.LoadEntries()
	if err != nil {
		return nil, fmt.Errorf("failed to load entries: %w", err)
	}

	return &Model{
		keys:     keys,
		help:     help.New(),
		viewport: &vp,
		choices:  entries,
		selected: make(map[int]struct{}),
		header:   header,
		input:    newTextInput(),
		state:    MAIN_STATE,
		cursor:   0,
	}, nil
}

// Helper methods
func (m *Model) moveCursor(down bool) {
	if len(m.choices) == 0 {
		return
	}

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

func (m *Model) setState(state State) {
	m.state = state
}

func (m *Model) getCurrentEntry() storage.Entry {
	if len(m.choices) == 0 || m.cursor >= len(m.choices) {
		return storage.Entry{}
	}
	return m.choices[m.cursor]
}

// Private helper functions
func newTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Path"
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 50
	return ti
}
