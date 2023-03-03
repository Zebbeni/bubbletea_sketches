package menu

import "github.com/charmbracelet/bubbles/list"

type item struct {
	name  string
	state State
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

func mainItems() []list.Item {
	return []list.Item{
		item{name: "Interpolation", state: Interpolation},
	}
}
