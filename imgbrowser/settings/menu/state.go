package menu

type State int

const (
	Main State = iota
	Interpolation
)

var StateNames = map[State]string{
	Main:          "Main",
	Interpolation: "Interpolation",
}
