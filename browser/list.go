package main

import (
	"fmt"
	"github.com/Zebbeni/bubbletea_sketches/browser/item"
	"github.com/charmbracelet/bubbles/list"
	"os"
	"path/filepath"
)

const (
	width  = 20
	height = 20
)

var (
	acceptedFileExts = []string{".jpg", ".png"}
)

func NewList(directory string) list.Model {
	items := buildItems(directory)
	return list.New(items, list.NewDefaultDelegate(), width, height)
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
		item.New(dirBaseName, dirParentName),
	}
	fileItems := make([]list.Item, 0)

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			dirItems = append(dirItems, item.New(entry.Name(), fullPath))
			continue
		}
		for _, ext := range acceptedFileExts {
			if filepath.Ext(entry.Name()) == ext {
				fileItems = append(fileItems, item.New(entry.Name(), fullPath))
				continue
			}
		}
	}

	return append(dirItems, fileItems...)
}
