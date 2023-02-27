package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"image"
)

type keymap struct {
	next, prev, quit key.Binding
}

func initKeymap() keymap {
	return keymap{
		next: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next"),
		),
		prev: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "prev"),
		),
		quit: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
			key.WithHelp("esc", "quit"),
		),
	}
}

type model struct {
	img     image.Image
	keymap  keymap
	content string
}

func newViewer(image image.Image) tea.Model {
	return &model{image, initKeymap(), ""}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		}
	}
	if m.content == "" {
		m.content = process(inputImg, width)
	}
	return m, nil
}

func (m *model) View() string {
	return m.content
}
