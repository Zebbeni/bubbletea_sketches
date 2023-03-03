package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app/process"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/viewer"
)

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)
	switch m.state {
	case Main:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			selectedItem := m.menu.SelectedItem().(item)
			m.state = selectedItem.state
		default:
			m.menu, cmds[0] = m.menu.Update(msg)
		}
	case Browser:
		m.browser, cmds[0] = m.browser.Update(msg)
		if m.browser.DidUpdate {
			m.browser.DidUpdate = false
			m.viewer.WaitingOnRender = true
			cmds[1] = m.newRenderCmd
		}
		if m.browser.ShouldClose {
			m.state = Main
		}
	case Settings:
		m.settings, cmds[0] = m.settings.Update(msg)
		if m.settings.DidUpdate {
			m.settings.DidUpdate = false
			m.viewer.WaitingOnRender = true
			cmds[1] = m.newRenderCmd
		}
		if m.settings.ShouldClose {
			m.state = Main
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) handleBackMsg() (Model, tea.Cmd) {
	m.state = Main
	return m, nil
}

func (m Model) handleRenderMsg(msg viewer.RenderMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewer, cmd = m.viewer.Update(msg)
	return m, cmd
}

func (m Model) newRenderCmd() tea.Msg {
	imgString := process.RenderImageFile(m.browser.File)
	// kick off a render job
	return viewer.RenderMsg{FilePath: m.browser.File, ImgString: imgString}
}
