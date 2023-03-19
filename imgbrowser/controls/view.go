package controls

import "github.com/charmbracelet/lipgloss"

var (
	activeStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#aaaaaa")).
			Foreground(lipgloss.Color("#aaaaaa"))
	focusStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#ffffff")).
			Foreground(lipgloss.Color("#ffffff"))
	normalStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#555555")).
			Foreground(lipgloss.Color("#555555"))
)

func (m Model) drawTitle() string {
	title := "▛▜▐▜▐▀▐▝▜▐▜▐ ▐▝▜▐▛▐█\n" +
		"▛▜▐▐▗▟▐▐▄▐▜▐▄▐▐▄▐▄▐▐"
	return lipgloss.NewStyle().Width(30).AlignHorizontal(lipgloss.Center).Padding(1, 1, 0, 1).Render(title)
}

func (m Model) drawButtons() string {
	buttons := make([]string, len(stateOrder))
	for i, state := range stateOrder {
		style := normalStyle
		if state == m.active {
			style = activeStyle
		} else if state == m.focus {
			style = focusStyle
		}
		buttons[i] = style.Copy().Width(8).Align(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}
