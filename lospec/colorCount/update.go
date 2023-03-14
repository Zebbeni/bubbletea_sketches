package colorCount

import tea "github.com/charmbracelet/bubbletea"

func (m Model) handleNavigation(msg tea.KeyMsg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) switchFocus(focus Focus) Model {
	switch m.focus {
	case ColorNumberFilterType:
		m.colorNumberFilterType.IsFocused = false
	case ColorNumberForm:
		m.colorNumberForm.IsFocused = false
	}

	m.focus = focus

	switch focus {
	case ColorNumberFilterType:
		m.colorNumberFilterType.IsFocused = true
	case ColorNumberForm:
		m.colorNumberForm.IsFocused = true
	}

	return m
}
