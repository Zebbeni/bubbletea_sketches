package colors

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
)

var navMap = map[Direction]map[State]State{
	Right: {TrueColor: Adaptive, Adaptive: Paletted},
	Left:  {Paletted: Adaptive, Adaptive: TrueColor},
	Up:    {TrueColor: None, Adaptive: None, Paletted: None},
}

func (m Model) handleMenuUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		}
	}
	return m, nil
}

func (m Model) handleAdaptiveUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Adaptive, cmd = m.Adaptive.Update(msg)
	if m.Adaptive.ShouldClose {
		m.Adaptive.ShouldClose = false
		m.active = None
		return m, cmd
	}
	return m, cmd
}

func (m Model) handlePaletteUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Palette, cmd = m.Palette.Update(msg)
	if m.Palette.ShouldClose {
		m.Palette.ShouldClose = false
		m.active = None
		return m, cmd
	}
	return m, cmd
}

func (m Model) handleEsc() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	m.active = m.focus
	return m, io.StartRenderCmd
}

func (m Model) handleNav(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Right):
		if next, hasNext := navMap[Right][m.focus]; hasNext {
			m.focus = next
		}
	case key.Matches(msg, io.KeyMap.Left):
		if next, hasNext := navMap[Left][m.focus]; hasNext {
			m.focus = next
		}
	case key.Matches(msg, io.KeyMap.Up):
		if next, hasNext := navMap[Up][m.active]; hasNext {
			m.active = next
		}
	}

	return m, cmd
}
