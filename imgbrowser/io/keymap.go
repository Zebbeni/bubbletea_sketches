package io

import (
	"github.com/charmbracelet/bubbles/key"
)

type Map struct {
	Enter key.Binding
	//Quit  key.Binding
	//Nav   key.Binding
	//Back  key.Binding
}

var KeyMap Map

func InitKeyMap() {
	KeyMap = Map{
		Enter: key.NewBinding(
			key.WithKeys("return", "enter"),
			key.WithHelp("↲", "select"),
		),
		//Quit: key.NewBinding(
		//	key.WithKeys("ctrl+c"),
		//	key.WithHelp("ctrl+c", "quit"),
		//),
		//Nav: key.NewBinding(
		//	key.WithKeys("up", "down"),
		//	key.WithHelp("↑↓", "navigate"),
		//),
		//Back: key.NewBinding(
		//	key.WithKeys("esc"),
		//	key.WithHelp("esc", "back"),
		//),
	}
}

func (k Map) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter}
	//return []key.Binding{k.Quit, k.Nav, k.Enter, k.Back}
}

func (k Map) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Enter}}
	//return [][]key.Binding{{k.Quit, k.Nav, k.Enter, k.Back}}
}
