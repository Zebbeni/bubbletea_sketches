package palette

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette/basic"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette/file"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette/lospec"
)

type State int

const (
	Basic State = iota
	FromFile
	Lospec
)

type Info struct {
	name   string
	colors color.Palette
}

type Model struct {
	state State

	basic  basic.Model
	file   file.Model
	lospec lospec.Model

	ShouldClose bool
}

func New() Model {
	return Model{
		state:       Basic,
		basic:       basic.New(),
		file:        file.New(),
		lospec:      lospec.New(),
		ShouldClose: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch m.state {
	case Basic:
		return m.handleBasicUpdate(msg)
	}
	return m, nil
}

func (m Model) View() string {
	switch m.state {
	case Basic:
		return m.basic.View()
	}
	return ""
}

func (m Model) SetState(s State) Model {
	if m.state == s {
		return m
	}
	m.state = s
	// This should also update the current palette
	return m
}

func (m Model) GetCurrent() (string, color.Palette) {
	switch m.state {
	case Basic:
		return m.basic.GetCurrent()
	case FromFile:
		return m.file.GetCurrent()
	case Lospec:
		return m.lospec.GetCurrent()
	}
	return "", nil
}
