package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"

	"github.com/Zebbeni/bubbletea_sketches/browser/item"
)

const (
	width  = 40
	height = 20
)

var (
	acceptedFileExts = []string{".jpg", ".png", ".go"}
)

func NewList(directory string) list.Model {
	items := buildItems(directory)
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetHeight(1)
	delegate.SetSpacing(0)
	return list.New(items, delegate, width, height)
}

func buildItems(dir string) []list.Item {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Unable to read directory entries", err)
		os.Exit(1)
	}

	// initialize directory items with an item that allows going up a directory
	dirBaseName := fmt.Sprintf("../%s", filepath.Base(dir))
	dirParentName := filepath.Dir(dir)
	dirItems := []list.Item{
		item.New(dirBaseName, dirParentName, true),
	}
	fileItems := make([]list.Item, 0)

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			dirItems = append(dirItems, item.New(entry.Name(), fullPath, true))
			continue
		}
		for _, ext := range acceptedFileExts {
			if filepath.Ext(entry.Name()) == ext {
				fileItems = append(fileItems, item.New(entry.Name(), fullPath, false))
				continue
			}
		}
	}

	return append(dirItems, fileItems...)
}
