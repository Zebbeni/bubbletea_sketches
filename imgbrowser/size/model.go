package size

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	ShouldUnfocus bool
	ShouldClose   bool
}

func New() Model {
	return Model{
		ShouldUnfocus: false,
		ShouldClose:   false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			m.ShouldUnfocus = true
		case key.Matches(msg, io.KeyMap.Esc):
			m.ShouldClose = true
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "Size"
}
