package process

import (
	"image"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/makeworld-the-better-one/dither/v2"
	"github.com/nfnt/resize"
)

// A list of Ascii characters by ascending brightness
var asciiChars = []rune("`.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@")
var asciiAZChars = []rune("rczsLTvJFiCfItluneoZYxjyaESwqkPhdVpOGbUAKXHmRDBgMNWQ")
var asciiNumChars = []rune("7315269480")
var asciiSpecChars = []rune("`.-':_,^=;><+!*/?)(|{}[]#$%&@")

func (m Renderer) processAscii(input image.Image, width int) string {
	imgW, imgH := float32(input.Bounds().Dx()), float32(input.Bounds().Dy())
	height := int(float32(width) * (imgH / imgW) * PROPORTION)

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

	content := ""
	rows := make([]string, height)
	row := make([]string, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			c, _ := colorful.MakeColor(refImg.At(x, y))

			if m.Settings.Colors.IsLimited() {
				paletteAvg := colorPalette.Convert(c)
				c, _ = colorful.MakeColor(paletteAvg)
			}

			fg, _ := colorful.MakeColor(c)
			lipFg := lipgloss.Color(fg.Hex())
			style := lipgloss.NewStyle().Foreground(lipFg)

			_, _, brightness := c.Hsl()

			chars := asciiChars
			char := chars[int(brightness*float64(len(chars)-1))]

			row[x] = style.Render(string(char))
		}
		rows[y] = lipgloss.JoinHorizontal(lipgloss.Top, row...)
	}
	content += lipgloss.JoinVertical(lipgloss.Left, rows...)
	return content
}
