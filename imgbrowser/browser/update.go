package browser

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleEsc() (Model, tea.Cmd) {
	// remove last list on escape if possible (navigate back to previous list)
	if len(m.lists) > 1 {
		m.lists = m.lists[:m.listIndex()]
		return m, nil
	}
	// send command to close the browser experience
	return m, io.CloseCmd
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	var isFile bool
	m, isFile = m.selectCurrentItem(true)
	if isFile {
		return m, io.CloseCmd
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.lists[m.listIndex()], cmd = m.currentList().Update(msg)
	m, _ = m.selectCurrentItem(false)
	return m, cmd
}
