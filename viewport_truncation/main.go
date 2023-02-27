package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	viewH   = 5
	viewW   = 10
	content = "line 1\nline 2\nline 3\nline 4\nline 5\nline 6\nline 7\nline 8\nline 9\nline 10"
)

var (
	borderStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	noBorderStyle = lipgloss.NewStyle()
)

type model struct {
	view1 viewport.Model
	view2 viewport.Model
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.view1, _ = m.view1.Update(msg)
	m.view2, _ = m.view2.Update(msg)
	return m, nil
}

func (m *model) View() string {
	view1Content := m.view1.View()
	view2Content := m.view2.View()
	combined := lipgloss.JoinHorizontal(lipgloss.Left, view1Content, view2Content)
	return combined
}

func main() {
	view1 := viewport.New(viewW, viewH)
	view1.Style = noBorderStyle
	view1.SetContent(content)

	view2 := viewport.New(viewW, viewH)
	view2.Style = borderStyle
	view2.SetContent(content)

	m := &model{
		view1: view1,
		view2: view2,
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Run error:", err)
		os.Exit(1)
	}
}
