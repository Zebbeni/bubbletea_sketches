package box

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/focus"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/shared"
)

var (
	unfocusedStyle = lipgloss.NewStyle().Padding(1)
	focusedStyle   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
)

type Model struct {
	name     string
	w, h     int
	children map[shared.ID]*Model

	focus.Handler
}

func New() *Model {
	// create left
	// create right
	// create map
	// create Container

	return &Model{
		"A",
		20, 10,
		nil,
		focus.NewNode(true),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			m.HandleNav(msg)
		}
	}
	return m, nil
}

func (m *Model) View() string {
	//return m.name
	if m.HasFocus() {
		return focusedStyle.Copy().Width(m.w).Height(m.h).Render(m.name)
	}
	return unfocusedStyle.Copy().Width(m.w).Height(m.h).Render(m.name)
}
