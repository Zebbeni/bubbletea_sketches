package focus

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/shared"
)

// Container is a Handler implementation for an object with focusable children.
type Container struct {
	hasFocus bool

	doFocusInternal bool
	internalFocusID shared.ID
	children        map[shared.ID]Component
	navMap          NavMap
}

func NewContainer(hasFocus, doInternal bool, id shared.ID, navMap NavMap) Container {
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

// HandleNav takes a KeyMsg already determined to be a navigation message and
// recursively redirects this navigation to its internal Handler objects.
// When it we reach a Handler object
func (f *Container) HandleNav(msg tea.KeyMsg) {
	if f.doFocusInternal {
		// update focused shared. this may cause the child to set
		// hasFocus = false, indicating that the focus needs to be
		// moved to a shared higher up the chain.
		f.children[f.internalFocusID].Update(msg)
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
