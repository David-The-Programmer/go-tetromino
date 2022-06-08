package engine

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type engine struct {
	state  gotetromino.State
	player gotetromino.Player
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

	return &e
}

func (e *engine) Step() {

}

func (e *engine) Reset() {

}

func (e *engine) State() gotetromino.State {
	return e.state
}

func (e *engine) Player(p gotetromino.Player) {
	e.player = p
}
