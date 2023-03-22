package characters

import "github.com/charmbracelet/lipgloss"

var (
	stateOrder = []State{Ascii, Blocks}
	stateNames = map[State]string{
		Ascii:  "Ascii",
		Blocks: "Blocks",
	}

	activeColor = lipgloss.Color("#aaaaaa")
	focusColor  = lipgloss.Color("#ffffff")
	normalColor = lipgloss.Color("#555555")
	titleStyle  = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))
)

func (m Model) drawButtons() string {
	buttons := make([]string, len(stateOrder))
	for i, state := range stateOrder {
		styleColor := normalColor
		if m.IsActive {
			if state == m.focus {
				styleColor = focusColor
			} else if state == m.active {
				styleColor = activeColor
			}
		}
		style := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(styleColor).
			Foreground(styleColor)
		buttons[i] = style.Copy().Width(11).AlignHorizontal(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}
