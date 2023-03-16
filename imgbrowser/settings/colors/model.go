package colors

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/makeworld-the-better-one/dither/v2"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/menu"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette"
)

type State int

// Colors consists of a few different components that are shown or hidden
// depending on which toggles have been set on / off. The Model state indicates
// which shared is currently focused. From top to bottom the components are:

// 1) Limited (on/off)
// 2) Palette (Name) (if Limited) -> [Enter] displays Palette menu
// 3) Dithering (on/off) (if Limited)
// 4) Serpentine (on/off) (if Dithering)
// 5) Matrix (Name) (if Dithering) -> [Enter] displays to Matrix menu

// These can all be part of a single list, but we need to onSelect the list items

const (
	Menu State = iota
	Palette
)

type Model struct {
	state State

	menu list.Model

	IsLimited    bool
	IsDithered   bool
	IsSerpentine bool

	Palette  palette.Model
	Ditherer dither.Ditherer
	Matrix   Matrix

	showPaletteMenu bool
	showMatrixMenu  bool

	ShouldClose bool
}

func New() Model {
	m := Model{
		state:           Menu,
		IsLimited:       true,
		IsDithered:      true,
		IsSerpentine:    true,
		Palette:         palette.New(),
		Matrix:          Matrix{Name: "FloydSteinberg", Method: dither.FloydSteinberg},
		showPaletteMenu: false,
		showMatrixMenu:  false,
		ShouldClose:     false,
	}
	m.menu = menu.New(buildMenuItems(m))
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case Menu:
		m, cmd = m.handleMenuUpdate(msg)
	case Palette:
		m, cmd = m.handlePaletteUpdate(msg)
	}
	m.menu.SetItems(buildMenuItems(m))
	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case Menu:
		return m.menu.View()
	case Palette:
		return m.Palette.View()
	}
	return ""
}

func (m Model) paletteName() string {
	name, _ := m.Palette.GetCurrent()
	return name
}

func (m Model) matrixName() string {
	return m.Matrix.Name
}
