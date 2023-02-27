package main

import (
	"fmt"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/app"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	m := app.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
