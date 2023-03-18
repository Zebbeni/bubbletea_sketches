package controls

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/browser"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/export"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
)

type State int

const (
	None State = iota
	Open
	Settings
	Export
)

var (
	stateOrder = []State{Open, Settings, Export}
	stateNames = map[State]string{
		Open:     "Open",
		Settings: "Settings",
		Export:   "Export",
	}
)

type Model struct {
	active State
	focus  State

	FileBrowser browser.Model
	Settings    settings.Model
	Export      export.Model
}

func New() Model {
	return Model{
		active: None,
		focus:  Open,

		FileBrowser: browser.New(),
		Settings:    settings.New(),
		Export:      export.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.active {
	case Open:
		return m.handleOpenUpdate(msg)
	case Settings:
		return m.handleSettingsUpdate(msg)
	case Export:
		return m.handleExportUpdate(msg)
	}
	return m.handleControlsUpdate(msg)
}

// View displays a row of 3 buttons above 1 of 3 control panels:
// Open | Settings | Export
func (m Model) View() string {
	// draw the top three buttons
	buttons := m.drawButtons()
	var controls string

	switch m.active {
	case Open:
		controls = m.FileBrowser.View()
	case Settings:
		controls = m.Settings.View()
	case Export:
		controls = m.Export.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top, buttons, controls)
}
