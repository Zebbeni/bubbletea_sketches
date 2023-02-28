package viewer

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type FileMsg string
type RenderMsg struct {
	Path     string
	Rendered string
}

type Model struct {
	filePath          string
	renderString      string
	isWaitingOnRender bool
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FileMsg:
		return m.handleFileMsg(msg)
	case RenderMsg:
		return m.handleRenderMsg(msg)
	}
	return m, nil
}

func (m Model) handleFileMsg(msg FileMsg) (Model, tea.Cmd) {
	if m.filePath == string(msg) {
		return m, nil
	}
	// kick off a render process
	m.filePath = string(msg)
	m.isWaitingOnRender = true

	return m, m.msgWhenRendered
}

func (m Model) handleRenderMsg(msg RenderMsg) (Model, tea.Cmd) {
	if msg.Path != m.filePath {
		return m, nil
	}
	m.isWaitingOnRender = false
	m.renderString = msg.Rendered
	return m, nil
}

func (m Model) View() string {
	if m.isWaitingOnRender {
		return "rendering..."
	}
	return "done: " + m.renderString
}

func (m Model) msgWhenRendered() tea.Msg {
	time.Sleep(time.Second)
	// kick off a render job
	return RenderMsg{Path: m.filePath, Rendered: m.filePath}
}
