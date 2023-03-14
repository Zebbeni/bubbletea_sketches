package app

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/component"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
)

//func (m *Model) handleNav(msg tea.KeyMsg) (Model, bool) {
//	// figure out if the currently-focused child can handle the navigation internally
//	// if not, handle navigation with the app's FocusHandler (?)
//	// ... or is there a way to make FocusHandler work recursively?
//	isHandledInternally := false
//	switch {
//	case key.Matches(msg, io.KeyMap.Up):
//		m.UpdateFocus(component.Up)
//	case key.Matches(msg, io.KeyMap.Down):
//		m.UpdateFocus(component.Down)
//	case key.Matches(msg, io.KeyMap.Left):
//		m.UpdateFocus(component.Left)
//	case key.Matches(msg, io.KeyMap.Right):
//		m.UpdateFocus(component.Right)
//	}
//
//	m.UpdateFocus(d)
//}

func (m *Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	case key.Matches(msg, io.KeyMap.Nav):
		dir := component.GetKeyDirection(msg)
		isHandledInternally := m.UpdateFocus(dir)
		if isHandledInternally {

		}
	}
	return m, nil
}