package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleMenuUpdate(msg tea.Msg) (Model, tea.Cmd) {
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
		m.state = Menu
		m.Colors.ShouldClose = false
	}
	return m, cmd
}

func (m Model) handleSamplingUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Sampling, cmd = m.Sampling.Update(msg)

	if m.Sampling.ShouldClose {
		m.state = Menu
		m.Sampling.ShouldClose = false
	}
	return m, cmd
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Esc):
		m, cmd = m.handleEsc()
	case key.Matches(msg, io.KeyMap.Enter):
		return m.handleEnter()
	default:
		return m.handleKey(msg)
	}
	return m, cmd
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	currentItem := m.menu.SelectedItem().(item)
	m.state = currentItem.state
	return m, nil
}

func (m Model) handleEsc() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}

func (m Model) handleKey(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)
	return m, cmd
}
