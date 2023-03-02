package menu

import "github.com/charmbracelet/bubbles/list"

type item struct {
	title string
	state State
}

func (i item) FilterValue() string {
	return i.title
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return ""
}

func mainItems() []list.Item {
	return []list.Item{
		item{title: "Open", state: FileMenu},
		item{title: "Settings", state: SettingsMenu},
	}
}
