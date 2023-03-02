package browser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	Dir  string
	File string

	lists []list.Model
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
			return m.handleEsc()
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
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

func (m Model) selectCurrentItem(selectDirectories bool) (Model, bool) {
	i, ok := m.currentList().SelectedItem().(item)
	if !ok {
		panic("Unexpected list item type")
	}

	isFileSelected := false
	if i.isDir {
		if selectDirectories {
			m = m.addListForDirectory(i.path)
		}
	} else {
		m.File = i.path
		isFileSelected = true
	}

	return m, isFileSelected
}

func (m Model) addListForDirectory(dir string) Model {
	newList := list.New(getItems(dir), NewDelegate(), 30, 30)
	newList.Title = fmt.Sprintf(".../%s/", filepath.Base(dir))
	newList.SetShowHelp(false)
	newList.SetShowStatusBar(false)

	newList.KeyMap.ForceQuit.Unbind()
	newList.KeyMap.Quit.Unbind()

	m.Dir = dir
	m.lists = append(m.lists, newList)

	return m
}

func (m Model) View() string {
	return m.currentList().View()
}
