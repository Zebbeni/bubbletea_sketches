package app

import (
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/browser"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/env"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	browser browser.Model
	w, h    int
}

func New() Model {
	return Model{
		browser: browser.New(),
		w:       100,
		h:       100,
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleSizeMsg(msg)
	case checkSizeMsg:
		return m.handleCheckSizeMsg()
	case tea.KeyMsg:
		m.browser, cmd = m.browser.Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	vp := viewport.New(m.w, m.h)

	browser := m.browser.View()
	vp.SetContent(lipgloss.NewStyle().Width(m.w).Height(m.h).Render(browser))

	vp.Style = lipgloss.NewStyle().Width(m.w).Height(m.h).BorderStyle(lipgloss.RoundedBorder())
	return vp.View()
}
