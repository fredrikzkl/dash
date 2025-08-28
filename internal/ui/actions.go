package ui

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	s "github.com/fredrikzkl/dash/internal/storage"
)

func dash(entry s.Entry) tea.Cmd {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && exec $SHELL", entry.Path))
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err != nil {
			return fmt.Errorf("failed to change directory: %w", err)
		}
		return tea.Quit()
	})
}

func addNewEntry(inputPath string) (s.Entry, tea.Cmd) {
	entry := s.NewEntry(inputPath)

	storedEntries, err := s.LoadEntries()
	if err != nil {
		return entry, func() tea.Msg { return err }
	}

	storedEntries = append(storedEntries, entry)

	err = s.SaveEntries(storedEntries)
	if err != nil {
		return entry, func() tea.Msg { return err }
	}

	return entry, nil
}

func getPwd() (string, error) {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
