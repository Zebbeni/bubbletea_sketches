package container

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/gui"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/gui/box"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/nav"
)

const (
	BoxA gui.ID = iota
	BoxB
)

type Model struct {
	w, h int

	components map[gui.ID]gui.Component

	nav.FocusHandler
}

func New() Model {
	left := map[gui.ID]gui.ID{BoxA: BoxB}
	right := map[gui.ID]gui.ID{BoxB: BoxA}
	navMap := nav.NewMap(left, right, nil, nil)
	handler := nav.NewContainer(true, true, BoxA, navMap)
	children := map[gui.ID]gui.Component{
		BoxA: box.New(true),
		BoxB: box.New(false),
	}

	return Model{
		components:   children,
		FocusHandler: &handler,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	boxA := m.components[BoxA].View()
	boxB := m.components[BoxB].View()
	return lipgloss.JoinHorizontal(lipgloss.Left, boxA, boxB)
}
