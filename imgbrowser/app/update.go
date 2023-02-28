package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/menu"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/viewer"
)

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)
	switch m.menu.State {
	case menu.MainMenu:
		m.menu, cmds[0] = m.menu.Update(msg)
	case menu.FileMenu:
		m.browser, cmds[0] = m.browser.Update(msg)
		// The viewer will need to kick off a new render if the file changed
		m.viewer, cmds[1] = m.viewer.Update(viewer.FileMsg(m.browser.File))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) handleRenderMsg(msg viewer.RenderMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewer, cmd = m.viewer.Update(msg)
	return m, cmd
}
