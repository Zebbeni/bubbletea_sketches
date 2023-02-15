package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
	"strings"
	"time"
)

const (
	text            = "Name:        Nien Nunb\nOccupation:  Smuggler\nGender:      Male\nHome Planet: Sullust\nSpecies:     Sullustan\n\nIn his youth, Nien Nunb became acquainted with Sian Tevv, a fellow Sullustan whose parents regarded Nunb as an unseemly individual and disapproved of his relationship with Tevv. The talented and ambitious Nunb formed a lasting friendship with Tevv, but they eventually went their separate ways, with Nunb mastering the skills necessary to be a spacer while Tevv studied to become a diplomat. Nunb achieved success in his field, entering a legitimate career as one of the top cargo runners of the SoroSuub Corporation, for which he transported minerals and other raw materials to outlying systems. The speed of Nunb's ship and the quality of his work earned him substantial pay and appreciation from his employer, but when SoroSuub chose to support the Galactic Empire and subsequently take control of the Sullust system from its people, Nunb quit the career; in the process, he faced blaster fire and hostile company starfighters.\n\nAfter some soul-searching, Nunb decided to use his skills—acquired through years of smuggling—to undermine the SoroSuub Corporation. He began stealing SoroSuub consignments from the company and shipping them to the Rebellion; as he hoped to inspire his people and rouse them to action, he committed these acts publicly and with bravado. Others soon began to join Nunb in his efforts, and he and his band of smugglers and outlaws became folk heroes to the Sullustan people. Sian Tevv, now a famous political agitator, provided quiet support to Nunb. SoroSuub established the Home Guard with the goal of defending Sullust and stopping Nunb, and while the Guard virtually eliminated pirate attacks by outsiders, it was less successful in operations against Nunb's gang and other Sullustans. Nunb decided that joining the Rebel Alliance and directly fighting the Empire was the only way to affect change, and his gang soon became affiliated with the Alliance, although an Imperial blockade destroyed their ships while they were attempting to establish contact with the Rebels. SoroSuub later allied with the Alliance as well, offering the Sullust system as a secret staging area for the Rebel fleet, and Nunb was hailed as a hero of his people upon this reversal in policy."
	timeToCheckSize = time.Millisecond * 250
	timeToReveal    = time.Millisecond * 3
)

var (
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	textStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("45"))
)

type sizeTickMsg int
type revealTickMsg int

type model struct {
	w, h, paraIdx, runeIdx int
	revealed               string
	paras                  [][]rune
	vp                     viewport.Model
}

func main() {
	vp := viewport.New(1, 1)
	paras := processText(text)
	m := &model{w: 1, h: 1, paraIdx: 0, runeIdx: 0, paras: paras, revealed: "", vp: vp}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(sizeCheckTick, textRevealTick)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case sizeTickMsg:
		return m.handleSizeTick()
	case revealTickMsg:
		return m.handleRevealTick()
	}
	return m, nil
}

func (m *model) View() string {
	return m.vp.View()
}

func processText(t string) [][]rune {
	paras := strings.Split(t, "\n")
	runeArrays := make([][]rune, 0, len(paras))
	for _, p := range paras {
		runeArrays = append(runeArrays, []rune(p))
	}
	return runeArrays
}

func (m *model) handleSizeTick() (tea.Model, tea.Cmd) {
	w, h, _ := term.GetSize(int(os.Stdout.Fd()))
	if w == m.w && h == m.h {
		return m, sizeCheckTick
	}

	m.w, m.h = w, h
	m.vp = viewport.New(w, h)
	m.vp.Style = borderStyle

	m.updateRenderedText()

	return m, tea.Batch(sizeCheckTick, func() tea.Msg { return tea.WindowSizeMsg{Width: w, Height: h} })
}

func (m *model) updateRenderedText() {
	tea.ClearScreen()
	rendered := textStyle.Copy().Width(m.w - 2).Height(m.h - 2).Render(m.revealed)
	m.vp.SetContent(rendered)
}

func (m *model) handleRevealTick() (tea.Model, tea.Cmd) {
	if m.paraIdx >= len(m.paras) {
		return m, nil
	}

	if m.runeIdx < len(m.paras[m.paraIdx]) {
		m.revealed += string((m.paras[m.paraIdx])[m.runeIdx])
		m.updateRenderedText()
	}

	m.runeIdx++
	if m.runeIdx >= len(m.paras[m.paraIdx]) {
		m.paraIdx++
		m.runeIdx = 0
		m.revealed += "\n"
	}
	return m, textRevealTick
}

func sizeCheckTick() tea.Msg {
	time.Sleep(timeToCheckSize)
	return sizeTickMsg(1)
}

func textRevealTick() tea.Msg {
	time.Sleep(timeToReveal)
	return revealTickMsg(1)
}
