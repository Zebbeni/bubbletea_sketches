package basic

import (
	"image/color"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/menu"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Model struct {
	menu list.Model

	name    string
	palette color.Palette

	ShouldClose bool
}

func New() Model {
	items := menuItems()
	newMenu := menu.New(items)

	delegate := menu.NewDelegate()
	delegate.ShowDescription = true
	newMenu.SetDelegate(delegate)

	return Model{
		menu:        newMenu,
		name:        items[0].(item).name,
		palette:     items[0].(item).palette,
		ShouldClose: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Enter):
			return m.handleEnter()
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
		case key.Matches(msg, io.KeyMap.Nav):
			return m.handleNav(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.menu.View()
}

func (m Model) GetCurrent() (string, color.Palette) {
	return m.name, m.palette
}
