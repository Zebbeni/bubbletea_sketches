package focus

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Handler interface {
	Focus()
	HasFocus() bool
	HandleNav(tea.KeyMsg)
}

type Component interface {
	Handler
	tea.Model
}
