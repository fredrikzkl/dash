package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

const (
	vp_width  = 78
	vp_height = 20
)

func newViewport() (viewport.Model, error) {
	vp := viewport.New(vp_width, vp_height)

	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingLeft(4).
		PaddingRight(4)

	return vp, nil
}

func render(viewport *viewport.Model, content string) error {
	viewport.SetContent(content)
	return nil
}
