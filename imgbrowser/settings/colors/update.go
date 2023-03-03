package colors

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleEsc() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}

func (m Model) handleNav(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	selectedItem := m.list.SelectedItem().(item)
	if selectedItem.value == m.IsPaletted {
		return m, nil
	}
	m.IsPaletted = selectedItem.value
	return m, tea.Batch(cmd, io.StartRenderCmd)
}
