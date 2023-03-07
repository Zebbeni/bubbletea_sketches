package settings

type State int

const (
	Menu State = iota
	Colors
	Sampling
	Characters
)

var States = []State{
	Menu,
	Colors,
	Sampling,
	Characters,
}
