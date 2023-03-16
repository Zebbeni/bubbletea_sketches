package nav

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Container is a FocusHandler implementation for an object with focusable children.
type Container struct {
	hasFocus bool

	doFocusInternal bool
	internalFocusID ID
	children        map[ID]FocusHandler
	navMap          Map
}

func NewContainer(hasFocus, doInternal bool, id ID, navMap Map) Container {
	return Container{
		hasFocus:        hasFocus,
		doFocusInternal: doInternal,
		internalFocusID: id,
		navMap:          navMap,
	}
}

func (f *Container) Focus() {
	f.hasFocus = true
}

func (f *Container) HasFocus() bool {
	return f.hasFocus
}

// HandleNav takes a KeyMsg already determined to be a nav message and
// recursively redirects this nav to its internal FocusHandler objects.
// When it we reach a FocusHandler object
func (f *Container) HandleNav(msg tea.KeyMsg) {
	if f.doFocusInternal {
		// update focused id. this may cause the child to set
		// hasFocus = false, indicating that the focus needs to be
		// moved to a id higher up the chain.
		f.children[f.internalFocusID].HandleNav(msg)
		if f.children[f.internalFocusID].HasFocus() {
			return
		}

		dir := GetKeyDirection(msg)

		if nextID, ok := f.navMap.get(f.internalFocusID, dir); ok {
			f.internalFocusID = nextID
			f.children[nextID].Focus()
			return
		}
	}

	f.hasFocus = false
}
