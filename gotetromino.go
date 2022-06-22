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
	NotifyObservers()
}

type Observer interface {
	Notify()
}

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

type Player interface {
	Subject
	Action() Action
}

// Interaction is the action a user does to use the game app
type Interaction int

const (
	Exit Interaction = iota
	Restart
)

type User interface {
	Subject
	Interaction() Interaction
}

type EngineService interface {
	Subject
	Step(a Action)
	Reset()
	State() State
	Stop()
}

type UI interface {
	Render()
}

// Matrix U.I component, Render, Listen, Re-render, Listen, ...
// Not every U.I component would need to listen, but all need to render
// U.I component listen to different things, like state, user actions/interaction
// Need game component, child components (matrix component, stats component, next component), instructions
