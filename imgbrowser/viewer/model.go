package viewer

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
)

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
	case io.FinishRenderMsg:
		return m.handleFinishRenderMsg(msg)
	}
	return m, nil
}

func (m Model) handleFinishRenderMsg(msg io.FinishRenderMsg) (Model, tea.Cmd) {
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
