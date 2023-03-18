package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
)

func init() {
	io.InitKeyMap()
}

func main() {
	m := app.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
