package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/lospec/app"
	"github.com/Zebbeni/bubbletea_sketches/lospec/io"
)

func init() {
	io.InitKeyMap()
}

func main() {
	m := app.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("run error:", err)
		os.Exit(1)
	}
}
