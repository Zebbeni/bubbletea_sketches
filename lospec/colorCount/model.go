package colorCount

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/lospec/colorCount/colorNumberFilterType"
	"github.com/Zebbeni/bubbletea_sketches/lospec/colorCount/colorNumberForm"
	"github.com/Zebbeni/bubbletea_sketches/lospec/component"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
)

type Focus int

const (
	ColorNumberFilterType Focus = iota
	ColorNumberForm
)

type Model struct {
	focus Focus

	colorNumberFilterType colorNumberFilterType.Model
	colorNumberForm       colorNumberForm.Model

	component.Model
}

func New() Model {
	m := Model{
		ColorNumberFilterType,
		colorNumberFilterType.New(),
		colorNumberForm.New(),
		component.New(),
	}
	m = m.switchFocus(ColorNumberFilterType)
	m.FocusInternal = true
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focus {
	case ColorNumberFilterType:
		if m.colorNumberFilterType.FocusInternal {
			m.colorNumberFilterType, cmd = m.colorNumberFilterType.Update(msg)
			return m, cmd
		}
	case ColorNumberForm:
		if m.colorNumberFilterType.FocusInternal {
			m.colorNumberFilterType, cmd = m.colorNumberFilterType.Update(msg)
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
	title := "Color Count:"
	filter := m.colorNumberFilterType.View()
	form := m.colorNumberForm.View()

	content := lipgloss.JoinVertical(lipgloss.Top, title, filter, form)
	return m.GetStyle().Render(content)
}
