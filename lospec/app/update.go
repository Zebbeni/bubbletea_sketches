package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
)

func (m Model) update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		}
	}
	return m, nil
}

func (m Model) handleNav(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	case key.Matches(msg, io.KeyMap.Up):
		return m.handleUp()
	case key.Matches(msg, io.KeyMap.Down):
		return m.handleDown()
		m = m.setFocus(navDownLookup[m.focus])
	}
	return m, nil
}

func (m Model) handleUp() (Model, tea.Cmd) {
	switch m.focus {
	case ColorCount:
		m.colorCount.FocusInternal = false
	case TagSearch:
		m.tagSearch.IsFocused = false
	case Sorting:
		m.sortingType.IsFocused = false
	case PaletteList:
		m.paletteList.IsFocused = false
	}
	return m, nil
}

func (m Model) handleDown() (Model, tea.Cmd) {

}

func (m Model) setFocus(focus Focus) Model {
	switch m.focus {
	case ColorCount:
		m.colorCount.IsFocused = false
	case TagSearch:
		m.tagSearch.IsFocused = false
	case Sorting:
		m.sortingType.IsFocused = false
	case PaletteList:
		m.paletteList.IsFocused = false
	}

	m.focus = focus

	switch focus {
	case ColorCount:
		m.colorCount.IsFocused = true
	case TagSearch:
		m.tagSearch.IsFocused = true
	case Sorting:
		m.sortingType.IsFocused = true
	case PaletteList:
		m.paletteList.IsFocused = true
	}

	return m
}

func (m Model) setActive(active Focus) Model {
	switch m.active {
	case ColorCount:
		m.colorCount.FocusInternal = false
	case TagSearch:
		m.tagSearch.FocusInternal = false
	case Sorting:
		m.sortingType.FocusInternal = false
	case PaletteList:
		m.paletteList.FocusInternal = false
	}

	m.active = active

	switch active {
	case ColorCount:
		m.colorCount.FocusInternal = true
	case TagSearch:
		m.tagSearch.FocusInternal = true
	case Sorting:
		m.sortingType.FocusInternal = true
	case PaletteList:
		m.paletteList.FocusInternal = true
	}

	return m
}
