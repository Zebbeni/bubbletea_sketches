package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

const (
	width  = 30
	height = 20
)

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

type item struct {
	name        string
	description string
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	return i.description
}

func (i item) FilterValue() string {
	return i.name
}

func main() {
	items := []list.Item{
		item{name: "item A", description: "first"},
		item{name: "item B", description: "second"},
		item{name: "item C", description: "another"},
		item{name: "item D", description: "last"},
	}
	m := model{list: list.New(items, list.NewDefaultDelegate(), width, height)}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
