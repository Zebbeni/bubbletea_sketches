package interpolation

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
	if selectedItem.Function == m.Function {
		return m, nil
	}
	m.Function = selectedItem.Function
	return m, tea.Batch(cmd, io.StartRenderCmd)
}
