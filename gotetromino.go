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
	// Bag is a bag of randomly shuffled tetrominos
	Bag [][][]int
	// GhostTetrominoPos reflects the position of the CurrentTetromino after it is drop from its current position
	GhostTetrominoPos []int
	// Reset is true if game was reset
	Reset bool
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

type UI interface {
	SetPos(x, y int)
	SetDimensions(w, h int)
	Render()
}

type App interface {
	Run()
}
