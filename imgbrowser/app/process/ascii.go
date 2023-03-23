package process

import (
	"image"
	"image/color"
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/makeworld-the-better-one/dither/v2"
	"github.com/nfnt/resize"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options/characters"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options/size"
)

// A list of Ascii characters by ascending brightness
var asciiChars = []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@")
var asciiAZChars = []rune(" rczsLTvJFiCfItluneoZYxjyaESwqkPhdVpOGbUAKXHmRDBgMNWQ")
var asciiNumChars = []rune(" 7315269480")
var asciiSpecChars = []rune(" `.-':_,^=;><+!*/?)(|{}[]#$%&@")
var unicodeShadeChars = []rune{' ', '░', '▒', '▓'}

func (m Renderer) processAsciiOld(input image.Image) string {
	imgW, imgH := float32(input.Bounds().Dx()), float32(input.Bounds().Dy())

	dimensionType, width, height := m.Settings.Size.Info()
	if dimensionType == size.Fit {
		fitHeight := float32(width) * (imgH / imgW) * PROPORTION
		fitWidth := (float32(height) * (imgW / imgH)) / PROPORTION
		if fitHeight > float32(height) {
			width = int(fitWidth)
		} else {
			height = int(fitHeight)
		}
	}

	resizeFunc := m.Settings.Sampling.Function
	refImg := resize.Resize(uint(width), uint(height), input, resizeFunc)

	_, colorPalette = m.Settings.Colors.Palette.GetCurrent()

	if m.Settings.Colors.IsDithered() {
		ditherer := dither.NewDitherer(colorPalette)
		ditherer.Matrix = m.Settings.Colors.Matrix()
		if m.Settings.Colors.IsSerpentine() {
			ditherer.Serpentine = true
		}
		refImg = ditherer.Dither(refImg)
	}

	var chars []rune
	_, charMode, _ := m.Settings.Characters.Selected()
	switch charMode {
	case characters.AzAscii:
		chars = asciiAZChars
	case characters.NumAscii:
		chars = asciiNumChars
	case characters.SpecAscii:
		chars = asciiSpecChars
	case characters.AllAscii:
		chars = asciiChars
	}

	content := ""
	rows := make([]string, height)
	row := make([]string, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			c, _ := colorful.MakeColor(refImg.At(x, y))
			_, _, brightness := c.Hsl()

			if m.Settings.Colors.IsLimited() {
				paletteAvg := colorPalette.Convert(c)
				c, _ = colorful.MakeColor(paletteAvg)
			}

			fg, _ := colorful.MakeColor(c)
			lipFg := lipgloss.Color(fg.Hex())
			style := lipgloss.NewStyle().Foreground(lipFg)

			char := chars[int(brightness*float64(len(chars)-1))]

			row[x] = style.Render(string(char))
		}
		rows[y] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	content += lipgloss.JoinVertical(lipgloss.Left, rows...)
	return content
}

