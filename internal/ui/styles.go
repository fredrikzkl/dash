package ui

import "github.com/charmbracelet/lipgloss"

const (
	vp_width  = 50
	vp_height = 20
)

var standardViewportStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("1")).
	Padding(1).
	PaddingLeft(2)

var headerStyle = lipgloss.NewStyle().
	Italic(true).
	MarginBottom(1)

var hoverStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("4"))

var cmdToggledStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("5"))

var cmdText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#efefef"))
