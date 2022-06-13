package engine

import (
	"sync"
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type engine struct {
	state       gotetromino.State
	action      <-chan gotetromino.Action
	stateChange chan gotetromino.State
	ticker      *time.Ticker
	tickerDelay time.Duration
	stop        chan bool
}

// New returns a new instance of gotetromino.Engine
func New(numRows int, numCols int) gotetromino.Engine {
	e := engine{}
	e.state = emptyMatrix(e.state, numRows, numCols)
	// set current tetromino & its position
	// TODO: Refactor this part
	e.state.CurrentTetromino = randTetromino()
	e.state.CurrentTetrominoPos = tetrominoStartPos(e.state.CurrentTetromino, e.state.Matrix)

	var coreEngine gotetromino.Engine = &e
	return coreEngine
}

// Start runs the engine (launches internal processes)
// receives a channel to receive actions and returns a channel to receive State when it changes
// engine will ignore any action once gameover, and will only continue to execute actions when Reset is invoked
func (e *engine) Start(a <-chan gotetromino.Action) <-chan gotetromino.State {
	// need delay to simulate moving onto next frame for renderer
	// TODO: Need to have different delays between falling and movement of pieces
	delay := 100 * time.Millisecond
	e.tickerDelay = 3 * delay
	e.ticker = time.NewTicker(e.tickerDelay)
	e.action = a
	e.stateChange = make(chan gotetromino.State)
	e.stop = make(chan bool)

	mutex := sync.Mutex{}

	// launch goroutine to handle updating of state from action and ticker events
	go func() {
		// send initial state
		e.stateChange <- e.state
		for {
			select {
			case <-e.stop:
				e.ticker.Stop()
				return
			case <-e.ticker.C:
				mutex.Lock()
				s := duplicate(e.state)
				mutex.Unlock()
				// do not continue updating state if game is over
				if s.Over {
					continue
				}
				// make current tetromino fall every 150ms
				s = moveTetromino(s, 1, 0)
				if !collision(s) {
					mutex.Lock()
					e.state = s
					mutex.Unlock()
					e.stateChange <- e.state
					continue
				}
				s = moveTetromino(s, -1, 0)
				// lock tetromino into matrix once collision occurs
				s = lockTetromino(s)
				temp := spawnTetromino(s)
				// if unable to spawn new tetromino due to existing pieces already there, game is over
				if collision(temp) {
					s = over(s)
				} else {
					s = temp
				}
				mutex.Lock()
				e.state = s
				mutex.Unlock()
				e.stateChange <- e.state
			case action := <-e.action:
				mutex.Lock()
				s := duplicate(e.state)
				mutex.Unlock()
				// do not continue updating state if game is over
				if s.Over {
					continue
				}
				// set the change in position of CurrentTetromino according to action
				switch action {
				case gotetromino.Drop:
					for !collision(s) {
						s = moveTetromino(s, 1, 0)
					}
					s = moveTetromino(s, -1, 0)
					// lock tetromino into matrix once collision occurs
					s = lockTetromino(s)
					temp := spawnTetromino(s)
					// if unable to spawn new tetromino due to existing pieces already there, game is over
					if collision(temp) {
						s = over(s)
					} else {
						s = temp
					}
					mutex.Lock()
					e.state = s
					mutex.Unlock()
					time.Sleep(delay)
					e.stateChange <- e.state
				case gotetromino.Left:
					s = moveTetromino(s, 0, -1)
					if !collision(s) {
						mutex.Lock()
						e.state = s
						mutex.Unlock()
						time.Sleep(delay)
						e.stateChange <- e.state
					}
				case gotetromino.Right:
					s = moveTetromino(s, 0, 1)
					if !collision(s) {
						mutex.Lock()
						e.state = s
						mutex.Unlock()
						time.Sleep(delay)
						e.stateChange <- e.state
					}
					// TODO: Complete logic to rotate tetromino
					// case gotetromino.Rotate:
				}
			}
		}
	}()

	return e.stateChange
}

// Stop ends the running of the engine
func (e *engine) Stop() {
	e.stop <- true
}

// Reset resets the state of the game back to its initial state
func (e *engine) Reset() {
	e.state = spawnTetromino(e.state)
	e.state = emptyMatrix(e.state, len(e.state.Matrix), len(e.state.Matrix[0]))
	e.state.Over = false
	e.state.Score = 0
	e.ticker.Reset(e.tickerDelay)
}

// collision returns true if non-space blocks of the matrix overlap with non-space blocks of the tetromino in the given state
func collision(s gotetromino.State) bool {
	x := s.CurrentTetrominoPos[1]
	y := s.CurrentTetrominoPos[0]
	for i := 0; i < len(s.CurrentTetromino); i++ {
		for j := 0; j < len(s.CurrentTetromino[i]); j++ {
			if s.CurrentTetromino[i][j] != int(Space) && s.Matrix[y+i][x+j] != int(Space) {
				return true
			}
		}
	}
	return false

}

// duplicate returns a deep copy of the given state
func duplicate(s gotetromino.State) gotetromino.State {
	state := gotetromino.State{}

	// copy CurrentTetromino
	for row := 0; row < len(s.CurrentTetromino); row++ {
		newRow := []int{}
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			newRow = append(newRow, s.CurrentTetromino[row][col])
		}
		state.CurrentTetromino = append(state.CurrentTetromino, newRow)
	}

	// copy CurrentTetrominoPos
	for i := 0; i < len(s.CurrentTetrominoPos); i++ {
		state.CurrentTetrominoPos = append(state.CurrentTetrominoPos, s.CurrentTetrominoPos[i])
	}

	// copy Matrix
	for row := 0; row < len(s.Matrix); row++ {
		newRow := []int{}
		for col := 0; col < len(s.Matrix[row]); col++ {
			newRow = append(newRow, s.Matrix[row][col])
		}
		state.Matrix = append(state.Matrix, newRow)
	}

	// copy Score
	state.Score = s.Score

	// copy Over
	state.Over = s.Over

	return state
}

