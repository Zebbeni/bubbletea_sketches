package viewer

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
)

type SettingsMsg settings.Model
type RenderMsg struct {
	FilePath  string
	ImgString string
}

type Model struct {
	imgString string
	settings  settings.Model

	WaitingOnRender bool
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case RenderMsg:
		return m.handleRenderMsg(msg)
	}
	return m, nil
}

func (m Model) handleRenderMsg(msg RenderMsg) (Model, tea.Cmd) {
	m.WaitingOnRender = false
	m.imgString = msg.ImgString
	return m, nil
}

func (m Model) View() string {
	if m.WaitingOnRender {
		return ""
	}
	return m.imgString
}
