package colors

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/makeworld-the-better-one/dither/v2"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/adaptive"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette"
)

type State int

// None consists of a few different components that are shown or hidden
// depending on which toggles have been set on / off. The Model state indicates
// which component is currently focused. From top to bottom the components are:

// 1) Limited (on/off)
// 2) Palette (Name) (if Limited) -> [Enter] displays Palette menu
// 3) Dithering (on/off) (if Limited)
// 4) Serpentine (on/off) (if Dithering)
// 5) Matrix (Name) (if Dithering) -> [Enter] displays to Matrix menu

// These can all be part of a single list, but we need to onSelect the list items

const (
	None State = iota
	TrueColor
	Adaptive
	Paletted
)

type Model struct {
	active State // the panel taking input
	focus  State // the focused button
	mode   State // the mode currently in use for rendering

	Adaptive adaptive.Model
	Palette  palette.Model

	ShouldClose bool
}

func New() Model {
	m := Model{
		active:      None,
		focus:       TrueColor,
		Adaptive:    adaptive.New(),
		Palette:     palette.New(),
		ShouldClose: false,
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.active {
	case Adaptive:
		return m.handleAdaptiveUpdate(msg)
	case Paletted:
		return m.handlePaletteUpdate(msg)
	}
	return m.handleMenuUpdate(msg)
}

func (m Model) View() string {
	title := m.drawTitle()
	buttons := m.drawButtons()
	var controls string
	switch m.active {
	case Adaptive:
		controls = m.Adaptive.View()
	case Paletted:
		controls = m.Palette.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top, title, buttons, controls)
}

func (m Model) IsLimited() bool {
	return true
}

func (m Model) IsDithered() bool {
	return true
}

func (m Model) IsSerpentine() bool {
	return true
}

func (m Model) Matrix() dither.ErrorDiffusionMatrix {
	return dither.FloydSteinberg
}
