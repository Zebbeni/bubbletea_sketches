package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/browser"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/env"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
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

	menu     list.Model
	browser  browser.Model
	settings settings.Model
	viewer   viewer.Model

	w, h int
}

func New() Model {
	return Model{
		state:    Main,
		menu:     newMenu(),
		browser:  browser.New(),
		settings: settings.New(),
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
	default:
		switch m.state {
		case Main:
			return m.handleMainUpdate(msg)
		case Browser:
			return m.handleBrowserUpdate(msg)
		case Settings:
			return m.handleSettingsUpdate(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	vpLeft := viewport.New(m.leftPanelWidth(), m.leftPanelHeight())
	var leftContent string
	switch m.state {
	case Main:
		leftContent = m.menu.View()
	case Browser:
		leftContent = m.browser.View()
	case Settings:
		leftContent = m.settings.View()
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