// emptyMatrix returns a new state, which comprises of the given state with an empty matrix of with numRows & numCols
func emptyMatrix(s gotetromino.State, numRows int, numCols int) gotetromino.State {
	matrix := [][]int{}
	// init matrix as space
	for r := 0; r < numRows; r++ {
		matrix = append(matrix, []int{})
		for c := 0; c < numCols; c++ {
			matrix[r] = append(matrix[r], int(Space))
		}
	}
	// encode boundaries into matrix
	for i := 0; i < len(matrix); i++ {
		// leftmost boundary
		matrix[i][0] = int(Boundary)
		// rightmost boundary
		matrix[i][len(matrix[i])-1] = int(Boundary)
	}
	// bottom boundary
	for i := 0; i < len(matrix[len(matrix)-1]); i++ {
		matrix[len(matrix)-1][i] = int(Boundary)
	}
	state := duplicate(s)
	state.Matrix = matrix
	return state

}

// lockTetromino returns a new state, which comprises of the given state that has its current tetromino locked in the matrix
func lockTetromino(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	x := state.CurrentTetrominoPos[1]
	y := state.CurrentTetrominoPos[0]
	for i := 0; i < len(state.CurrentTetromino); i++ {
		for j := 0; j < len(state.CurrentTetromino[i]); j++ {
			matrixRow := y + i
			matrixCol := x + j
			if matrixRow > len(state.Matrix)-1 || matrixCol > len(state.Matrix[0])-1 {
				continue
			}

			// only override space blocks with tetromino blocks
			if state.Matrix[matrixRow][matrixCol] == int(Space) {
				state.Matrix[matrixRow][matrixCol] = state.CurrentTetromino[i][j]
			}

		}
	}
	return state

}

// moveTetromino returns a new state, which comprises of the given state that has its current tetromino moved by numRows & numCols
func moveTetromino(s gotetromino.State, numRows int, numCols int) gotetromino.State {
	state := duplicate(s)
	state.CurrentTetrominoPos[0] += numRows
	state.CurrentTetrominoPos[1] += numCols
	return state
}

// spawnTetromino returns a new state, which comprises of the given state that has new tetromino is spawned
func spawnTetromino(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	state.CurrentTetromino = randTetromino()
	state.CurrentTetrominoPos = tetrominoStartPos(state.CurrentTetromino, state.Matrix)
	return state
}

// tetrominoStartPos returns the starting position of a tetromino
func tetrominoStartPos(tetromino [][]int, matrix [][]int) []int {
	return []int{
		0,
		(len(matrix[0]) - 2 - len(tetromino)) / 2,
	}
}

// over returns a new state, which comprises of the given state that has the Over field set to true
func over(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	state.Over = true
	return state
}

// TODO: Finish clearing of tetromino blocks
// TODO: Finish rotation of tetromino

// TODO: Make comment terms that relate to the code be of the specific constant/field
