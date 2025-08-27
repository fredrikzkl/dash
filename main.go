package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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

	err error
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
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

	s += m.getDefaultInputView()

	render(m.viewport, s)

	helpView := m.help.View(m.keys)
	height := 1 - strings.Count(helpView, "\n")
	spacing := strings.Repeat("\n", height)

	return m.viewport.View() + spacing + helpView
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Down):
			m.moveCursor(true)
			return m, nil

		case key.Matches(msg, m.keys.Up):
			m.moveCursor(false)
			return m, nil

		case key.Matches(msg, m.keys.Confirm):
			return m, dash(m.choices[m.cursor])

		case key.Matches(msg, m.keys.Add):
			// TODO: Add add!
			return m, nil

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func dash(entry entry) tea.Cmd {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && exec $SHELL", entry.Path))
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err != nil {
			return fmt.Errorf("failed to change directory: %w", err)
		}
		return tea.Quit()
	})
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
	}, nil
}
