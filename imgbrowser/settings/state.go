package settings

type State int

const (
	Main State = iota
	Interpolation
	Dithering
	Palette
	Characters
)

var States = []State{
	Main,
	Interpolation,
	Dithering,
	Palette,
	Characters,
}
