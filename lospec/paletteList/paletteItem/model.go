package paletteItem

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
)

type Model struct {
	name    string
	palette color.Palette

	component.Model
}

func New(name string, palette color.Palette) Model {
	return Model{
		name,
		palette,
		component.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Palette Item"
}
