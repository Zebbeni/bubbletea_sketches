package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component/box"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/component/stretch"
)

type Model struct {
	content stretch.Model
}

func New() Model {
	s := stretch.New(box.New())
	return Model{
		content: s,
	}
}

func (m Model) Init() tea.Cmd {
	return m.content.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.content, cmd = m.content.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.content.View()
}
