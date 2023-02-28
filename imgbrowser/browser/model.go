package browser

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	Dir  string
	File string

	list list.Model
}

func New() Model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting starting directory:", err)
		os.Exit(1)
	}

	browserList := list.New(getItems(dir), NewDelegate(), 30, 30)
	browserList.SetShowHelp(false)
	browserList.SetShowTitle(false)

	browserList.KeyMap.ForceQuit.Unbind()
	browserList.KeyMap.Quit.Unbind()

	m := Model{Dir: dir, list: browserList}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			m = m.selectCurrentItem()
			return m, nil
		default:
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) selectCurrentItem() Model {
	i, ok := m.list.SelectedItem().(item)
	if !ok {
		panic("Unexpected list item type")
	}
	if i.isDir {
		m.Dir = i.path
		m.list.SetItems(getItems(m.Dir))
	} else {
		m.File = i.path
	}
	return m
}

func (m Model) View() string {
	return m.list.View()
}
