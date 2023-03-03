package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/interpolation"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/menu"
)

type Model struct {
	Interpolation interpolation.Model
	state         menu.State

	list menu.Model

	DidUpdate   bool
	ShouldClose bool
}

func New() Model {
	return Model{
		state: menu.Main,
		list:  menu.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc(), nil
			//case key.Matches(msg, io.KeyMap.Enter):
			//	return m.handleEnter()
			//default:
			//	return m.handleKey(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "View Settings"
}
