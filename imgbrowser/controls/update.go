package controls

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
	Down
)

var navMap = map[Direction]map[State]State{
	Right: {Open: Settings, Settings: Export},
	Left:  {Export: Settings, Settings: Open},
}

func (m Model) handleOpenUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.FileBrowser, cmd = m.FileBrowser.Update(msg)

	if m.FileBrowser.ShouldClose {
		m.FileBrowser.ShouldClose = false
		return m.handleControlsUpdate(msg)
	}
	return m, cmd
}

func (m Model) handleSettingsUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Settings, cmd = m.Settings.Update(msg)

	// if Settings no longer focused, handle the message at the control level
	if m.Settings.ShouldClose {
		m.Settings.ShouldClose = false
		return m.handleControlsUpdate(msg)
	}

	return m, cmd
}

func (m Model) handleExportUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Export, cmd = m.Export.Update(msg)

	if m.Export.ShouldClose {
		m.Export.ShouldClose = false
		return m.handleControlsUpdate(msg)
	}

	return m, cmd
}

func (m Model) handleControlsUpdate(msg tea.Msg) (Model, tea.Cmd) {
	// This should only happen if
	m.active = None

	// if key message, determine if we need to switch our button
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			if key.Matches(msg, io.KeyMap.Right) {
				if next, hasNext := navMap[Right][m.focus]; hasNext {
					m.focus = next
				}
			} else if key.Matches(msg, io.KeyMap.Left) {
				if next, hasNext := navMap[Left][m.focus]; hasNext {
					m.focus = next
				}
			}
		case key.Matches(msg, io.KeyMap.Enter):
			m.active = m.focus
			switch m.active {
			case Open:
				m.FileBrowser = m.FileBrowser
			case Settings:
				m.Settings = m.Settings
			case Export:
				m.Export = m.Export
			}
			// Okay this is the problem. Right here we want to call
			// 'Focus()' on the currently focused model. But I really
			// don't want to do another switch case on m.active and do
			// the same thing for each panel type. It would be nice if
			// we could store all panels in a map of Focusable pointers
			// (Actually, I'm referring to what we should do when we set
			// m.active. But same point. My mistake)
		}
	}
	return m, nil
}
