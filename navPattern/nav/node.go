package nav

import tea "github.com/charmbracelet/bubbletea"

// Node is a simple FocusHandler implementation for an object
// with no focusable children.
type Node struct {
	hasFocus bool
}

func NewNode(hasFocus bool) *Node {
	return &Node{
		hasFocus: hasFocus,
	}
}

func (n *Node) Focus() {
	n.hasFocus = true
}

func (n *Node) HasFocus() bool {
	return n.hasFocus
}

// HandleNav is only called on currently-focused objects, allowing them to
// transfer focus to another internal id. Since Node has no internal
// components, we set it to false.
func (n *Node) HandleNav(msg tea.KeyMsg) {
	n.hasFocus = false
}
