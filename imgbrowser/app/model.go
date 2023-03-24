package app

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/env"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/viewer"
)

type State int

const (
	Main State = iota
	Browser
	Settings
)

type Model struct {
	state State

	controls controls.Model
	viewer   viewer.Model

	w, h int
}

func New() Model {
	return Model{
		state:    Main,
		controls: controls.New(),
		viewer:   viewer.New(),
		w:        100,
		h:        100,
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
	case checkSizeMsg:
		return m.handleCheckSizeMsg()
	case tea.WindowSizeMsg:
		return m.handleSizeMsg(msg)
	case io.StartRenderMsg:
		return m.handleStartRenderMsg()
	case io.FinishRenderMsg:
		return m.handleFinishRenderMsg(msg)
	}
	return m.handleControlsUpdate(msg)
}

func (m Model) View() string {
	lViewport := viewport.New(m.leftPanelWidth(), m.leftPanelHeight())

	leftContent := m.controls.View()
	lViewport.SetContent(lipgloss.NewStyle().Width(m.leftPanelWidth()).Height(m.leftPanelHeight()).Render(leftContent))
	leftPanel := lViewport.View()

	viewer := m.viewer.View()
	rViewport := viewport.New(m.rPanelWidth(), m.rPanelHeight())

	vpRightStyle := lipgloss.NewStyle().Align(lipgloss.Center).AlignVertical(lipgloss.Center)
	rightContent := vpRightStyle.Copy().Width(m.rPanelWidth()).Height(m.rPanelHeight()).Render(viewer)
	rViewport.SetContent(rightContent)
	rightPanel := rViewport.View()

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

	vp := viewport.New(m.w, m.h)
	vp.SetContent(content)
	//vp.Style = lipgloss.NewStyle().Width(m.w).Height(m.h).BorderStyle(lipgloss.RoundedBorder())
	vp.Style = lipgloss.NewStyle().Width(m.w).Height(m.h)
	return vp.View()
}
