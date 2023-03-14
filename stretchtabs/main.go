package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/viewport"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"

	"github.com/Zebbeni/bubbletea_sketches/stretchtabs/tabs"
)

type tickMsg int

func tick() tea.Msg {
	time.Sleep(time.Second / 4)
	return tickMsg(1)
}

type app struct {
	w, h int
	vp   viewport.Model
	tm   tabs.Manager
}

func (m *app) Init() tea.Cmd {
	return tea.Batch(tick)
}

func (m *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		w, h, _ := term.GetSize(int(os.Stdout.Fd()))
		if w == m.w && h == m.h {
			return m, tick
		}
		m.updateSize(w, h)
		return m, tea.Batch(tick, func() tea.Msg { return tea.WindowSizeMsg{Width: w, Height: h} })
	}
	return m, nil
}

func (m *app) View() string {
	content := m.tm.View()
	m.vp.SetContent(content)
	return m.vp.View()
}

func (m *app) updateSize(w, h int) {
	m.w, m.h = w, h
	m.vp.Width, m.vp.Height = w, h

	m.tm = tabs.New(m.w, m.h)
	m.tm.Resize(m.w, m.h)

	tea.ClearScreen()
}

func main() {
	m := &app{w: 1, h: 1, vp: viewport.New(1, 1), tm: tabs.New(1, 1)}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
	}
}
