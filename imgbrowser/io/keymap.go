package io

import (
	"github.com/charmbracelet/bubbles/key"
)

type Map struct {
	Enter key.Binding
	Nav   key.Binding
	Right key.Binding
	Left  key.Binding
	Esc   key.Binding
}

var KeyMap Map

func InitKeyMap() {
	KeyMap = Map{
		Enter: key.NewBinding(
			key.WithKeys("return", "enter"),
			key.WithHelp("↲", "select"),
		),
		Nav: key.NewBinding(
			key.WithKeys("up", "down", "right", "left"),
			key.WithHelp("↕", "nav"),
		),
		Right: key.NewBinding(
			key.WithKeys("right"),
		),
		Left: key.NewBinding(
			key.WithKeys("left"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back"),
		),
	}
}

func (k Map) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Esc}
}

func (k Map) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Enter, k.Esc}}
}
