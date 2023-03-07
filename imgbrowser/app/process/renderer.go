package process

import (
	"github.com/lucasb-eyer/go-colorful"

	"github.com/Zebbeni/bubbletea_sketches/imgbrowser/settings"
)

type Renderer struct {
	Settings   settings.Model
	blockFuncs map[rune]blockFunc
}

func New(s settings.Model) Renderer {
	m := Renderer{
		Settings: s,
	}
	m.blockFuncs = m.createBlockFuncs()
	return m
}

type blockFunc func(r1, r2, r3, r4 colorful.Color) (colorful.Color, colorful.Color, float64)

func (m Renderer) createBlockFuncs() map[rune]blockFunc {
	//return map[rune]blockFunc{
	//	'▀': m.calcTop,
	//	'▐': m.calcRight,
	//	'▞': m.calcDiagonal,
	//	'▖': m.calcBotLeft,
	//	'▘': m.calcTopLeft,
	//	'▝': m.calcTopRight,
	//	'▗': m.calcBotRight,
	//}
	return map[rune]blockFunc{
		'▀': m.calcTop,
	}
}
