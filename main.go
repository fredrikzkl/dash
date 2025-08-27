package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport *viewport.Model
	choices  []entry
	cursor   int
	selected map[int]struct{}
	header   string
	err      error
}

func (m model) Init() tea.Cmd {
	return nil
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
		s += fmt.Sprintf("No entries")
	}

	render(m.viewport, s)
	return m.viewport.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// fmt.Println("Key pressed: %s", msg.String())
		switch msg.String() {
		case "j", "down":
			m.moveCursor(true)
			return m, nil

		case "k", "up":
			m.moveCursor(false)
			return m, nil

		case "enter":
			return m, dash(m.choices[m.cursor])
		case "f": // Testing
			cmd := func() tea.Msg { return tea.ExitAltScreen() }
			return m, cmd

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
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
		viewport: &vp,
		choices:  loadedEntries,
		selected: make(map[int]struct{}),
		header:   header,
	}, nil
}
