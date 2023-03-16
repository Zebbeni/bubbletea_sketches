package box

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/focus"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/shared"
)

var (
	unfocusedStyle = lipgloss.NewStyle()
	focusedStyle   = lipgloss.NewStyle()
)

type Model struct {
	name     string
	w, h     int
	children map[shared.ID]*Model

	focus.Handler
}

func New() *Model {
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

func (m *Model) Update(msg tea.Msg) (component.Resizable, tea.Cmd) {
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
	vp := viewport.New(m.w, m.h)

	content := lipgloss.NewStyle().Width(m.w).Height(m.h).Render(m.name)
	vp.SetContent(content)
	return vp.View()
	//if m.HasFocus() {
	//	return focusedStyle.Copy().Width(m.w).Height(m.h).Render(m.name)
	//}
	//return unfocusedStyle.Copy().Width(m.w).Height(m.h).Render(m.name)
}

func (m *Model) Resize(w, h int) component.Resizable {
	m.w, m.h = w, h
	return m
}
