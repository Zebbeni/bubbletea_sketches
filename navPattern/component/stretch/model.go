package stretch

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/env"
)

type State int

type Model struct {
	content component.Resizable

	w, h int
}

func New(content component.Resizable) Model {
	return Model{
		content: content,
	}
}

func (m Model) Init() tea.Cmd {
	// This initiates the polling cycle for window size updates
	// but shouldn't be necessary on non-Windows computers.
	if env.PollForSizeChange {
		return pollForSizeChange
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case checkSizeMsg:
		return m.handleCheckSizeMsg()
	case tea.WindowSizeMsg:
		return m.handleSizeMsg(msg)
	}
	m.content, cmd = m.content.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	vp := viewport.New(m.w, m.h)
	vp.SetContent(m.content.View())
	return vp.View()
}
