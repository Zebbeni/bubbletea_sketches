package process

import (
	"bufio"
	"image"
	"image/color"
	"os"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/controls/options/characters"
)

const PROPORTION = 0.46

// var color color.Palette
var colorPalette color.Palette

func RenderImageFile(s options.Model, imgFilePath string) string {
	if imgFilePath == "" {
		return "Choose an image to render"
	}

	var img image.Image
	imgFile, err := os.Open(imgFilePath)
	if err != nil {
		return "Could not open image " + imgFilePath
	}
	defer imgFile.Close()
	imageReader := bufio.NewReader(imgFile)
	img, _, err = image.Decode(imageReader)
	if err != nil {
		return "Could not decode image " + imgFilePath
	}

	renderer := New(s)
	imgString := renderer.process(img)
	return imgString
}

func (m Renderer) process(input image.Image) string {
	mode, _, _ := m.Settings.Characters.Selected()
	if mode == characters.Ascii {
		return m.processAscii(input)
	}
	return m.processUnicode(input)
}
