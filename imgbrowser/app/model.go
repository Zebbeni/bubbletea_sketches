package app

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/browser"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/env"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/menu"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/viewer"
)

type Model struct {
	menu    menu.Model
	browser browser.Model
	viewer  viewer.Model
	w, h    int
}

func New() Model {
	return Model{
		menu:    menu.New(),
		browser: browser.New(),
		viewer:  viewer.New(),
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleSizeMsg(msg)
	case checkSizeMsg:
		return m.handleCheckSizeMsg()
	case io.BackMsg:
		return m.handleBackMsg()
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case viewer.RenderMsg:
		return m.handleRenderMsg(msg)
	}
	return m, nil
}

func (m Model) View() string {
	vpLeft := viewport.New(m.leftPanelWidth(), m.leftPanelHeight())
	var leftContent string
	switch m.menu.State {
	case menu.MainMenu:
		leftContent = m.menu.View()
	case menu.FileMenu:
		leftContent = m.browser.View()
	}
	vpLeft.SetContent(lipgloss.NewStyle().
		Width(m.leftPanelWidth()).
		Height(m.leftPanelHeight()).
		Render(leftContent))
	panelLeft := vpLeft.View()

	viewer := m.viewer.View()
	vpRight := viewport.New(m.rightPanelWidth(), m.rightPanelHeight())
	vpRight.SetContent(lipgloss.NewStyle().
		Width(m.rightPanelWidth()).
		Height(m.rightPanelHeight()).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(viewer))
	panelRight := vpRight.View()

	content := lipgloss.JoinHorizontal(lipgloss.Top, panelLeft, panelRight)

	vp := viewport.New(m.w, m.h)
	vp.SetContent(content)
	vp.Style = lipgloss.NewStyle().Width(m.w).Height(m.h).BorderStyle(lipgloss.RoundedBorder())
	return vp.View()
}
