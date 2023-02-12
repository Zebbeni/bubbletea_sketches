package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
var textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))

type tickMsg int

type model struct {
	w, h int
	vp   viewport.Model
}

func tick() tea.Msg {
	time.Sleep(time.Second / 4)
	return tickMsg(1)
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(tick)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		w, h, _ := term.GetSize(int(os.Stdout.Fd()))
		if w != m.w || h != m.h {
			m.updateSize(w, h)
		}
		resize := func() tea.Msg { return tea.WindowSizeMsg{Width: w, Height: h} }
		return m, tea.Batch(tick, resize)
	}
	return m, nil
}

func (m *model) View() string {
	return m.vp.View()
}

func (m *model) updateSize(w, h int) {
	m.w = w
	m.h = h

	m.vp = viewport.New(m.w, m.h)
	m.vp.Style = borderStyle

	text := fmt.Sprintf("Size: %d, %d", m.w, m.h)
	rendered := textStyle.Copy().Width(m.w).Height(m.h).Render(text)
	m.vp.SetContent(rendered)

	tea.ClearScreen()
}

func main() {
	m := &model{w: 1, h: 1, vp: viewport.New(1, 1)}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
	}
}
