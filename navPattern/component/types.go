package component

import tea "github.com/charmbracelet/bubbletea"

type Resizable interface {
	Resize(int, int) Resizable
	Update(msg tea.Msg) (Resizable, tea.Cmd)
	View() string
}
