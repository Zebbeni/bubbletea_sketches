package item

type Model struct {
	name string
	path string
}

func New(n, d string) Model {
	return Model{name: n, path: d}
}

func (i Model) Title() string {
	return i.name
}

func (i Model) Description() string {
	return i.path
}

func (i Model) FilterValue() string {
	return i.name
}
