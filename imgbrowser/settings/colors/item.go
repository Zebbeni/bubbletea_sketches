package colors

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/io"
	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings/colors/palette"
)

type updateFunc func(m Model) (Model, tea.Cmd)

type item struct {
	name     string
	onSelect updateFunc
}

func (i item) FilterValue() string {
	return i.name
}

func (i item) Title() string {
	return i.name
}

func (i item) Description() string {
	return ""
}

//  1. Limited (on/off)
//  2. Palette (Name) (if Limited) -> shows / hides Palette submenu
//     2a) Basic (if Palette submenu shown)
//     2b) From File (if Palette submenu shown)
//     2c) From Lospec (if Palette submenu shown)
//  3. Dithering (on/off) (if Limited)
//  4. Serpentine (on/off) (if Dithering)
//  5. Matrix (Name) (if Dithering) -> shows / hides Matrix submenu
//     5a) Simple2D (if Matrix submenu shown...)
//     5b) FloydSteinberg
//     5c) FalseFloydSteinberg
//     5d) JarvisJudiceNinke
//     5e) Atkinson
//     5f) Stucki
//     5g) Burkes
//     5h) Sierra
//     5i) StevenPigeon
func buildMenuItems(m Model) []list.Item {

	items := make([]list.Item, 0, 20)
	items = append(items, getLimitedToggle(m))
	if !m.IsLimited {
		return items
	}

	items = append(items, getShowPaletteMenuToggle(m))
	if m.showPaletteMenu {
		items = append(items, getPaletteMenuItems()...)
	}

	items = append(items, ditheringToggle(m))
	if m.IsDithered {
		items = append(items, serpentineToggle(m))
		items = append(items, showMatrixMenuToggle(m))
		if m.showMatrixMenu {
			items = append(items, getMatrixMenuItems()...)
		}
	}

	return items
}

func getLimitedToggle(m Model) list.Item {
	name := "Limited (Off)"
	if m.IsLimited {
		name = "Limited (On)"
	}
	return item{name: name, onSelect: func(m Model) (Model, tea.Cmd) {
		m.IsLimited = !m.IsLimited
		return m, io.StartRenderCmd
	}}
}

func ditheringToggle(m Model) list.Item {
	name := "Dithering (Off)"
	if m.IsDithered {
		name = "Dithering (On)"
	}
	return item{name: name, onSelect: func(m Model) (Model, tea.Cmd) {
		m.IsDithered = !m.IsDithered
		return m, io.StartRenderCmd
	}}
}

func serpentineToggle(m Model) list.Item {
	name := "Serpentine (Off)"
	if m.IsSerpentine {
		name = "Serpentine (On)"
	}
	return item{name: name, onSelect: func(m Model) (Model, tea.Cmd) {
		m.IsSerpentine = !m.IsSerpentine
		return m, io.StartRenderCmd
	}}
}

func getShowPaletteMenuToggle(m Model) list.Item {
	name := fmt.Sprintf("Palette (%s)", m.paletteName())
	return item{name: name, onSelect: func(m Model) (Model, tea.Cmd) {
		m.showPaletteMenu = !m.showPaletteMenu
		return m, nil
	}}
}

func showMatrixMenuToggle(m Model) list.Item {
	name := fmt.Sprintf("Matrix (%s)", m.matrixName())
	return item{name: name, onSelect: func(m Model) (Model, tea.Cmd) {
		m.showMatrixMenu = !m.showMatrixMenu
		return m, nil
	}}
}

func getPaletteMenuItems() []list.Item {
	return []list.Item{standardPalette(), filePalette(), lospecPalette()}
}

func standardPalette() list.Item {
	return item{name: "Basic", onSelect: func(m Model) (Model, tea.Cmd) {
		m.state = Palette
		m.Palette = m.Palette.SetState(palette.Basic)
		return m, io.StartRenderCmd
	}}
}

func filePalette() list.Item {
	return item{name: "From File", onSelect: func(m Model) (Model, tea.Cmd) {
		m.state = Palette
		m.Palette = m.Palette.SetState(palette.FromFile)
		return m, io.StartRenderCmd
	}}
}

func lospecPalette() list.Item {
	return item{name: "From Lospec", onSelect: func(m Model) (Model, tea.Cmd) {
		m.state = Palette
		m.Palette = m.Palette.SetState(palette.Lospec)
		return m, io.StartRenderCmd
	}}
}
