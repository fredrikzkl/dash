package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport *viewport.Model
	choices  []string
	cursor   int
	selected map[int]struct{}
	header   string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().
		MarginBottom(1)

	s := headerStyle.Render(m.header) + "\n"

	// Iterate over choices
	for i, choice := range m.choices {
		cursor := " " // nor cursor
		// Cursor at point
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // Seleceted
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	render(m.viewport, s)
	return m.viewport.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.moveCursor(true)
			return m, nil

		case "k", "up":
			m.moveCursor(false)
			return m, nil

		case "enter":
			// TODO:Something
			_, ok := m.selected[m.cursor]
			if ok {
				fmt.Printf("%s\n", m.choices[m.cursor])
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
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

	return &model{
		viewport: &vp,
		choices:  []string{"Vippsnummer", "Shopingbasket"},
		selected: make(map[int]struct{}),
		header:   header,
	}, nil
}
