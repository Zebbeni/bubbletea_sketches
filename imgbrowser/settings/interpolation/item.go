package interpolation

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/nfnt/resize"
)

type item struct {
	name     string
	Function resize.InterpolationFunction
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

func getItems() []list.Item {
	items := make([]list.Item, len(nameMap))

	for i, f := range Functions {
		items[i] = item{name: nameMap[f], Function: f}
	}

	return items
}

func NewDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()
	delegate.SetSpacing(0)
	delegate.ShowDescription = false
	return delegate
}
