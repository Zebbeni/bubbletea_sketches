package nav

import tea "github.com/charmbracelet/bubbletea"

type FocusHandler interface {
	Focus()
	HasFocus() bool
	HandleNav(tea.KeyMsg)
}
