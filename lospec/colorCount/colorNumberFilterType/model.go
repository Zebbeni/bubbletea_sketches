package colorNumberFilterType

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
)

type FilterType int

const (
	Any FilterType = iota
	Max
	Min
	Exact
)

type Model struct {
	Selected FilterType
	items    []option

	component.Model
}

func New() Model {
	m := Model{
		Any,
		buildOptions(),
		component.New(),
	}
	m.FocusInternal = true
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			m, cmd = m.handleNavigation(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {
	options := make([]string, len(m.items))
	for i, o := range m.items {
		options[i] = lipgloss.NewStyle().Padding(1).Render(o.name)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, options...)
}
