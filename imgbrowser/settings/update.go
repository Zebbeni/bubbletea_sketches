package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleMainUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	}
	return m, nil
}

func (m Model) handleColorsUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Colors, cmd = m.Colors.Update(msg)

	if m.Colors.ShouldClose {
		m.state = Main
		m.Colors.ShouldClose = false
	}
	return m, cmd
}

func (m Model) handleInterpolationUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Sampling, cmd = m.Sampling.Update(msg)

	if m.Sampling.ShouldClose {
		m.state = Main
		m.Sampling.ShouldClose = false
	}
	return m, cmd
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Esc):
		m, cmd = m.handleEsc(msg)
	case key.Matches(msg, io.KeyMap.Enter):
		return m.handleEnter(msg)
	default:
		return m.handleKey(msg)
	}
	return m, cmd
}

func (m Model) handleEnter(msg tea.Msg) (Model, tea.Cmd) {
	currentItem := m.menu.SelectedItem().(item)
	m.state = currentItem.state
	return m, nil
}

func (m Model) handleEsc(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.state {
	case Main:
		m.ShouldClose = true
	case Interpolation:
		m.Sampling, cmd = m.Sampling.Update(msg)
	}
	return m, cmd
}

func (m Model) handleKey(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return m, cmd
}

func (m Model) handleCloseFlags() Model {
	if m.Sampling.ShouldClose {
		m.Sampling.ShouldClose = false
		m.state = Main
	}
	return m
}
