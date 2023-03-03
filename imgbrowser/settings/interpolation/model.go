package interpolation

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nfnt/resize"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	Function resize.InterpolationFunction

	list list.Model

	ShouldClose bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
		default:
			return m.handleKey(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.list.View()
}

func New() Model {
	items := getItems()
	selected := items[0].(item)
	l := list.New(items, NewDelegate(), 30, 30)

	return Model{Function: selected.Function, list: l}
}
