package browser

type item struct {
	name  string
	path  string
	isDir bool
}

func (i item) FilterValue() string {
	return i.name
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	if i.isDir {
		return "directory"
	}
	return "file"
}
