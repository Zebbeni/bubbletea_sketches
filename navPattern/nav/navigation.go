package nav

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
)

type Direction int

const (
	Right Direction = iota
	Left
	Down
	Up
	None
)

type ID int

type Map map[Direction]map[ID]ID

func NewMap(right, left, down, up map[ID]ID) Map {
	return Map{
		Right: right,
		Left:  left,
		Down:  down,
		Up:    up,
	}
}

func (n Map) get(id ID, dir Direction) (ID, bool) {
	next, hasNext := id, false
	if dirMap, ok := n[dir]; ok {
		// locate and focus the get child id given the direction
		// of the nav
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
