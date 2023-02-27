package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"math/rand"
	"os"
)

var characters = []rune("▀▞▖▘▝▗")

var textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("45")).Background(lipgloss.Color("21"))
var highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
var windowStyle = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()

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
	keymap keymap
}

func main() {
	rand.Seed(0)

	k := initKeymap()
	m := &model{k}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
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
	return m, nil
}

func (m *model) View() string {
	v := viewport.New(20, 10)
	v.Style = lipgloss.NewStyle()
	rows := make([]string, 5)
	row := make([]string, 10)
	content := ""
	for y := 0; y < 5; y++ {
		for x := 0; x < 10; x++ {
			char := string(characters[rand.Intn(len(characters))])
			fg := lipgloss.Color(fmt.Sprintf("%d", rand.Intn(16)))
			bg := lipgloss.Color(fmt.Sprintf("%d", rand.Intn(16)))
			style := lipgloss.NewStyle().Foreground(fg).Background(bg)
			row[x] = style.Render(char)
		}
		rows[y] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	content += lipgloss.JoinVertical(lipgloss.Left, rows...)
	v.SetContent(content)
	return v.View()
}
