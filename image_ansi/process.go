package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

const PROPORTION = 0.46

// var palette color.Palette
var colorPalette color.Palette

type blockFunc func(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64)

var blockFuncs = map[rune]blockFunc{
	'▀': calcTop,
	'▐': calcRight,
	'▞': calcDiagonal,
	'▖': calcBotLeft,
	'▘': calcTopLeft,
	'▝': calcTopRight,
	'▗': calcBotRight,
}

func process(input image.Image, width int) string {
	imgW, imgH := float32(input.Bounds().Dx()), float32(input.Bounds().Dy())
	height := int(float32(width) * (imgH / imgW) * PROPORTION)
	// resize the sample to be twice the width and height we want to render
	// since we'll try to use each block character to mimic 4 pixels
	refImg := resize.Resize(uint(width)*2, uint(height)*2, input, resize.Lanczos3)
	//refImg := resize.Resize(uint(width)*2, uint(height)*2, input, resize.NearestNeighbor)

	colorPalette = make(color.Palette, 0)
	//p, _ := colorful.SoftPalette(512)
	////p, _ := colorful.HappyPalette(1000)
	////p, _ := colorful.WarmPalette(1000)
	//
	//addColorfulColors := func(cols []colorful.Color) {
	//	for _, c := range cols {
	//		colorPalette = append(colorPalette, c)
	//	}
	//}
	//
	//addColorfulColors(p)

	// sweet 16
	//hexColors := []string{
	//	"#1a1c2c", "#5d275d", "#b13e53", "#ef7d57",
	//	"#ffcd75", "#a7f070", "#38b764", "#257179",
	//	"#29366f", "#3b5dc9", "#41a6f6", "#73eff7",
	//	"#f4f4f4", "#94b0c2", "#566c86", "#333c57",
	//}

	// cucumber richard
	hexColors := []string{
		"#efeaa1", "#b6c157", "#749938", "#1d6e1c",
		"#124f25", "#103934", "#abffea", "#80dcdf",
		"#5890b5", "#395c8c", "#232a65", "#1d164d",
		"#f2db6d", "#dda43e", "#c67a26", "#bf481d",
		"#ac1512", "#6f0611", "#b0a7b4", "#84778a",
		"#585063", "#413b4d", "#333042", "#212030",
	}
	for _, h := range hexColors {
		c, _ := colorful.Hex(h)
		colorPalette = append(colorPalette, c)
	}

	//colorPalette = palette.Plan9
	//colorPalette = palette.WebSafe

	//colorPalette = palette.AnsiWinPowershell16()
	//colorPalette = palette.Ansi256()
	//colorPalette = palette.AnsiVga16()

	//addGamutPalette := func(p gamut.Palette) {
	//	for _, c := range p.Colors() {
	//		colorPalette = append(colorPalette, c.Color)
	//	}
	//}
	//
	//gamutPalette := palette.Crayola
	//addGamutPalette(gamutPalette)

	// TRIADIC
	//addColors := func(cols []color.Color) {
	//	for _, c := range cols {
	//		colorPalette = append(colorPalette, c)
	//	}
	//}

	//b1 := color.RGBA{R: 225, G: 100, B: 100, A: 255}
	//b2 := color.RGBA{R: 100, G: 225, B: 100, A: 255}
	//b3 := color.RGBA{R: 100, G: 100, B: 225, A: 255}
	//b4 := color.RGBA{R: 255, G: 250, B: 100, A: 255}
	////cols := gamut.Triadic(b1)
	//b1s := gamut.Monochromatic(b1, 16)
	//b2s := gamut.Monochromatic(b2, 16)
	//b3s := gamut.Monochromatic(b3, 16)
	//b4s := gamut.Monochromatic(b4, 16)
	////b3s := gamut.Monochromatic(b3, 4)
	////for i := 0; i < len(b1s); i++ {
	////	addColors(gamut.Blends(b1s[i], b2s[i], 4))
	////	addColors(gamut.Blends(b1s[i], b3s[i], 4))
	////	addColors(gamut.Blends(b2s[i], b3s[i], 4))
	////}
	//addColors(b1s)
	//addColors(b2s)
	//addColors(b3s)
	//addColors(b4s)

	//addColors(gamut.Monochromatic(b1, 8))
	//addColors(gamut.Monochromatic(b2, 8))
	//addColors(gamut.Monochromatic(cols[1], 8))

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
			r, fg, bg := getBlock(r1, r2, r3, r4)

			//paletteFg := colorPalette.Convert(fg)
			//paletteBg := colorPalette.Convert(bg)
			//pFg, _ := colorful.MakeColor(paletteFg)
			//pBg, _ := colorful.MakeColor(paletteBg)
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
func getBlock(r1, r2, r3, r4 colorful.Color) (r rune, fg, bg colorful.Color) {
	minDist := 100.0
	for bRune, bFunc := range blockFuncs {
		f, b, dist := bFunc(r1, r2, r3, r4)
		if dist < minDist {
			minDist = dist
			r, fg, bg = bRune, f, b
		}
	}
	return
}

func avgCol(colors ...colorful.Color) (colorful.Color, float64) {
	rSum, gSum, bSum := 0.0, 0.0, 0.0
	for _, col := range colors {
		rSum += col.R
		gSum += col.G
		bSum += col.B
	}
	count := float64(len(colors))
	avg := colorful.Color{R: rSum / count, G: gSum / count, B: bSum / count}

	paletteAvg := colorPalette.Convert(avg)
	avg, _ = colorful.MakeColor(paletteAvg)

	// compute sum of squares
	totalDist := 0.0
	for _, col := range colors {
		totalDist += math.Pow(col.DistanceCIEDE2000(avg), 2)
	}
	return avg, totalDist
}

func calcTop(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r1, r2)
	bg, bDist := avgCol(r3, r4)
	return fg, bg, fDist + bDist
}

func calcRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r2, r4)
	bg, bDist := avgCol(r1, r3)
	return fg, bg, fDist + bDist
}

func calcDiagonal(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r2, r3)
	bg, bDist := avgCol(r1, r4)
	return fg, bg, fDist + bDist
}

func calcBotLeft(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r3)
	bg, bDist := avgCol(r1, r2, r4)
	return fg, bg, fDist + bDist
}

func calcTopLeft(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r1)
	bg, bDist := avgCol(r2, r3, r4)
	return fg, bg, fDist + bDist
}

func calcTopRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r2)
	bg, bDist := avgCol(r1, r3, r4)
	return fg, bg, fDist + bDist
}

func calcBotRight(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64) {
	fg, fDist := avgCol(r4)
	bg, bDist := avgCol(r1, r2, r3)
	return fg, bg, fDist + bDist
}
