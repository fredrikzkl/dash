package main

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func dash(entry entry) tea.Cmd {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && exec $SHELL", entry.Path))
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err != nil {
			return fmt.Errorf("failed to change directory: %w", err)
		}
		return tea.Quit()
	})
}

func addNewEntry(inputPath string) (entry, tea.Cmd) {
	entry := newEntry(inputPath)

	storedEntries, err := loadEntries()
	if err != nil {
		return entry, func() tea.Msg { return err }
	}

	storedEntries = append(storedEntries, entry)

	err = saveEntries(storedEntries)
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
