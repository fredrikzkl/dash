package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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
