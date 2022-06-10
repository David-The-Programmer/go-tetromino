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
	stop        chan bool
	mutex       *sync.Mutex
}

// New returns a new instance of gotetromino.Engine
func New(numRows int, numCols int) gotetromino.Engine {
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
	e.mutex = &sync.Mutex{}

	var coreEngine gotetromino.Engine = &e
	return coreEngine
}

// Start runs the engine (launches internal processes)
// receives a channel to receive actions and returns a channel to receive State when it changes
func (e *engine) Start(a <-chan gotetromino.Action) <-chan gotetromino.State {
	// need delay to simulate moving onto next frame for renderer
	// TODO: Need to have different delays between falling and movement of pieces
	delay := 100 * time.Millisecond
	e.ticker = time.NewTicker(3 * delay)
	e.action = a
	e.stateChange = make(chan gotetromino.State)
	e.stop = make(chan bool)

	// launch goroutine to handle updating of state from action and ticker events
	go func() {
		// send initial state
		e.stateChange <- e.state
		for {
			select {
			case <-e.stop:
				// release all resources when engine is stopped
				e.ticker.Stop()
				return
			case <-e.ticker.C:
				// make current tetromino fall every 150ms
				e.mutex.Lock()
				e.state.CurrentTetrominoPos[0] += 1
				e.mutex.Unlock()
				e.stateChange <- e.state
			case action := <-e.action:
				// set the change in position of CurrentTetromino according to action
				switch action {
				// case gotetromino.Drop:
				// TODO: Complete logic to drop tetromino
				// e.mutex.Lock()
				// e.state.CurrentTetrominoPos[0] += 1
				// e.mutex.Unlock()
				case gotetromino.Left:
					e.mutex.Lock()
					e.state.CurrentTetrominoPos[1] += -1
					e.mutex.Unlock()
					time.Sleep(delay)
					e.stateChange <- e.state
				case gotetromino.Right:
					e.mutex.Lock()
					e.state.CurrentTetrominoPos[1] += 1
					e.mutex.Unlock()
					time.Sleep(delay)
					e.stateChange <- e.state
					// case gotetromino.Rotate:
				}
			}
		}
	}()
	// TODO: Replace timer based game over for actual game over
	go func() {
		time.Sleep(2 * time.Second)
		e.mutex.Lock()
		e.state.Over = true
		e.mutex.Unlock()
		e.stop <- e.state.Over
	}()

	return e.stateChange
}

// Stop ends the running of the engine
func (e *engine) Stop() {
	e.stop <- true
}

// TODO: Finish Reset
// Reset sets the state of the game back its inital state (before Start was invoked)
func (e *engine) Reset() {}
