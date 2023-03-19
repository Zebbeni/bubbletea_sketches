package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleSettingsUpdate(msg tea.Msg) (Model, tea.Cmd) {
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
		m.active = Colors
		m.Colors.ShouldClose = false
		m.ShouldClose = true
	}
	return m, cmd
}

func (m Model) handleCharactersUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Characters, cmd = m.Characters.Update(msg)

	if m.Characters.ShouldClose {
		m.active = Colors
		m.Characters.ShouldClose = false
		m.ShouldClose = true
	}
	return m, cmd
}

func (m Model) handleSizeUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Size, cmd = m.Size.Update(msg)
	if m.Size.ShouldClose {
		m.active = Colors
		m.Size.ShouldClose = false
		m.ShouldClose = true
	}
	if m.Size.ShouldUnfocus {
		return m.handleSettingsUpdate(msg)
	}
	return m, cmd
}

func (m Model) handleSamplingUpdate(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Sampling, cmd = m.Sampling.Update(msg)

	if m.Sampling.ShouldClose {
		m.active = Colors
		m.Sampling.ShouldClose = false
		m.ShouldClose = true
	}
	return m, cmd
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Esc):
		m.ShouldClose = true
	case key.Matches(msg, io.KeyMap.Nav):
		m.ShouldUnfocus = true
	}
	return m, cmd
}

func (m Model) handleEsc() (Model, tea.Cmd) {
	m.ShouldClose = true
	return m, nil
}
