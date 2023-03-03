package settings

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/sampling"
)

type Model struct {
	state State
	menu  list.Model

	Colors   colors.Model
	Sampling sampling.Model

	ShouldClose bool
}

func New() Model {
	return Model{
		state:    Main,
		menu:     newMenu(),
		Colors:   colors.New(),
		Sampling: sampling.New(),
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
	case Colors:
		m, cmd = m.handleColorsUpdate(msg)
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
	case Colors:
		return m.Colors.View()
	case Interpolation:
		return m.Sampling.View()
		//case Palette:
		//case Colors:
		//case Characters:
	}
	return m.menu.View()
}
