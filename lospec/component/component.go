package component

import (
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	IsFocused     bool
	FocusInternal bool
}

func New() Model {
	return Model{
		IsFocused: false,
	}
}

func (m Model) GetStyle() lipgloss.Style {
	if m.IsFocused && !m.FocusInternal {
		return Focused
	}
	return Unfocused
}
