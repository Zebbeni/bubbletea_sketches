package menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type State int

const (
	MainMenu State = iota
	FileMenu
	SettingsMenu
)

type Model struct {
	State State
	list  list.Model
}

func New() Model {
	menuList := list.New(mainItems(), NewDelegate(), 20, 20)
	menuList.SetShowHelp(false)
	menuList.SetShowFilter(false)
	menuList.SetShowTitle(false)

	menuList.KeyMap.ForceQuit.Unbind()
	menuList.KeyMap.Quit.Unbind()

	return Model{
		State: MainMenu,
		list:  menuList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	}
	return m, nil
}

func (m Model) View() string {
	return m.list.View()
}

func (m Model) SetState(state State) Model {
	m.State = state
	return m
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch {
	case key.Matches(msg, io.KeyMap.Enter):
		item, ok := m.list.SelectedItem().(item)
		if !ok {
			panic("Unexpected list item type")
		}
		m.State = item.state
		return m, nil
	default:
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}
