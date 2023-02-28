package browser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
)

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

func getItems(dir string) []list.Item {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		os.Exit(1)
	}

	parentPath := filepath.Dir(dir)
	parentName := fmt.Sprintf("←%s/", filepath.Base(parentPath))
	parentItem := item{name: parentName, path: parentPath, isDir: true}

	dirItems := []list.Item{parentItem}
	fileItems := make([]list.Item, 0)

	for _, e := range entries {
		path := fmt.Sprintf("%s/%s", dir, e.Name())
		if e.IsDir() {
			name := fmt.Sprintf("%s/", e.Name())
			dirItem := item{name: name, path: path, isDir: true}
			dirItems = append(dirItems, dirItem)
		} else {
			fileItem := item{name: e.Name(), path: path, isDir: false}
			fileItems = append(fileItems, fileItem)
		}
	}

	return append(dirItems, fileItems...)
}
