package gui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/nav"
)

type Component interface {
	tea.Model
	nav.FocusHandler
}
