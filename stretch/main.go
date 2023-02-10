package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
	"log"
	"os"
	"time"
)

func main() {
	// Initialize our program
	p := tea.NewProgram(&model{100, 100}, tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	w, h int
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m *model) Init() tea.Cmd {
	m.updateSize()
	return tick
}

// Update is called when messages are received. In this case we're going to
// listen to 'tick' messages and check to see if the window size has changed.
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		m.updateSize()
		return m, tick
	}
	return m, nil
}

func (m *model) View() string {
	return fmt.Sprintf("Size: %d, %d", m.w, m.h)
}

func (m *model) updateSize() {
	m.w, m.h, _ = term.GetSize(int(os.Stdout.Fd()))
}

type tickMsg int

func tick() tea.Msg {
	time.Sleep(time.Second / 4)
	return tickMsg(1)
}
