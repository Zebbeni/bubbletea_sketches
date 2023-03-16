package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/box"
)

type Model struct {
	box *box.Model
}

func New() *Model {
	return &Model{
		box: box.New(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.box.Update(msg)
	return m, nil
}

func (m *Model) View() string {
	return m.box.View()
}
