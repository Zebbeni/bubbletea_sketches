package tagList

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
)

type Model struct {
	component.Model
}

func New() Model {
	m := Model{
		component.New(),
	}
	m.FocusInternal = true
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "tagList"
}
