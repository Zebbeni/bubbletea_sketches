package colors

import "github.com/charmbracelet/lipgloss"

var (
	stateOrder = []State{TrueColor, Adaptive, Paletted}
	stateNames = map[State]string{
		TrueColor: "True Color",
		Adaptive:  "Adaptive",
		Paletted:  "Palette",
	}

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

func (m Model) drawButtons() string {
	buttons := make([]string, len(stateOrder))
	for i, state := range stateOrder {
		style := normalStyle
		if state == m.active {
			style = activeStyle
		} else if state == m.focus {
			style = focusStyle
		}
		buttons[i] = style.Copy().Width(10).Align(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}
