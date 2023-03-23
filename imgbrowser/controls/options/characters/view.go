package characters

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	stateOrder         = []State{AsciiButton, UnicodeButton}
	asciiButtonOrder   = []State{AsciiAzButton, AsciiNumButton, AsciiSpecButton, AsciiAllButton}
	unicodeButtonOrder = []State{UnicodeFullButton, UnicodeHalfButton, UnicodeQuartButton, UnicodeShadeButton}
	colorsButtonsOrder = []State{OneColor, TwoColor}

	stateNames = map[State]string{
		AsciiButton:        "Ascii",
		UnicodeButton:      "Unicode",
		AsciiAzButton:      "AZ",
		AsciiNumButton:     "0-9",
		AsciiSpecButton:    "!$",
		AsciiAllButton:     "All",
		UnicodeFullButton:  "1",
		UnicodeHalfButton:  "1/2",
		UnicodeQuartButton: "1/4",
		UnicodeShadeButton: "░▒▓",
		OneColor:           "1 Color",
		TwoColor:           "2 Color",
	}

	activeColor = lipgloss.Color("#aaaaaa")
	focusColor  = lipgloss.Color("#ffffff")
	normalColor = lipgloss.Color("#555555")
	titleStyle  = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))
)

func (m Model) drawModeButtons() string {
	buttons := make([]string, len(stateOrder))
	for i, state := range stateOrder {
		styleColor := normalColor
		if m.IsActive && state == m.focus {
			styleColor = focusColor
		} else if state == m.active {
			styleColor = activeColor
		}
		style := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(styleColor).
			Foreground(styleColor)
		buttons[i] = style.Copy().Width(11).AlignHorizontal(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}

func (m Model) drawCharButtons() string {
	var buttonOrder []State
	switch m.charButtons {
	case Ascii:
		buttonOrder = asciiButtonOrder
	case Unicode:
		buttonOrder = unicodeButtonOrder
	}
	buttons := make([]string, len(buttonOrder))
	for i, state := range buttonOrder {
		styleColor := normalColor
		if m.IsActive && state == m.focus {
			styleColor = focusColor
		} else if state == m.active {
			styleColor = activeColor
		}
		style := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(styleColor).
			Foreground(styleColor)

		// quick dirty stuff to make buttons fit nicely. Let's do this in a smarter / cleaner way later
		if m.charButtons == Unicode {
			buttons[i] = style.Copy().Padding(0, 1, 0, 1).Render(stateNames[state])
		} else {
			buttons[i] = style.Copy().Padding(0, 1, 0, 1).Render(stateNames[state])
		}
	}
	content := lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
	return lipgloss.NewStyle().Width(27).AlignHorizontal(lipgloss.Center).Render(content)
}

func (m Model) drawColorsButtons() string {
	buttons := make([]string, len(colorsButtonsOrder))
	for i, state := range colorsButtonsOrder {
		styleColor := normalColor
		if m.IsActive && state == m.focus {
			styleColor = focusColor
		} else if state == m.useFgBg {
			styleColor = activeColor
		}
		style := lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(styleColor).
			Foreground(styleColor)
		buttons[i] = style.Copy().Width(11).AlignHorizontal(lipgloss.Center).Render(stateNames[state])
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, buttons...)
}
