package characters

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Mode int

const (
	Ascii Mode = iota
	Unicode
)

type CharMode int

const (
	AzAscii CharMode = iota
	NumAscii
	SpecAscii
	AllAscii
	FullBlockUnicode
	HalfBlockUnicode
	QuartBlockUnicode
	ShadeBlockUnicode
)

type State int

const (
	AsciiButton State = iota
	UnicodeButton
	AsciiAzButton
	AsciiNumButton
	AsciiSpecButton
	AsciiAllButton
	UnicodeFullButton
	UnicodeHalfButton
	UnicodeQuartButton
	UnicodeShadeButton
	OneColor
	TwoColor
)

type Model struct {
	focus         State
	active        State
	mode          Mode
	charButtons   Mode
	unicodeMode   CharMode
	asciiMode     CharMode
	useFgBg       State
	ShouldClose   bool
	ShouldUnfocus bool
	IsActive      bool
}

func New() Model {
	return Model{
		focus:         AsciiButton,
		active:        AsciiButton,
		mode:          Ascii,
		charButtons:   Ascii,
		asciiMode:     NumAscii,
		unicodeMode:   NumAscii,
		useFgBg:       OneColor,
		ShouldClose:   false,
		ShouldUnfocus: false,
		IsActive:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
		}
	}
	return m, nil
}

func (m Model) View() string {
	colorsButtons := m.drawColorsButtons()
	modeButtons := m.drawModeButtons()
	charButtons := m.drawCharButtons()
	return lipgloss.JoinVertical(lipgloss.Top, colorsButtons, modeButtons, charButtons)
}

func (m Model) Selected() (Mode, CharMode, State) {
	charMode := m.asciiMode
	if m.mode == Unicode {
		charMode = m.unicodeMode
	}
	return m.mode, charMode, m.useFgBg
}
