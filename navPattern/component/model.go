package component

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
)

type ID int

type Focusable interface {
	Focus()
	IsFocused() bool
	FocusInternal(Direction)
}

type Model interface {
	tea.Model
}

type Component struct {
	hasFocus            bool
	shouldFocusInternal bool

	internalFocusID ID
	children        map[ID]*Component

	focusMap map[ID]map[Direction]ID
}

func (c *Component) Init() tea.Cmd {
	return nil
}

func (c *Component) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Nav):
			if c.shouldFocusInternal {
				// update focused component. this may cause the child to set
				// hasFocus = false, indicating that the focus needs to be
				// moved to a component higher up the chain.
				c.children[c.internalFocusID].Update(msg)
				if c.children[c.internalFocusID].HasFocus() {
					return c, nil
				}

				if idMap, ok := c.focusMap[c.internalFocusID]; ok {
					// locate and focus the next child component given the direction
					// of the navigation
					dir := GetKeyDirection(msg)
					if nextID, hasNext := idMap[dir]; hasNext {
						c.internalFocusID = nextID
						c.children[c.internalFocusID].Focus()
						return c, nil
					}
				}
			}

			c.hasFocus = false
		}
	}
	return c, nil
}

func (c *Component) View() string {
	return "view"
}

func (c *Component) Focus() {
	c.hasFocus = true
}

func (c *Component) HasFocus() bool {
	return c.hasFocus
}
