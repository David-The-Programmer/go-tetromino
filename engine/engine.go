package engine

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type engine struct {
	state gotetromino.State
}

// New returns a new instance of gotetromino.Engine
func New(numRows int, numCols int) gotetromino.Engine {
	e := engine{}
	e.state = emptyMatrix(e.state, numRows, numCols)
	e.state = spawnTetromino(e.state)

	var coreEngine gotetromino.Engine = &e
	return coreEngine
}

// State returns the current game state
func (e *engine) State() gotetromino.State {
	return e.state
}

// Step updates the game state based on the given action
// Step will ignore any action once gameover, and will only continue to execute actions when Reset is invoked
func (e *engine) Step(a gotetromino.Action) {
	// do not continue updating state if game is over
	if e.state.Over {
		return
	}
	s := duplicate(e.state)
	switch a {
	case gotetromino.None:
		s = moveTetromino(s, 1, 0)
		if !collision(s) {
			e.state = s
			return
		}
		s = moveTetromino(s, -1, 0)
		// lock tetromino into matrix once collision occurs
		s = lockTetromino(s)
		// check for any lines (full rows)
		first, last, ok := findLines(s)
		if ok {
			// clear any full rows in the matrix after tetromino is locked
			s = clearLines(s, first, last)
		}

		/*
		   s, numLines := clearLines(s)
		   if numLines % 10 {
		       s = levelUp(s)
		   }
		*/
		temp := spawnTetromino(s)
		// if unable to spawn new tetromino due to existing pieces already there, game is over
		if collision(temp) {
			s = over(s)
		} else {
			s = temp
		}
		e.state = s
	case gotetromino.SoftDrop:
		const maxSteps = 2
		s = moveTetromino(s, maxSteps, 0)
		if collision(s) {
			// rollback on movement if step exceeds the bottom boundary
			for collision(s) {
				s = moveTetromino(s, -1, 0)
			}
			// lock tetromino into matrix if collision occurs
			s = lockTetromino(s)
			// check for any lines (full rows)
			first, last, ok := findLines(s)
			if ok {
				// clear any full rows in the matrix after tetromino is locked
				s = clearLines(s, first, last)
			}
			temp := spawnTetromino(s)
			// if unable to spawn new tetromino due to existing pieces already there, game is over
			if collision(temp) {
				s = over(s)
			} else {
				s = temp
			}
		}
		e.state = s
	case gotetromino.HardDrop:
		for !collision(s) {
			s = moveTetromino(s, 1, 0)
		}
		s = moveTetromino(s, -1, 0)
		// lock tetromino into matrix once collision occurs
		s = lockTetromino(s)
		// check for any lines (full rows)
		first, last, ok := findLines(s)
		if ok {
			// clear any full rows in the matrix after tetromino is locked
			s = clearLines(s, first, last)
		}
		temp := spawnTetromino(s)
		// if unable to spawn new tetromino due to existing pieces already there, game is over
		if collision(temp) {
			s = over(s)
		} else {
			s = temp
		}
		e.state = s
	case gotetromino.Left:
		s = moveTetromino(s, 0, -1)
		if !collision(s) {
			e.state = s
		}
	case gotetromino.Right:
		s = moveTetromino(s, 0, 1)
		if !collision(s) {
			e.state = s
		}
	case gotetromino.RotateCW:
		s = rotateTetrimino(s, clockwise)
		if !collision(s) {
			e.state = s
		}
	case gotetromino.RotateACW:
		s = rotateTetrimino(s, antiClockwise)
		if !collision(s) {
			e.state = s
		}
	}
}

// Reset resets the state of the game back to its initial state
func (e *engine) Reset() {
	e.state = spawnTetromino(e.state)
	e.state = emptyMatrix(e.state, len(e.state.Matrix), len(e.state.Matrix[0]))
	e.state.Over = false
	e.state.Score = 0
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

// over returns a new state, which comprises of the given state that has the Over field set to true
func over(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	state.Over = true
	return state
}

// findLines check the given state for lines and returns the first & last rowIndexes which have lines in the matrix
func findLines(s gotetromino.State) (int, int, bool) {
	rowIndexes := []int{}
	// skip checking bottom boundary
	for row := 0; row < len(s.Matrix)-1; row++ {
		rowHasLine := true
		// skip checking left & right boundaries
		for col := 1; col < len(s.Matrix[row])-1; col++ {
			if s.Matrix[row][col] == int(Space) {
				rowHasLine = false
			}
		}
		if rowHasLine {
			rowIndexes = append(rowIndexes, row)
		}
	}
	if len(rowIndexes) == 0 {
		return 0, 0, false
	}
	return rowIndexes[0], rowIndexes[len(rowIndexes)-1], true
}

// clearLines returns a new state, which comprises of the given state where the rows of specified range of indexes are cleared
func clearLines(s gotetromino.State, firstRowIdx int, lastRowIdx int) gotetromino.State {
	state := duplicate(s)
	for row := firstRowIdx; row <= lastRowIdx; row++ {
		for r := row; r >= 0; r-- {
			for col := 1; col < len(state.Matrix[row])-1; col++ {
				if r == 0 {
					state.Matrix[r][col] = int(Space)
					continue
				}
				state.Matrix[r][col] = state.Matrix[r-1][col]
			}
		}

	}

	return state
}

// rotateTetrimino returns a new state, which comprises of the given state which current tetromino is rotated in specified direction
func rotateTetrimino(s gotetromino.State, d direction) gotetromino.State {
	state := duplicate(s)
	rotate(state.CurrentTetromino, d)
	return state
}

// levelUp calculates the returns a new state

// tetrominoStartPos returns the starting position of a tetromino
func tetrominoStartPos(tetromino [][]int, matrix [][]int) []int {
	return []int{
		0,
		(len(matrix[0]) - 2 - len(tetromino)) / 2,
	}
}

// TODO: Finish levels, put in state (leveling up, speed up game (delay) according to level)
// TODO: Finish scoring
// TODO: Finish returning lines cleared after actions in state, for renderer to render animation or something
// TODO: Finish having next tetromino
// TODO: Finish U.I (show scoring, next tetromino, instructions to restart game, game controls, etc)
// TODO: Finish ghost piece
// TODO: Need to have different delays between falling and movement of pieces

// TODO: Make comment terms that relate to the code be of the specific constant/field
