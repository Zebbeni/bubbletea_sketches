package component

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
)

func GetKeyDirection(msg tea.KeyMsg) Direction {
	switch {
	case key.Matches(msg, io.KeyMap.Up):
		return Up
	case key.Matches(msg, io.KeyMap.Down):
		return Down
	case key.Matches(msg, io.KeyMap.Left):
		return Left
	case key.Matches(msg, io.KeyMap.Right):
		return Right
	}
	return None
}
