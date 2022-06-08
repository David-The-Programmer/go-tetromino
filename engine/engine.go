package engine

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type engine struct {
	state  gotetromino.State
	player gotetromino.Player
	ticker *time.Ticker
}

func New(numRows int, numCols int) *engine {
	e := engine{}
	// init matrix as space
	for r := 0; r < numRows; r++ {
		e.state.Matrix = append(e.state.Matrix, []int{})
		for c := 0; c < numCols; c++ {
			e.state.Matrix[r] = append(e.state.Matrix[r], int(Space))
		}
	}
	// encode boundaries into matrix
	for i := 0; i < len(e.state.Matrix); i++ {
		// leftmost boundary
		e.state.Matrix[i][0] = int(Boundary)
		// rightmost boundary
		e.state.Matrix[i][len(e.state.Matrix[i])-1] = int(Boundary)
	}
	// bottom boundary
	for i := 0; i < len(e.state.Matrix[len(e.state.Matrix)-1]); i++ {
		e.state.Matrix[len(e.state.Matrix)-1][i] = int(Boundary)
	}

	// set current tetromino & its position
	e.state.CurrentTetromino = randTetromino()
	e.state.CurrentTetrominoPos = []int{
		(len(e.state.Matrix) - 2 - len(e.state.CurrentTetromino)) / 2,
		(len(e.state.Matrix[0]) - 2 - len(e.state.CurrentTetromino)) / 2,
	}

	// set ticker to tick every 150ms
	e.ticker = time.NewTicker(150 * time.Millisecond)

	return &e
}

func (e *engine) Step() {
	// make current tetromino fall every 150ms
    done := make(chan bool)
    go func() {
        for {
            select {
            case <- done:
                return
            case <- e.ticker.C:
                e.state.CurrentTetrominoPos[0] += 1
            }
        }
    }()
    go func() {
        time.Sleep(2*time.Second)
        e.ticker.Stop()
        e.state.Over = true
        done <- true
    }()

}

func (e *engine) Reset() {

}

func (e *engine) State() gotetromino.State {
	return e.state
}

func (e *engine) Player(p gotetromino.Player) {
	e.player = p
}
