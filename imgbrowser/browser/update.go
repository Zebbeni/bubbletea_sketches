package browser

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleEsc() (Model, tea.Cmd) {
	// remove last list on escape if possible (navigate back to previous list)
	if len(m.lists) > 1 {
		m.lists = m.lists[:m.listIndex()]
		return m, nil
	}
	// send command to close the browser experience
	return m, io.CloseCmd
}

func (m Model) handleEnter() (Model, tea.Cmd) {
	var isFile bool
	m, isFile = m.selectCurrentItem(true)
	if isFile {
		return m, io.CloseCmd
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.lists[m.listIndex()], cmd = m.currentList().Update(msg)
	m, _ = m.selectCurrentItem(false)
	return m, cmd
}

// TODO: Stop overloading the term 'Select' here.
// The way we're using this function, selectCurrentItem either means that the
// user selected the item with [Enter] or merely arrowed over it. It's a fuzzy
// area since we're using the current highlighted image to preview the settings
// even if it hasn't been explicitly 'selected'
func (m Model) selectCurrentItem(selectDirectories bool) (Model, bool) {
	itm, ok := m.currentList().SelectedItem().(item)
	if !ok {
		panic("Unexpected list item type")
	}
	isFile := itm.isDir == false

	if isFile {
		if m.File != itm.path {
			m.DidUpdate = true
		}
		m.File = itm.path
		return m, true
	}

	if itm.isDir && selectDirectories {
		m = m.addListForDirectory(itm.path)
	}

	return m, false
}

func (m Model) addListForDirectory(dir string) Model {
	newList := list.New(getItems(dir), NewDelegate(), 30, 30)
	newList.Title = fmt.Sprintf(".../%s/", filepath.Base(dir))
	newList.SetShowHelp(false)
	newList.SetShowStatusBar(false)

	newList.KeyMap.ForceQuit.Unbind()
	newList.KeyMap.Quit.Unbind()

	m.Dir = dir
	m.lists = append(m.lists, newList)

	return m
}
