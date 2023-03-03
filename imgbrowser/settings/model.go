package settings

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/interpolation"
)

type Model struct {
	state State
	menu  list.Model

	Interpolation interpolation.Model

	ShouldClose bool
}

func New() Model {
	return Model{
		state:         Main,
		menu:          newMenu(),
		Interpolation: interpolation.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case Main:
		m, cmd = m.handleMainUpdate(msg)
	case Interpolation:
		m, cmd = m.handleInterpolationUpdate(msg)
	}
	m = m.handleCloseFlags()
	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case Main:
		return m.menu.View()
	case Interpolation:
		return m.Interpolation.View()
		//case Palette:
		//case Dithering:
		//case Characters:
	}
	return m.menu.View()
}
