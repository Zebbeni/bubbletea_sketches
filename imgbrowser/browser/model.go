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

	lists []list.Model

	DidUpdate   bool
	ShouldClose bool
}

func New() Model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting starting directory:", err)
		os.Exit(1)
	}

	m := Model{}.addListForDirectory(dir)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc(), nil
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter(), nil
		default:
			return m.handleKey(msg)
		}
	}
	return m, nil
}

func (m Model) currentList() list.Model {
	return m.lists[m.listIndex()]
}

func (m Model) listIndex() int {
	return len(m.lists) - 1
}

func (m Model) View() string {
	return m.currentList().View()
}
