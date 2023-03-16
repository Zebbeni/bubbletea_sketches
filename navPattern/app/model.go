package app

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/env"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/gui/box"
)

type Model struct {
	w, h int

	content box.Model
}

func New() Model {
	content := box.New(true)
	return Model{
		content: content,
	}
}

func (m Model) Init() tea.Cmd {
	if env.PollForSizeChange {
		return pollForSizeChange
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
