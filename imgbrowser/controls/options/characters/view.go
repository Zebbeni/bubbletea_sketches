package characters

import "github.com/charmbracelet/lipgloss"

var (
	stateOrder = []State{Ascii, Blocks}
	stateNames = map[State]string{
		Ascii:  "Ascii",
		Blocks: "Blocks",
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
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))
)

func (m Model) drawButtons() string {
	buttons := make([]string, len(stateOrder))
	for i, state := range stateOrder {
		style := normalStyle
		if m.IsActive {
			if state == m.active {
				style = activeStyle
			} else if state == m.focus {
				style = focusStyle
			}
		}
		buttons[i] = style.Copy().Width(11).AlignHorizontal(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}
