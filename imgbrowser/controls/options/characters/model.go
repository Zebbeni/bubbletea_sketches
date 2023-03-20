package characters

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type State int

const (
	Ascii State = iota
	Blocks
)

type Model struct {
	focus  State
	active State

	ShouldClose   bool
	ShouldUnfocus bool

	IsActive bool
}

func New() Model {
	return Model{
		focus:         Ascii,
		active:        Ascii,
		ShouldClose:   false,
		ShouldUnfocus: false,
		IsActive:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.drawButtons()
}

func (m Model) Selected() State {
	return m.active
}
