package box

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component"
)

var (
	unfocusedStyle = lipgloss.NewStyle().Padding(1)
	focusedStyle   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
)

type Model struct {
	name string
	w, h int
	component.FocusHandler
}

func New(n string, w, h int) Model {
	return Model{
		n,
		w, h,
		component.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}

func (m Model) View() string {
	return unfocusedStyle.Copy().Width(m.w).Height(m.h).Render(m.name)
}
