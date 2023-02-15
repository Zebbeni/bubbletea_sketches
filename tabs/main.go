package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
)

type Tab int

const (
	MapTab Tab = iota
	DiploTab
	HoldTab
	EconTab
	AgentsTab
	PlanetsTab
	FleetsTab
)

var tabNames = map[Tab]string{
	MapTab:     "Map",
	DiploTab:   "Diplomacy",
	HoldTab:    "Holdings",
	EconTab:    "Economy",
	PlanetsTab: "Planets",
	AgentsTab:  "Agents",
	FleetsTab:  "Fleets",
}

type model struct {
	activeTab  Tab
	tabs       []Tab
	tabContent []string
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "l", "n", "tab":
			m.activeTab = Tab(min(int(m.activeTab)+1, len(m.tabs)-1))
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = Tab(max(int(m.activeTab)-1, 0))
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabs)-1, t == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(tabNames[t]))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.tabContent[m.activeTab]))
	return docStyle.Render(doc.String())
}

func main() {
	//v := // Raw Materials
	//	"" +
	//		"	Nutrients(☘)		125		-13" +
	//		"	Fibers(☙)			233	 	+4" +
	//		"	Minerals(⯁)		2401	-124" +
	//		"   Rare Metals(⬙)		120		-14" +
	//		"   Mythril(❖)			24  	-3" +
	//		// Refined Materials
	//		"	Cloth(▨)			120		+24" +
	//		"	Ghostweave(▩)		540		-12" +
	//		"	Plasteel(🞜) 		1234	+92" +
	//		"	MoonMetal(◈)		123		-12" +
	//		"	DreamGlass(◇)		120		-20" +
	//		// Fuels
	//		"	Fossoil()			4120	-13" +
	//		"	Brutine(☀) 			123   	-12" +
	//		"	Antimatter(❇) 		123		+23" +
	//		// Energy
	//		"	Energy(🗲)			123		-24" +
	//		// Goods
	//		"	Hard Goods(🛍)		40		-231" +
	//		"	Amenities(♫)		1239	-203" +
	//		"	Luxury Goods(✧)		230	 	+2" +
	//		" 	Medicine(⛨)		124		+92" +
	//		"	Drugs(♨) 			23	 	-41" +
	//		"	Weapons(⚔)			410	 	-23" +
	//		"	Heavy Weapons(♜)	20		+14" +
	//		// Labor
	//		"	Services(⛏)		1204	120" +
	//		"	Engineering(🛠)		1243	-20" +
	//		"	Slaves(☹) 			23	 	+2"

	//phases := "🌑︎ 🌒︎ 🌓︎ 🌖︎ 🌕︎ 🌖︎ 🌗︎ 🌘︎"

	m := &model{
		activeTab:  MapTab,
		tabs:       []Tab{MapTab, DiploTab, HoldTab, EconTab, AgentsTab, PlanetsTab, FleetsTab},
		tabContent: []string{"Map Tab", "Diplomacy Tab", "Holdings Tab", "Economy Tab", "Agents Tab", "Planets Tab", "Fleets Tab"},
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
