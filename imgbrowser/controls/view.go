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
	//title1Runes := "▛▜▐▜▐▀▐▝▜▐▜▐ ▐▝▜▐▛▐▀▌"
	//title2Runes := "▛▜▐▐▗▟▐▐▄▐▜▐▄▐▐▄▐▄▐▜"
	title1Runes := []rune{' ', '▛', '▜', '▐', '▜', '▐', '▀', '▐', '▝', '▜', '▐', '▜', '▐', ' ', '▐', '▝', '▜', '▐', '▛', '▐', '▀', '▌'}
	title2Runes := []rune{'▛', '▜', '▐', '▐', '▗', '▟', '▐', '▐', '▄', '▐', '▜', '▐', '▄', '▐', '▐', '▄', '▐', '▄', '▐', '▜'}
	title1 := lipgloss.JoinHorizontal(lipgloss.Left, string(title1Runes))
	title2 := lipgloss.JoinHorizontal(lipgloss.Left, string(title2Runes))
	title := lipgloss.JoinVertical(lipgloss.Left, title1, title2)
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
		buttons[i] = style.Copy().Width(7).AlignHorizontal(lipgloss.Center).Render(stateNames[state])
	}
	buttonRow := lipgloss.JoinHorizontal(lipgloss.Left, buttons...)

	return lipgloss.NewStyle().Padding(0, 1, 0, 1).Render(buttonRow)
}
