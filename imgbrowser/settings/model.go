package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nfnt/resize"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	Interpolation resize.InterpolationFunction

	DidUpdate bool
}

func New() Model {
	return Model{
		Interpolation: resize.Lanczos3,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
			//case key.Matches(msg, io.KeyMap.Enter):
			//	return m.handleEnter()
			//default:
			//	return m.handleKey(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "View Settings"
}
