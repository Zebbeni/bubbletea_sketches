package process

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/makeworld-the-better-one/dither/v2"
	"github.com/nfnt/resize"
)

const PROPORTION = 0.46

// var color color.Palette
var colorPalette color.Palette

func (m Renderer) process(input image.Image, width int) string {
	imgW, imgH := float32(input.Bounds().Dx()), float32(input.Bounds().Dy())
	height := int(float32(width) * (imgH / imgW) * PROPORTION)

	// resize the sample to be twice the width and height we want to render
	// since we'll try to use each block character to mimic 4 pixels
	// TODO: Revisit this assumption when we can change allowed characters
	resizeFunc := m.Settings.Sampling.Function
	refImg := resize.Resize(uint(width)*2, uint(height)*2, input, resizeFunc)

	_, colorPalette = m.Settings.Colors.Palette.GetCurrent()

	if m.Settings.Colors.IsDithered {
		ditherer := dither.NewDitherer(colorPalette)
		ditherer.Matrix = m.Settings.Colors.Matrix.Method
		if m.Settings.Colors.IsSerpentine {
			ditherer.Serpentine = true
		}
		refImg = ditherer.Dither(refImg)
	}

	content := ""
	rows := make([]string, height)
	row := make([]string, width)
	for y := 0; y < height*2; y += 2 {
		for x := 0; x < width*2; x += 2 {
			// TODO: Handle when MakeColor returns false if alpha == 0
			// r1 r2
			// r3 r4
			r1, _ := colorful.MakeColor(refImg.At(x, y))
			r2, _ := colorful.MakeColor(refImg.At(x+1, y))
			r3, _ := colorful.MakeColor(refImg.At(x, y+1))
			r4, _ := colorful.MakeColor(refImg.At(x+1, y+1))

			// pick the block, fg and bg color with the lowest total difference
			// convert the colors to ansi, render the block and add it at row[x]
			r, fg, bg := m.getBlock(r1, r2, r3, r4)

			pFg, _ := colorful.MakeColor(fg)
			pBg, _ := colorful.MakeColor(bg)

			lipFg := lipgloss.Color(pFg.Hex())
			lipBg := lipgloss.Color(pBg.Hex())
			style := lipgloss.NewStyle().Foreground(lipFg).Background(lipBg)
			row[x/2] = style.Render(string(r))
		}
		rows[y/2] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	content += lipgloss.JoinVertical(lipgloss.Left, rows...)
	return content
}

// find the best block character and foreground and background colors to match
// a set of 4 pixels. return
func (m Renderer) getBlock(r1, r2, r3, r4 colorful.Color) (r rune, fg, bg colorful.Color) {
	minDist := 100.0
	for bRune, bFunc := range m.blockFuncs {
		f, b, dist := bFunc(r1, r2, r3, r4)
		if dist < minDist {
			minDist = dist
			r, fg, bg = bRune, f, b
		}
	}
	return
}

func (m Renderer) avgCol(colors ...colorful.Color) (colorful.Color, float64) {
	rSum, gSum, bSum := 0.0, 0.0, 0.0
	for _, col := range colors {
		rSum += col.R
		gSum += col.G
		bSum += col.B
	}
	count := float64(len(colors))
	avg := colorful.Color{R: rSum / count, G: gSum / count, B: bSum / count}

	if m.Settings.Colors.IsLimited {
		paletteAvg := colorPalette.Convert(avg)
		avg, _ = colorful.MakeColor(paletteAvg)
	}

	// compute sum of squares
	totalDist := 0.0
	for _, col := range colors {
		totalDist += math.Pow(col.DistanceCIEDE2000(avg), 2)
	}
	return avg, totalDist
}

func (m Renderer) calcTop(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r1, r2)
	bg, bDist := m.avgCol(r3, r4)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r2, r4)
	bg, bDist := m.avgCol(r1, r3)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcDiagonal(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r2, r3)
	bg, bDist := m.avgCol(r1, r4)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcBotLeft(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r3)
	bg, bDist := m.avgCol(r1, r2, r4)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcTopLeft(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r1)
	bg, bDist := m.avgCol(r2, r3, r4)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcTopRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r2)
	bg, bDist := m.avgCol(r1, r3, r4)
	return fg, bg, fDist + bDist
}

func (m Renderer) calcBotRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := m.avgCol(r4)
	bg, bDist := m.avgCol(r1, r2, r3)
	return fg, bg, fDist + bDist
}
