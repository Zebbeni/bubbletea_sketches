package lospec

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/controlPanel/colorCount"
)

type Focus int

const (
	ColorCount Focus = iota
	TagSearch
	Sorting
	PaletteList
)

type Model struct {
	focus Focus

	colorCount colorCount.Model
}

func New() Model {
	return Model{
		focus: ColorCount,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focus {
	case ColorCount:
		if m.colorCount.IsActive() {
			m.colorCount, cmd = m.colorCount.Update(msg)
			return m, cmd
		}
		//case TagSearch:
		//	if m.tagSearch.IsActive() {
		//		m.tagSearch, cmd = m.tagSearch.Update(msg)
		//		return m, cmd
		//	}
		//case Sorting:
		//	m.sorting = m.sorting.Update(msg)
		//case PaletteList:
		//	m.paletteList = m.paletteList.Update(msg)
	}
	return m, cmd
}

// View displays the following components
//  1. Color Count
//     1a. Title [string-unselectable]
//     1b. Number [Radio Buttons]
//     1c. Colors [Form]
//  2. Tag: [Form]
//     Tag1
//     Tag2
//  3. Sorting
//     : Default | A-Z | Downloads | Newest |
//  4. Palette list (page)
//     Palette 1
//     Palette 2
//     Palette 3
//     Palette 4
//     ▞▞▞▞▞▞▞▞▞▞
//     Palette 5
//     Palette 6
//  5. Page 5 / 23
func (m Model) View() string {
	count := m.colorCount.View()
	return lipgloss.JoinVertical(lipgloss.Top, count)
}
