package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/controlPanel/io"
	"github.com/Zebbeni/bubbletea_sketches/controlPanel/lospec"
)

func init() {
	io.InitKeyMap()
}

func main() {
	m := lospec.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("run error:", err)
		os.Exit(1)
	}
}
