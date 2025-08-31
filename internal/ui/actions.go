package ui

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	s "github.com/fredrikzkl/dash/internal/storage"
)

func dash(entry s.Entry, executeCommand bool) tea.Cmd {
	var customCommand string
	if executeCommand && entry.Command != "" {
		customCommand += fmt.Sprintf("&& %s", entry.Command)
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s %s && exec $SHELL", entry.Path, customCommand))
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		if err != nil {
			return fmt.Errorf("failed to exectue dash: %w", err)
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

func editCommand(m *Model) error {
	if !choiceExists(*m) {
		return nil
	}

	m.choices[m.cursor].Command = m.input.Value()

	if err := s.SaveEntries(m.choices); err != nil {
		return err
	}

	return nil
}

func toggleCmd(m *Model) {
	m.cmdToggled = !m.cmdToggled
}

func getPwd() (string, error) {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
