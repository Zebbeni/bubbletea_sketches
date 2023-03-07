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
		state:    Menu,
		menu:     newMenu(),
		Colors:   colors.New(),
		Sampling: sampling.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.state {
	case Menu:
		return m.handleMenuUpdate(msg)
	case Colors:
		return m.handleColorsUpdate(msg)
	case Sampling:
		return m.handleSamplingUpdate(msg)
	}
	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case Menu:
		return m.menu.View()
	case Colors:
		return m.Colors.View()
	case Sampling:
		return m.Sampling.View()
		//case Limited:
		//case Colors:
		//case Characters:
	}
	return m.menu.View()
}
