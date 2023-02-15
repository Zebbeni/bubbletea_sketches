package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
)

type model struct {
	activeTabIdx int
	tabs         []Tab
	tabContent   []string
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "l", "n", "tab":
			m.activeTabIdx = min(m.activeTabIdx+1, len(m.tabs)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTabIdx = max(int(m.activeTabIdx)-1, 0)
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabs)-1, i == m.activeTabIdx
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t.label))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	tabContent := m.tabs[m.activeTabIdx].content

	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width(lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize()).Render(tabContent))
	return docStyle.Render(doc.String())
}

func main() {

	m := &model{
		activeTabIdx: 0,
		tabs:         tabs,
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
