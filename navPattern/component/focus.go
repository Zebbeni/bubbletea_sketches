package component

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

type Focus int

type NextMap map[Direction]map[Focus]Focus

type FocusHandler struct {
	IsFocused bool

	shouldFocusInternal bool
	// internalFocus should always be non-nil if shouldFocusInternal == true
	internalFocus Focus

	nextMap NextMap
}

func New(focus Focus, isFoc bool, navMap NextMap) FocusHandler {
	return FocusHandler{
		internalFocus: focus,
		IsFocused:     isFoc,
		nextMap:       navMap,
	}
}

// UpdateFocus takes a direction and updates the internalFocus according to the nextMap
// (if possible) before returning if the navigation could be done internally.
// A 'false' response indicates to the caller that navigation needs to occur
// higher up the chain.
func (n *FocusHandler) UpdateFocus(d Direction) bool {
	if !n.shouldFocusInternal {
		return false
	}
	if newFocus, ok := n.nextMap.Get(n.internalFocus, d); ok {
		n.internalFocus = newFocus
		return true
	}
	return false
}

func (n *FocusHandler) SetIsFocused(isFocused bool) {
	n.IsFocused = isFocused
}

func (n *FocusHandler) SetFocusInternal(focusInternal bool) {
	n.shouldFocusInternal = focusInternal
}

func (m NextMap) Get(f Focus, d Direction) (Focus, bool) {
	dirMap, ok := m[d]
	if !ok {
		return f, false
	}
	newFocus, hasNew := dirMap[f]
	return newFocus, hasNew
}
