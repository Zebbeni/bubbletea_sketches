package process

import (
	"bufio"
	"image"
	"os"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
)

func RenderImageFile(s settings.Model, imgFilePath string) string {
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

	imgString := process(s, img, 50)
	return imgString
}
