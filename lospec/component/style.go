package component

import "github.com/charmbracelet/lipgloss"

var (
	Unfocused = lipgloss.NewStyle().Padding(1).Width(50)
	Focused   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Width(50)
	Active    = lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).Width(50)
)
