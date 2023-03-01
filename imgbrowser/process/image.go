package process

import (
	"bufio"
	"image"
	"os"
)

// RenderImageFile will eventually be part of a
func RenderImageFile(imgFilePath string) string {
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

	imgString := process(img, 60)
	return imgString
}
