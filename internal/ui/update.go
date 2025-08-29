package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fredrikzkl/dash/internal/storage"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case MAIN_STATE:
		return mainUpdate(msg, m)
	case ADD_STATE, COMMAND_STATE:
		return inputUpdate(msg, m, newEntryInputView)
	}

	return m, nil
}

func mainUpdate(msg tea.Msg, m Model) (Model, tea.Cmd) {
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
			m.setState(ADD_STATE)

			pwd, err := getPwd()
			if err == nil {
				presetInput(&m, pwd)
			}
			return m, nil

		case key.Matches(msg, m.keys.Command):
			m.setState(COMMAND_STATE)

			if !choiceExists(m) {
				return m, nil
			}

			currentCmd := m.choices[m.cursor].Command
			if currentCmd != "" {
				presetInput(&m, currentCmd)
			}

			return m, nil

		case key.Matches(msg, m.keys.Back):
			m.setState(MAIN_STATE)
			return m, nil

		case key.Matches(msg, m.keys.Delete):
			if len(m.choices) == 0 {
				return m, nil
			}
			entries, _ := storage.DeleteEntry(m.choices[m.cursor])
			m.choices = entries
			m.cursor = 0
			return m, nil

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func inputUpdate(msg tea.Msg, m Model, iw inputView) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Confirm):
			iw.confirmAction(&m, m.input.Value())
			m.input.SetValue("")
			m.setState(MAIN_STATE)
			return m, cmd

		case key.Matches(msg, m.keys.Back):
			m.setState(MAIN_STATE)
			return m, nil
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func choiceExists(m Model) bool {
	return m.cursor >= 0 && m.cursor < len(m.choices)
}
