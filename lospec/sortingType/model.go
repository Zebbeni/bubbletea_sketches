package sortingType

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
)

type SelectedSortingType int

const (
	Default SelectedSortingType = iota
	Alphabetical
	Downloads
	Newest
)

type Model struct {
	Selected SelectedSortingType
	items    []option

	component.Model
}

func New() Model {
	return Model{
		Default,
		buildOptions(),
		component.New(),
	}
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
	content := lipgloss.JoinHorizontal(lipgloss.Top, options...)
	return m.GetStyle().Render(content)
}
