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

	list list.Model
}

func New() Model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting starting directory:", err)
		os.Exit(1)
	}

	m := Model{Dir: dir, list: buildList(dir)}
	return m
}

func buildList(dir string) list.Model {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		os.Exit(1)
	}

	parentPath := filepath.Dir(dir)
	parentName := fmt.Sprintf("←%s/", filepath.Base(parentPath))
	parentItem := item{name: parentName, path: parentPath, isDir: true}

	dirItems := []list.Item{parentItem}
	fileItems := make([]list.Item, 0)

	for _, e := range entries {
		path := fmt.Sprintf("%s/%s", dir, e.Name())
		if e.IsDir() {
			name := fmt.Sprintf("%s/", e.Name())
			dirItem := item{name: name, path: path, isDir: true}
			dirItems = append(dirItems, dirItem)
		} else {
			fileItem := item{name: e.Name(), path: path, isDir: false}
			fileItems = append(fileItems, fileItem)
		}
	}

	items := append(dirItems, fileItems...)
	l := list.New(items, NewDelegate(), 30, 30)
	l.SetShowHelp(false)

	return l
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
		m.list = buildList(m.Dir)
	} else {
		m.File = i.path
	}
	return m
}

func (m Model) View() string {
	return m.list.View()
}