func (m Renderer) processAscii(input image.Image) string {
	imgW, imgH := float32(input.Bounds().Dx()), float32(input.Bounds().Dy())

	dimensionType, width, height := m.Settings.Size.Info()
	if dimensionType == size.Fit {
		fitHeight := float32(width) * (imgH / imgW) * PROPORTION
		fitWidth := (float32(height) * (imgW / imgH)) / PROPORTION
		if fitHeight > float32(height) {
			width = int(fitWidth)
		} else {
			height = int(fitHeight)
		}
	}

	resizeFunc := m.Settings.Sampling.Function
	refImg := resize.Resize(uint(width)*2, uint(height)*2, input, resizeFunc)

	_, colorPalette = m.Settings.Colors.Palette.GetCurrent()

	if m.Settings.Colors.IsDithered() {
		ditherer := dither.NewDitherer(colorPalette)
		ditherer.Matrix = m.Settings.Colors.Matrix()
		if m.Settings.Colors.IsSerpentine() {
			ditherer.Serpentine = true
		}
		refImg = ditherer.Dither(refImg)
	}

	var chars []rune
	_, charMode, _ := m.Settings.Characters.Selected()
	switch charMode {
	case characters.AzAscii:
		chars = asciiAZChars
	case characters.NumAscii:
		chars = asciiNumChars
	case characters.SpecAscii:
		chars = asciiSpecChars
	case characters.AllAscii:
		chars = asciiChars
	}

	content := ""
	rows := make([]string, height)
	row := make([]string, width)

	for y := 0; y < height*2; y += 2 {
		for x := 0; x < width*2; x += 2 {
			r1, _ := colorful.MakeColor(refImg.At(x, y))
			r2, _ := colorful.MakeColor(refImg.At(x+1, y))
			r3, _ := colorful.MakeColor(refImg.At(x, y+1))
			r4, _ := colorful.MakeColor(refImg.At(x+1, y+1))

			fg, bg, brightness := m.fgBgBrightness(r1, r2, r3, r4)

			lipFg := lipgloss.Color(fg.Hex())
			lipBg := lipgloss.Color(bg.Hex())
			style := lipgloss.NewStyle().Foreground(lipFg).Background(lipBg)

			index := int(brightness * float64(len(chars)-1))
			char := chars[index]
			charString := string(char)

			row[x/2] = style.Render(charString)
		}
		rows[y/2] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	content += lipgloss.JoinVertical(lipgloss.Left, rows...)
	return content
}

func (m Renderer) fgBgBrightness(c ...colorful.Color) (fg, bg colorful.Color, b float64) {
	// find the darkest and lightest among given colors
	lightestCol, _ := lightDark(c...)

	avg := m.avgColTrue(c...)
	avgCol, _ := colorful.MakeColor(avg)

	dist := avgCol.DistanceCIEDE2000(lightestCol)
	brightness := math.Max(0.0, math.Min(1.0, math.Abs(dist*2.0)))

	// if paletted:
	//   convert the darkest to its closest paletted color
	//   convert the lightest to its closest paletted color (excluding the previously found color)
	if m.Settings.Colors.IsLimited() {
		_, colorPalette = m.Settings.Colors.Palette.GetCurrent()

		//index := colorPalette.Index(darkestCol)
		//paletteDark := colorPalette.Convert(darkestCol)
		index := colorPalette.Index(avgCol)
		paletteAvg := colorPalette.Convert(avgCol)

		palette := make([]color.Color, len(colorPalette))
		copy(palette, colorPalette)

		//paletteMinusDarkest := color.Palette(append(palette[:index], palette[index+1:]...))
		//paletteLight := paletteMinusDarkest.Convert(lightestCol)

		paletteMinusAvg := color.Palette(append(palette[:index], palette[index+1:]...))
		paletteLight := paletteMinusAvg.Convert(avgCol)

		lightestCol, _ = colorful.MakeColor(paletteLight)
		//darkestCol, _ = colorful.MakeColor(paletteDark)
		avgCol, _ = colorful.MakeColor(paletteAvg)

		// Account for the fact that both darkest and lightest might be lighter than the lighest
		// color in the limited color palette. In this case the lightest col would be darker than
		// the 'darkest' color (which wouldn't make sense). Instead, we'll just indicate to
		// use the lowest color brightess and only display the darkest color
		//_, _, darkLightness := darkestCol.Hsl()
		//_, _, lightLightness := lightestCol.Hsl()
		//if lightLightness <= darkLightness {
		//	temp := lightestCol
		//	lightestCol = darkestCol
		//	darkestCol = temp
		//}
		_, _, avgLightness := avgCol.Hsl()
		_, _, lightLightness := lightestCol.Hsl()
		if lightLightness < avgLightness {
			temp := lightestCol
			lightestCol = avgCol
			avgCol = temp
		}
	}

	return lightestCol, avgCol, brightness
}

func (m Renderer) avgColTrue(colors ...colorful.Color) colorful.Color {
	rSum, gSum, bSum := 0.0, 0.0, 0.0
	for _, col := range colors {
		rSum += col.R
		gSum += col.G
		bSum += col.B
	}
	count := float64(len(colors))
	avg := colorful.Color{R: rSum / count, G: gSum / count, B: bSum / count}

	return avg
}

func lightDark(c ...colorful.Color) (light, dark colorful.Color) {
	mostLight, mostDark := 0.0, 1.0
	for _, col := range c {
		_, _, l := col.Hsl()
		if l < mostDark {
			mostDark = l
			dark = col
		}
		if l > mostLight {
			mostLight = l
			light = col
		}
	}
	return
}
