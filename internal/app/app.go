package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fredrikzkl/dash/internal/ui"
)

func Run() error {
	m, err := ui.InitialModel()
	if err != nil {
		return fmt.Errorf("failed to initialize model: %w", err)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	return nil
}
