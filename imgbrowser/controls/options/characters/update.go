package characters

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var navMap = map[Direction]map[State]State{
	Right: {AsciiButton: UnicodeButton,
		AsciiAzButton:      AsciiNumButton,
		AsciiNumButton:     AsciiSpecButton,
		AsciiSpecButton:    AsciiAllButton,
		UnicodeFullButton:  UnicodeHalfButton,
		UnicodeHalfButton:  UnicodeQuartButton,
		UnicodeQuartButton: UnicodeShadeButton,
		OneColor:           TwoColor,
	},
	Left: {UnicodeButton: AsciiButton,
		AsciiAllButton:     AsciiSpecButton,
		AsciiSpecButton:    AsciiNumButton,
		AsciiNumButton:     AsciiAzButton,
		UnicodeShadeButton: UnicodeQuartButton,
		UnicodeQuartButton: UnicodeHalfButton,
		UnicodeHalfButton:  UnicodeFullButton,
		TwoColor:           OneColor,
	},
	Up: {
		AsciiButton:        OneColor,
		UnicodeButton:      TwoColor,
		AsciiAzButton:      AsciiButton,
		AsciiNumButton:     AsciiButton,
		AsciiSpecButton:    AsciiButton,
		AsciiAllButton:     AsciiButton,
		UnicodeFullButton:  UnicodeButton,
		UnicodeHalfButton:  UnicodeButton,
		UnicodeQuartButton: UnicodeButton,
		UnicodeShadeButton: UnicodeButton,
	},
	Down: {
		OneColor: AsciiButton,
		TwoColor: UnicodeButton,

		AsciiButton:   AsciiAzButton,
		UnicodeButton: UnicodeShadeButton,
	},
}

var (
	asciiCharModeMap = map[State]CharMode{
		AsciiAzButton:   AzAscii,
		AsciiNumButton:  NumAscii,
		AsciiSpecButton: SpecAscii,
		AsciiAllButton:  AllAscii,
	}
	unicodeCharModeMap = map[State]CharMode{
		UnicodeFullButton:  FullBlockUnicode,
		UnicodeHalfButton:  HalfBlockUnicode,
		UnicodeQuartButton: QuartBlockUnicode,
		UnicodeShadeButton: ShadeBlockUnicode,
	}
)

func (m Model) handleEsc() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	m.active = m.focus

	switch m.active {
	case AsciiButton:
		m.mode = Ascii
	case UnicodeButton:
		m.mode = Unicode
	case OneColor, TwoColor:
		m.useFgBg = m.active
	default:
		switch m.mode {
		case Ascii:
			if charMode, ok := asciiCharModeMap[m.active]; ok {
				m.asciiMode = charMode
			}
		case Unicode:
			if charMode, ok := unicodeCharModeMap[m.active]; ok {
				m.unicodeMode = charMode
			}
		}
	}
	return m, io.StartRenderCmd
}

func (m Model) handleNav(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Right):
		if next, hasNext := navMap[Right][m.focus]; hasNext {
			return m.setFocus(next)
		}
	case key.Matches(msg, io.KeyMap.Left):
		if next, hasNext := navMap[Left][m.focus]; hasNext {
			return m.setFocus(next)
		}
	case key.Matches(msg, io.KeyMap.Up):
		if next, hasNext := navMap[Up][m.focus]; hasNext {
			return m.setFocus(next)
		}
	case key.Matches(msg, io.KeyMap.Down):
		if next, hasNext := navMap[Down][m.focus]; hasNext {
			return m.setFocus(next)
		}
	}
	return m, cmd
}

func (m Model) setFocus(focus State) (Model, tea.Cmd) {
	m.focus = focus
	if m.focus == AsciiButton {
		m.charButtons = Ascii
	} else if m.focus == UnicodeButton {
		m.charButtons = Unicode
	}
	return m, nil
}
