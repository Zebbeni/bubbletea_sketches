package settings

type State int

const (
	None State = iota
	Colors
	Characters
	Size
	Sampling
)

var States = []State{
	Colors,
	Characters,
	Size,
	Sampling,
}

var stateOrder = []State{Colors, Characters, Size, Sampling}
