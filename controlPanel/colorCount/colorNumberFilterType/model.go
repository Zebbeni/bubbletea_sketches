package colorNumberFilterType

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/controlPanel/component"
)

type Model struct {
	component.Model
}

func New() Model {
	return Model{component.New()}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "colorNumberFilterType"
}
