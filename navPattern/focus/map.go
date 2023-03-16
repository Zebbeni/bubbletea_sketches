package focus

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/shared"
)

type Direction int

const (
	Right Direction = iota
	Left
	Down
	Up
	None
)

type NavMap map[Direction]map[shared.ID]shared.ID

func NewNavMap(right, left, down, up map[shared.ID]shared.ID) NavMap {
	return NavMap{
		Right: right,
		Left:  left,
		Down:  down,
		Up:    up,
	}
}

func (n NavMap) get(id shared.ID, dir Direction) (shared.ID, bool) {
	next, hasNext := id, false
	if dirMap, ok := n[dir]; ok {
		// locate and focus the get child shared given the direction
		// of the navigation
		next, hasNext = dirMap[id]
	}
	return next, hasNext
}

func GetKeyDirection(msg tea.KeyMsg) Direction {
	switch {
	case key.Matches(msg, io.KeyMap.Right):
		return Right
	case key.Matches(msg, io.KeyMap.Left):
		return Left
	case key.Matches(msg, io.KeyMap.Down):
		return Down
	case key.Matches(msg, io.KeyMap.Up):
		return Up
	}
	return None
}
