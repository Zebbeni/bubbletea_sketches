package main

import (
	"bufio"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"image"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// DefaultWidth is a thing
const DefaultWidth = 100

var (
	width    = DefaultWidth
	inputImg image.Image
)

func init() {
	rand.Seed(0)
}

func main() {
	// get the input image
	getArgs()
	// create the viewer model
	m := newViewer(inputImg)
	// update the viewer to display it
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func getArgs() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) < 2 {
		fmt.Println("Usage: ansize <image> [width]>")
		return
	}
	imageName := os.Args[1]
	if len(os.Args) >= 3 {
		var err error
		width, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid width " + os.Args[2] + ". Please enter an integer.")
			return
		}
	}
	imageFile, err := os.Open(imageName)
	if err != nil {
		fmt.Println("Could not open image " + imageName)
		return
	}
	defer imageFile.Close()
	imageReader := bufio.NewReader(imageFile)
	inputImg, _, err = image.Decode(imageReader)
	if err != nil {
		fmt.Println("Could not decode image")
		return
	}

	process(inputImg, width)
}
