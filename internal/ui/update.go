package ui

import (
	"regexp"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fredrikzkl/dash/internal/storage"
)

var digitKey = regexp.MustCompile(`^[1-9]$`)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case MainState:
		return mainUpdate(msg, m)
	case AddState:
		return inputUpdate(msg, m, newEntryInputView)
	case CommandState:
		return inputUpdate(msg, m, editCmdInputView)
	}

	return m, nil
}

func mainUpdate(msg tea.Msg, m Model) (Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(msg, m.keys.Down):
			m.moveCursor(true)
			return m, nil

		case key.Matches(msg, m.keys.Up):
			m.moveCursor(false)
			return m, nil

		case key.Matches(msg, m.keys.Confirm):
			chosenEntry := m.choices[m.cursor]

			cmd := dash(chosenEntry, m.cmdToggled)
			moveEntryToTop(&m, chosenEntry)

			return m, cmd

			// New entry input
		case key.Matches(msg, m.keys.Add):
			m.setState(AddState)

			pwd, err := getPwd()
			if err == nil {
				presetInput(&m, pwd)
			}
			return m, nil

			// Command Input
		case key.Matches(msg, m.keys.Command):
			m.setState(CommandState)

			if !choiceExists(m) {
				return m, nil
			}

			currentCmd := m.choices[m.cursor].Command
			if currentCmd != "" {
				presetInput(&m, currentCmd)
			}

			return m, nil

		case key.Matches(msg, m.keys.ToggleCommand):
			if m.choices[m.cursor].Command == "" {
				return m, nil
			}
			toggleCmd(&m)
			return m, nil

		case key.Matches(msg, m.keys.Back):
			m.setState(MainState)
			return m, nil

		case key.Matches(msg, m.keys.Delete):
			if len(m.choices) == 0 {
				return m, nil
			}
			entries, _ := storage.DeleteEntry(m.choices[m.cursor])
			m.choices = entries
			m.moveCursor(false)
			return m, nil

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case digitKey.MatchString(msg.String()):
			tryJumpToEntry(&m, msg.String())
			return m, nil
		}
	}
	return m, nil
}

func inputUpdate(msg tea.Msg, m Model, iw inputView) (Model, tea.Cmd) {
	var cmd tea.Cmd
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(msg, m.keys.Confirm):
			iw.confirmAction(&m)
			m.input.SetValue("")
			m.setState(MainState)
			return m, cmd

		case key.Matches(msg, m.keys.Back):
			m.setState(MainState)
			return m, nil
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func choiceExists(m Model) bool {
	return m.cursor >= 0 && m.cursor < len(m.choices)
}

func tryJumpToEntry(m *Model, digitStr string) error {
	digit, err := strconv.Atoi(digitStr)
	if err != nil {
		return err
	}

	if digit > 0 && digit <= len(m.choices) {
		m.cursor = digit - 1
	}

	return nil
}

func moveEntryToTop(m *Model, entry storage.Entry) {
	for i, e := range m.choices {
		if e.Name == entry.Name {
			copy(m.choices[i:], m.choices[i+1:])
			m.choices = m.choices[:len(m.choices)-1]
			break
		}
	}
	m.choices = append([]storage.Entry{entry}, m.choices...)
	storage.SaveEntries(m.choices)
}
