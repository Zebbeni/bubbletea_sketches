package paletteList

import (
	"fmt"
	"image/color"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
	"github.com/Zebbeni/bubbletea_sketches/lospec/paletteList/paletteItem"
)

type Model struct {
	items []paletteItem.Model

	component.Model
}

func New() Model {
	return Model{
		buildItems(),
		component.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			m, cmd = m.handleNavigation(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {

	itemStrings := make([]string, len(m.items))
	for i, itm := range m.items {
		itemStrings[i] = itm.View()
	}

	content := lipgloss.JoinVertical(lipgloss.Top, itemStrings...)
	return m.GetStyle().Render(content)
}

func buildItems() []paletteItem.Model {
	items := make([]paletteItem.Model, 5)
	for i := 0; i < 5; i++ {
		items[i] = paletteItem.New(fmt.Sprintf("Palette %d", i), Ansi256())
	}
	return items
}

func Ansi256() color.Palette {
	p := make(color.Palette, 0, 256)
	for i := 0; i < 256; i++ {
		ansi := termenv.ANSI256.Color(fmt.Sprintf("%d", i))
		col := termenv.ConvertToRGB(ansi)
		p = append(p, col)
	}
	return p
}
