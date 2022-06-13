package gotetromino

type Game interface {
	Run()
}

type Engine interface {
    State() State
	Step(a Action)
	Reset()
}

/*
// TODO: Abstract out all timing/animation logic to UI/renderer
type Engine interface {
    Step(a Action) State
    Reset()
}
*/

type State struct {
	CurrentTetromino    [][]int
	CurrentTetrominoPos []int
	Matrix              [][]int
	Score               int
	Over                bool
}

type Action int

const (
	SoftDrop Action = iota
	HardDrop
	Left
	Right
	Rotate
)

type Interaction int

const (
	Exit Interaction = iota
	Restart
)

type Renderer interface {
	Render(s State)
	Stop()
}

type Player interface {
	Action() <-chan Action
}

type User interface {
	Player
	Interaction() <-chan Interaction
}

type UI interface {
	Renderer
	User
}
