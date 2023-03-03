package settings

type State int

const (
	Main State = iota
	Colors
	Interpolation
	Characters
)

var States = []State{
	Main,
	Colors,
	Interpolation,
	Characters,
}
