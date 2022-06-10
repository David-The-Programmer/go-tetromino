package gotetromino

type Game interface {
	// Run runs tetris
	Run()
}

type Engine interface {
	Start(a <-chan Action) <-chan State
	Stop()
	Reset()
}

// Main game loop
/*
   currentState := engine.State()
   renderer.Render(currentState)
   engine.Step()
*/

// Need to have internal timer, every 150ms, (30*0.15 frames if 30 FPS, 60*0.15 frames if 60 FPS), move CurrentTetromino down
// One Step is one frame of tetris

/*
func (ge *GameEngine) Step() {
    if keyEvent {
        player.SetAction(keyEvent)
    }
    if player.Action() == drop {
        if !collision() {
            updateCurrentTetrominoPos()
            return
        }
    }
    if player.Action() == left {
        if !collision() {
            updateCurrentTetrominoPos()
            return
        }
    }
    if player.Action() == right {
        if !collision() {
            updateCurrentTetrominoPos()
            return
        }
    }
    if player.Action() == rotate {
        if !collision() {
            updateCurrentTetrominoPos()
            return
        }
    }
    if timer == 150ms {
        if !collision() {
            updateCurrentTetrominoPos()
        }
    }
}
*/

// 1) if any key event occurs, set player action for current Step
// 2) Upon player action, do collision checks, update CurrentTetrominoPos if possible render state
// 4) Do collision check and cause CurrentTetromino to fall if possible
// 5) Render State

type State struct {
	CurrentTetromino    [][]int
	CurrentTetrominoPos []int
	Matrix              [][]int
	Score               int
	Over                bool
}

type Action int

const (
	Left Action = iota
	Right
	Drop
	Rotate
)

type Renderer interface {
	Render(s State)
	Stop()
}
