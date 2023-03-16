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
	unfocusedStyle = lipgloss.NewStyle().Padding(1)
	focusedStyle   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
)

type Model struct {
	name     string
	w, h     int
	children map[shared.ID]*Model

	focus.Handler
}

func New(hasFocus bool) *Model {
	return &Model{
		"A",
		20, 10,
		nil,
		focus.NewNode(hasFocus),
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

	textStyle := unfocusedStyle
	if m.HasFocus() {
		textStyle = focusedStyle
	}
	w := m.w - textStyle.GetHorizontalFrameSize()
	h := m.h - textStyle.GetVerticalFrameSize()
	textStyle = textStyle.Copy().Width(w).Height(h)
	textContent := textStyle.Render(m.name)

	vp.SetContent(textContent)
	return vp.View()
}

func (m *Model) Resize(w, h int) component.Resizable {
	m.w, m.h = w, h
	return m
}
