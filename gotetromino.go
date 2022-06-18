package gotetromino

type UI interface {
	Run()
}

type Engine interface {
	State() State
	Step(a Action)
	Reset()
}

type State struct {
	CurrentTetromino    [][]int
	CurrentTetrominoPos []int
	Matrix              [][]int
	Score               int
	Over                bool
	Level               int
	LineCount           int
	ClearedLinesRows    []int
}

type Action int

const (
	None Action = iota
	SoftDrop
	HardDrop
	Left
	Right
	RotateCW
	RotateACW
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
