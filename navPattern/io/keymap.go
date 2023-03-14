package io

import (
	"github.com/charmbracelet/bubbles/key"
)

type Map struct {
	Enter key.Binding
	Nav   key.Binding
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
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
			key.WithKeys("up", "down", "left", "right"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right"),
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
