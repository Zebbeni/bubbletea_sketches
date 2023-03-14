package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component"
)

type Model struct {
	// Hmm does it really make sense to have all children be navigable?
	childMap map[component.Focus]component.Focusable

	*component.FocusHandler
}

func New() Model {
	return Model{
		childMap: component.New(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	}
	return m, nil
}

func (m *Model) View() string {
	style := lipgloss.NewStyle()
}
