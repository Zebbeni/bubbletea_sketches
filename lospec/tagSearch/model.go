package tagSearch

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
	"github.com/Zebbeni/bubbletea_sketches/lospec/tagSearch/tagForm"
	"github.com/Zebbeni/bubbletea_sketches/lospec/tagSearch/tagList"
)

type Focus int

const (
	TagForm Focus = iota
	TagList
)

type Model struct {
	focus Focus

	tagForm tagForm.Model
	tagList tagList.Model

	component.Model
}

func New() Model {
	m := Model{
		TagForm,
		tagForm.New(),
		tagList.New(),
		component.New(),
	}
	m.FocusInternal = true
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focus {
	case TagForm:
		if m.tagForm.FocusInternal {
			m.tagForm, cmd = m.tagForm.Update(msg)
			return m, cmd
		}
	case TagList:
		if m.tagList.FocusInternal {
			m.tagList, cmd = m.tagList.Update(msg)
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			m, cmd = m.handleNavigation(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	form := m.tagForm.View()
	list := m.tagList.View()
	content := lipgloss.JoinVertical(lipgloss.Top, form, list)
	return m.GetStyle().Render(content)
}
