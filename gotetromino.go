package gotetromino

// Engine contains all core game logic
type Engine interface {
	State() State
	Step(a Action)
	Reset()
}

// State is the state of the tetris game after a given action
type State struct {
	// CurrentTetromino reflects the current state of the tetromino after the given action
	CurrentTetromino [][]int
	// CurrentTetrominoPos reflects the current position of the tetromino after the given action
	CurrentTetrominoPos []int
	// Matrix reflects the current state of the matrix after the given action
	Matrix [][]int
	// Score reflects the current score since the start of the game
	Score int
	// Over equals to true if game is over after given action
	Over bool
	// Level reflects the current level since the start of the game
	Level int
	// LineCount reflects the total no. lines cleared since start of the game
	LineCount int
	// ClearedLinesRows contains all rows of cleared lines from the previous state, after the given action
	ClearedLinesRows []int
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	NotifyAll()
}

type Observer interface {
    Subscribe(s Subject)
	Notify(v any)
}

type KeyEventListener interface {
	Subject
	Listen()
	Stop()
}

type Key int

const (
	Esc Key = iota
	DownArrow
	LeftArrow
	RightArrow
	SpaceBar
	R
	X
	Z
)

// Action is the in-game action a player makes to play the tetris
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

// Interaction is the action a user does to use the game app
type Interaction int

const (
	Exit Interaction = iota
	Restart
)

type EngineService interface {
	Engine
	Subject
	Stop()
}

type UI interface {
	Run()
}
