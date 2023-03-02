package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app/menu"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app/process"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/viewer"
)

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)
	switch m.menu.State {
	case menu.MainMenu:
		m.menu, cmds[0] = m.menu.Update(msg)
	case menu.FileMenu:
		m.browser, cmds[0] = m.browser.Update(msg)
		if m.browser.DidUpdate {
			m.browser.DidUpdate = false
			m.viewer.WaitingOnRender = true
			cmds[1] = m.newRenderCmd
		}
	case menu.SettingsMenu:
		// it'd be nice to have settings update a flag if 'dirty'
		m.settings, cmds[0] = m.settings.Update(msg)
		if m.settings.DidUpdate {
			m.settings.DidUpdate = false
			m.viewer.WaitingOnRender = true
			cmds[1] = m.newRenderCmd
		}
		m.viewer, cmds[1] = m.viewer.Update(viewer.SettingsMsg(m.settings))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) handleBackMsg() (Model, tea.Cmd) {
	m.menu = m.menu.SetState(menu.MainMenu)
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
