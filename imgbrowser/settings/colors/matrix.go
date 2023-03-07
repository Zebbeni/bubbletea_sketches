package colors

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/makeworld-the-better-one/dither/v2"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

type Matrix struct {
	Name   string
	Method dither.ErrorDiffusionMatrix
}

func getMatrixMenuItems() []list.Item {
	return []list.Item{
		getMatrixItem(Matrix{Name: "Simple2D", Method: dither.Simple2D}),
		getMatrixItem(Matrix{Name: "FloydSteinberg", Method: dither.FloydSteinberg}),
		getMatrixItem(Matrix{Name: "JarvisJudiceNinke", Method: dither.JarvisJudiceNinke}),
		getMatrixItem(Matrix{Name: "Atkinson", Method: dither.Atkinson}),
		getMatrixItem(Matrix{Name: "Stucki", Method: dither.Stucki}),
		getMatrixItem(Matrix{Name: "Burkes", Method: dither.Burkes}),
		getMatrixItem(Matrix{Name: "Sierra", Method: dither.Sierra}),
		getMatrixItem(Matrix{Name: "StevenPigeon", Method: dither.StevenPigeon}),
	}
}

func getMatrixItem(matrix Matrix) list.Item {
	return item{name: matrix.Name, onSelect: func(m Model) (Model, tea.Cmd) {
		_, palette := m.Palette.GetCurrent()
		m.Ditherer = *dither.NewDitherer(palette)

		m.Matrix = matrix
		m.Ditherer.Matrix = matrix.Method
		return m, io.StartRenderCmd
	}}
}
