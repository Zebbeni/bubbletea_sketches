package browser

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func (m Model) handleEnter() (Model, tea.Cmd) {
	return m.updateSelected()
}

func (m Model) handleNav(msg tea.KeyMsg) (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)
	m.lists[m.listIndex()], cmds[0] = m.currentList().Update(msg)
	m, cmds[1] = m.updateActive()
	return m, tea.Batch(cmds...)
}

func (m Model) handleEsc() (Model, tea.Cmd) {
	// remove last list if possible (go back to previous)
	if len(m.lists) > 1 {
		m.lists = m.lists[:m.listIndex()]
		return m, nil
	}

	m.ShouldClose = true
	return m, nil
}

func (m Model) updateActive() (Model, tea.Cmd) {
	itm, ok := m.currentList().SelectedItem().(item)
	if !ok {
		panic("Unexpected list item type")
	}

	if itm.isDir == false && m.ActiveFile != itm.path {
		m.ActiveFile = itm.path
		return m, io.StartRenderCmd
	}
	return m, nil
}

func (m Model) updateSelected() (Model, tea.Cmd) {
	itm, ok := m.currentList().SelectedItem().(item)
	if !ok {
		panic("Unexpected list item type")
	}

	if itm.isDir {
		m.SelectedDir = itm.path
		m = m.addListForDirectory(itm.path)
	} else {
		m.SelectedFile = itm.path
		m.ShouldClose = true
	}

	return m, nil
}

func (m Model) addListForDirectory(dir string) Model {
	newList := list.New(getItems(dir), NewDelegate(), 30, 30)
	newList.Title = fmt.Sprintf(".../%s/", filepath.Base(dir))
	newList.SetShowHelp(false)
	newList.SetShowStatusBar(false)

	newList.KeyMap.ForceQuit.Unbind()
	newList.KeyMap.Quit.Unbind()

	m.SelectedDir = dir
	m.lists = append(m.lists, newList)

	return m
}
