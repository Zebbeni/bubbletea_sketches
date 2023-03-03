package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app/process"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleMainUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			selectedItem := m.menu.SelectedItem().(item)
			m.state = selectedItem.state
		}
	}
	m.menu, cmd = m.menu.Update(msg)
	return m, cmd
}

func (m Model) handleBrowserUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.browser, cmd = m.browser.Update(msg)
	if m.browser.ShouldClose {
		m.browser.ShouldClose = false
		m.state = Main
	}
	return m, cmd
}

func (m Model) handleSettingsUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.settings, cmd = m.settings.Update(msg)
	if m.settings.ShouldClose {
		m.settings.ShouldClose = false
		m.state = Main
	}
	return m, cmd
}

func (m Model) handleStartRenderMsg() (Model, tea.Cmd) {
	m.viewer.WaitingOnRender = true
	return m, m.processRenderCmd
}

func (m Model) handleFinishRenderMsg(msg io.FinishRenderMsg) (Model, tea.Cmd) {
	// cut out early if the finished render is for an previously selected image
	if msg.FilePath != m.browser.ActiveFile {
		return m, nil
	}

	var cmd tea.Cmd
	m.viewer, cmd = m.viewer.Update(msg)
	return m, cmd
}

func (m Model) processRenderCmd() tea.Msg {
	imgString := process.RenderImageFile(m.browser.ActiveFile)
	return io.FinishRenderMsg{FilePath: m.browser.ActiveFile, ImgString: imgString}
}
