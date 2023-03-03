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
	selectedItem := m.list.SelectedItem().(item)
	m.ShouldClose = true

	if selectedItem.Function == m.Function {
		return m, nil
	}

	m.Function = selectedItem.Function
	return m, io.StartRenderCmd
}

func (m Model) handleKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
