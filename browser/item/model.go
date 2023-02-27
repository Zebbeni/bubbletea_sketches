package item

type Model struct {
	Name  string
	Path  string
	IsDir bool
}

func New(n, d string, isDir bool) Model {
	return Model{Name: n, Path: d, IsDir: isDir}
}

func (i Model) Title() string {
	return i.Name
}

func (i Model) Description() string {
	return i.Path
}

func (i Model) FilterValue() string {
	return i.Name
}
