package browser

import (
	"fmt"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"path/filepath"
)

type Model struct {
	dirPath  string
	filePath string

	list list.Model
}

func New() Model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting starting directory:", err)
		os.Exit(1)
	}

	m := Model{dirPath: dir, list: buildList(dir)}
	return m
}

func buildList(dir string) list.Model {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		os.Exit(1)
	}

	parentPath := filepath.Dir(dir)
	parentName := fmt.Sprintf("%s/", filepath.Base(parentPath))
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
	return list.New(items, list.NewDefaultDelegate(), 30, 30)
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
		m.dirPath = i.path
		m.list = buildList(m.dirPath)
	} else {
		m.filePath = i.path
	}
	return m
}

func (m Model) View() string {
	return m.list.View()
}
