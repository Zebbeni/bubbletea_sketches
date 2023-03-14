package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/lospec/colorCount"
	"github.com/Zebbeni/bubbletea_sketches/lospec/paletteList"
	"github.com/Zebbeni/bubbletea_sketches/lospec/sortingType"
	"github.com/Zebbeni/bubbletea_sketches/lospec/tagSearch"
)

type Focus int

const (
	ColorCount Focus = iota
	TagSearch
	Sorting
	PaletteList
)

type Model struct {
	focus  Focus
	active Focus

	colorCount  colorCount.Model
	tagSearch   tagSearch.Model
	sortingType sortingType.Model
	paletteList paletteList.Model
}

func New() Model {
	m := Model{
		focus:       ColorCount,
		colorCount:  colorCount.New(),
		tagSearch:   tagSearch.New(),
		sortingType: sortingType.New(),
		paletteList: paletteList.New(),
	}
	m = m.setFocus(ColorCount)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		}
	}
	var cmd tea.Cmd
	switch m.focus {
	case ColorCount:
		if m.colorCount.FocusInternal {
			m.colorCount, cmd = m.colorCount.Update(msg)
			return m, cmd
		}
	case TagSearch:
		if m.tagSearch.FocusInternal {
			m.tagSearch, cmd = m.tagSearch.Update(msg)
			return m, cmd
		}
	case Sorting:
		if m.sortingType.FocusInternal {
			m.sortingType, cmd = m.sortingType.Update(msg)
			return m, cmd
		}
	case PaletteList:
		if m.paletteList.FocusInternal {
			m.paletteList, cmd = m.paletteList.Update(msg)
			return m, cmd
		}
	}

	return m.update(msg)
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
	tags := m.tagSearch.View()
	sorting := m.sortingType.View()
	palettes := m.paletteList.View()

	return lipgloss.JoinVertical(lipgloss.Top, count, tags, sorting, palettes)
}
