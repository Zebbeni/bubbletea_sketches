package export

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	isFocused bool
}

func New() Model {
	return Model{
		isFocused: false,
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

		}
	}
	return m, nil
}

func (m Model) View() string {
	return "Export Menu"
}

func (m Model) Focus() Model {
	m.isFocused = true
	return m
}

func (m Model) IsFocused() bool {
	return m.isFocused
}
