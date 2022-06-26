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
	// ClearedPrevLevel equals to true if level has been incremented (cleared previous level) after given action
	ClearedPrevLevel bool
	// LineCount reflects the total no. lines cleared since start of the game
	LineCount int
	// ClearedLinesRows contains all rows of cleared lines from the previous state, after the given action
	ClearedLinesRows []int
}

type Service interface {
	Start()
	Stop()
}

type StateEventListener interface {
	Service
	Attach(h StateEventHandler)
	Detach(h StateEventHandler)
	Publish()
}

type StateEventHandler interface {
	HandleNewState(s State)
}

// Store manages changes to the state
type Store interface {
	Engine
	StateEventListener
}

type KeyEventListener interface {
	Service
	Attach(h KeyEventHandler)
	Detach(h KeyEventHandler)
	Publish()
}

type KeyEventHandler interface {
	HandleNewKey(k Key)
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

type UI interface {
	Run()
}
