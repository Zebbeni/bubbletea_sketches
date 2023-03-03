package colors

import (
	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	name  string
	value bool
}

func (i item) FilterValue() string {
	return i.name
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	return ""
}

func menuItems() []list.Item {
	return []list.Item{
		item{name: "TrueColor", value: false},
		item{name: "Paletted", value: true},
	}
}

func newMenu(items []list.Item) list.Model {
	newList := list.New(items, NewDelegate(), 30, 30)
	newList.SetShowHelp(false)
	newList.SetShowFilter(false)
	newList.SetShowTitle(false)
	newList.SetShowStatusBar(false)

	newList.KeyMap.ForceQuit.Unbind()
	newList.KeyMap.Quit.Unbind()
	return newList
}

func NewDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.SetSpacing(0)
	delegate.ShowDescription = false
	return delegate
}
