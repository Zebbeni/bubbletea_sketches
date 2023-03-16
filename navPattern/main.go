package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/navPattern/app"
	"github.com/Zebbeni/bubbletea_sketches/navPattern/io"
)

func init() {
	io.InitKeyMap()
}

func main() {
	m := app.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Run failed:", err)
		os.Exit(1)
	}
}
