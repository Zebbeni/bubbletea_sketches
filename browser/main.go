package main

import (
	"fmt"
	"github.com/Zebbeni/bubbletea_sketches/browser/io"
	"github.com/Zebbeni/bubbletea_sketches/browser/item"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	directory string

	list list.Model

	keymap help.KeyMap
}

func New() model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current directory:", err)
		os.Exit(1)
	}

	return model{
		directory: dir,
		list:      NewList(dir),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			i, ok := m.list.SelectedItem().(item.Model)
			if !ok {
				return m, tea.Quit
			}
			if i.IsDir {
				m.directory = i.Path
				m.list = NewList(m.directory)
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func main() {
	io.InitKeyMap()
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}
}
