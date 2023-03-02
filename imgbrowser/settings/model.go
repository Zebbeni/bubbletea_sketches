package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nfnt/resize"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

var interpolationNames = map[resize.InterpolationFunction]string{
	resize.NearestNeighbor:   "Nearest Neighbor",
	resize.Bilinear:          "Bilinear",
	resize.Bicubic:           "Bicubic",
	resize.MitchellNetravali: "MitchellNetravali",
	resize.Lanczos2:          "Lanczos2",
	resize.Lanczos3:          "Lanczos3",
}

type Model struct {
	interpolation resize.InterpolationFunction
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, io.KeyMap.Esc):
			return m.handleEsc()
			//case key.Matches(msg, io.KeyMap.Enter):
			//	return m.handleEnter()
			//default:
			//	return m.handleKey(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "View Settings"
}
