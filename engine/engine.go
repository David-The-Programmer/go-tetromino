package engine

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type engine struct {
	state gotetromino.State
}

// New returns a new instance of gotetromino.Engine
func New() gotetromino.Engine {
	const (
		numMatrixRows = 20
		numMatrixCols = 10
	)
	e := engine{}
	e.state = emptyMatrix(e.state, numMatrixRows, numMatrixCols)
	e.state = spawnTetromino(e.state)
	e.state = calcGhostTetrominoPos(e.state)
	return &e
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
	s = setClearedLinesRows(s, nil)
	s = setClearedPrevLevel(s, false)
	switch a {
	case gotetromino.None:
		s = moveTetromino(s, 1, 0)
		if !bottomBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			e.state = s
			break
		}
		s = moveTetromino(s, -1, 0)
		// lock tetromino into matrix once collision occurs
		s = lockTetromino(s)
		// check for any lines (full rows)
		rows := findLines(s)
		if len(rows) > 0 {
			// clear any lines if found
			s = clearLines(s, rows)
			s = setClearedLinesRows(s, rows)
			s = incrementLineCount(s, len(rows))
			s = incrementScore(s, calcPts(s.Level, len(rows)))
		}
		if clearedLevel(s) {
			s = setClearedPrevLevel(s, true)
			s = levelUp(s)
		}
		temp := spawnTetromino(s)
		// if unable to spawn new tetromino due to existing pieces already there, game is over
		if overlapExistingTBlock(temp) {
			s = over(s)
		} else {
			s = temp
		}
		e.state = s
	case gotetromino.SoftDrop:
		const maxSteps = 2
		s = moveTetromino(s, maxSteps, 0)
		if bottomBoundaryExceeded(s) || overlapExistingTBlock(s) {
			for bottomBoundaryExceeded(s) || overlapExistingTBlock(s) {
				s = moveTetromino(s, -1, 0)
			}
			// lock tetromino into matrix if collision occurs
			s = lockTetromino(s)
			// check for any lines (full rows)
			rows := findLines(s)
			if len(rows) > 0 {
				// clear any lines if found
				s = clearLines(s, rows)
				s = setClearedLinesRows(s, rows)
				s = incrementLineCount(s, len(rows))
				s = incrementScore(s, calcPts(s.Level, len(rows)))
			}
			if clearedLevel(s) {
				s = setClearedPrevLevel(s, true)
				s = levelUp(s)
			}
			temp := spawnTetromino(s)
			// if unable to spawn new tetromino due to existing pieces already there, game is over
			if overlapExistingTBlock(temp) {
				s = over(s)
			} else {
				s = temp
			}
		}
		e.state = s
	case gotetromino.HardDrop:
		for !bottomBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			s = moveTetromino(s, 1, 0)
		}
		s = moveTetromino(s, -1, 0)
		// lock tetromino into matrix once collision occurs
		s = lockTetromino(s)
		// check for any lines (full rows)
		rows := findLines(s)
		if len(rows) > 0 {
			// clear any lines if found
			s = clearLines(s, rows)
			s = setClearedLinesRows(s, rows)
			s = incrementLineCount(s, len(rows))
			s = incrementScore(s, calcPts(s.Level, len(rows)))
		}
		if clearedLevel(s) {
			s = setClearedPrevLevel(s, true)
			s = levelUp(s)
		}
		temp := spawnTetromino(s)
		// if unable to spawn new tetromino due to existing pieces already there, game is over
		if overlapExistingTBlock(temp) {
			s = over(s)
		} else {
			s = temp
		}
		e.state = s
	case gotetromino.Left:
		s = moveTetromino(s, 0, -1)
		if !leftBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			e.state = s
		}
	case gotetromino.Right:
		s = moveTetromino(s, 0, 1)
		if !rightBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			e.state = s
		}
	case gotetromino.RotateCW:
		s = rotateTetrimino(s, clockwise)
		if !leftBoundaryExceeded(s) && !rightBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			e.state = s
		}
	case gotetromino.RotateACW:
		s = rotateTetrimino(s, antiClockwise)
		if !leftBoundaryExceeded(s) && !rightBoundaryExceeded(s) && !overlapExistingTBlock(s) {
			e.state = s
		}
	}
	e.state = calcGhostTetrominoPos(e.state)
}

// Reset resets the state of the game back to its initial state
func (e *engine) Reset() {
	e.state = spawnTetromino(e.state)
	e.state = emptyMatrix(e.state, len(e.state.Matrix), len(e.state.Matrix[0]))
	e.state = calcGhostTetrominoPos(e.state)
	e.state.Score = 0
	e.state.Over = false
	e.state.Level = 0
	e.state.ClearedPrevLevel = false
	e.state.LineCount = 0
	e.state.ClearedLinesRows = nil
	e.state.Bag = nil
}

// leftBoundaryExceeded returns true if CurrentTetrominoPos exceeds the left boundary
func leftBoundaryExceeded(s gotetromino.State) bool {
	tMatrixCol := s.CurrentTetrominoPos[1]
	for r := 0; r < len(s.CurrentTetromino); r++ {
		for c := 0; c < len(s.CurrentTetromino); c++ {
			if s.CurrentTetromino[r][c] == int(Space) {
				continue
			}
			tBlockCol := tMatrixCol + c
			if tBlockCol < 0 {
				return true
			}
		}
	}
	return false
}

// rightBoundaryExceeded returns true if CurrentTetrominoPos exceeds the right boundary
func rightBoundaryExceeded(s gotetromino.State) bool {
	tMatrixCol := s.CurrentTetrominoPos[1]
	for r := 0; r < len(s.CurrentTetromino); r++ {
		for c := 0; c < len(s.CurrentTetromino); c++ {
			if s.CurrentTetromino[r][c] == int(Space) {
				continue
			}
			tBlockCol := tMatrixCol + c
			if tBlockCol > len(s.Matrix[0])-1 {
				return true
			}
		}
	}
	return false
}

// bottomBoundaryExceeded returns true if CurrentTetrominoPos exceeds the bottom boundary
func bottomBoundaryExceeded(s gotetromino.State) bool {
	tMatrixRow := s.CurrentTetrominoPos[0]
	for r := 0; r < len(s.CurrentTetromino); r++ {
		for c := 0; c < len(s.CurrentTetromino); c++ {
			if s.CurrentTetromino[r][c] == int(Space) {
				continue
			}
			tBlockRow := tMatrixRow + r
			if tBlockRow > len(s.Matrix)-1 {
				return true
			}
		}
	}
	return false
}

// overlapExistingTBlock returns true if CurrentTetromino overlaps with another tetromino in the Matrix
// overlapExistingTBlock ignores any CurrentTetromino block that exceeds left, right or bottom boundary for overlapping
func overlapExistingTBlock(s gotetromino.State) bool {
	tMatrixRow := s.CurrentTetrominoPos[0]
	tMatrixCol := s.CurrentTetrominoPos[1]
	for r := 0; r < len(s.CurrentTetromino); r++ {
		for c := 0; c < len(s.CurrentTetromino); c++ {
			if s.CurrentTetromino[r][c] == int(Space) {
				continue
			}
			tBlockRow := tMatrixRow + r
			tBlockCol := tMatrixCol + c
			if tBlockRow > len(s.Matrix)-1 || tBlockCol < 0 || tBlockCol > len(s.Matrix[0])-1 {
				continue
			}
			if s.Matrix[tBlockRow][tBlockCol] != int(Space) {
				return true
			}
		}
	}
	return false
}

// duplicate returns a deep copy of the given state
func duplicate(s gotetromino.State) gotetromino.State {
	state := gotetromino.State{}
	for row := 0; row < len(s.CurrentTetromino); row++ {
		newRow := []int{}
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			newRow = append(newRow, s.CurrentTetromino[row][col])
		}
		state.CurrentTetromino = append(state.CurrentTetromino, newRow)
	}
	for i := 0; i < len(s.CurrentTetrominoPos); i++ {
		state.CurrentTetrominoPos = append(state.CurrentTetrominoPos, s.CurrentTetrominoPos[i])
	}
	for row := 0; row < len(s.Matrix); row++ {
		newRow := []int{}
		for col := 0; col < len(s.Matrix[row]); col++ {
			newRow = append(newRow, s.Matrix[row][col])
		}
		state.Matrix = append(state.Matrix, newRow)
	}
	state.Score = s.Score
	state.Over = s.Over
	state.Level = s.Level
	state.ClearedPrevLevel = s.ClearedPrevLevel
	state.LineCount = s.LineCount
	state.ClearedLinesRows = append([]int{}, s.ClearedLinesRows...)
	state.Bag = append([][][]int{}, s.Bag...)
	state.GhostTetrominoPos = append([]int{}, s.GhostTetrominoPos...)

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
			if matrixRow < 0 || matrixCol < 0 || matrixRow > len(state.Matrix)-1 || matrixCol > len(state.Matrix[0])-1 {
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
	if len(state.Bag) == 0 {
		state = generateBag(state)
	}
	state, tetromino := pickTetromino(state)
	state.CurrentTetromino = tetromino
	state.CurrentTetrominoPos = tetrominoStartPos(state.CurrentTetromino, state.Matrix)
	return state
}

// calcGhostTetrominoPos returns a new state, which comprises of the given state with the new position of the ghost tetromino
func calcGhostTetrominoPos(s gotetromino.State) gotetromino.State {
	temp := duplicate(s)
	for !bottomBoundaryExceeded(temp) && !overlapExistingTBlock(temp) {
		temp = moveTetromino(temp, 1, 0)
	}
	temp = moveTetromino(temp, -1, 0)
	state := duplicate(s)
	state.GhostTetrominoPos = append([]int{}, temp.CurrentTetrominoPos...)
	return state
}

// over returns a new state, which comprises of the given state that has the Over field set to true
func over(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	state.Over = true
	return state
}

// findLines check the given state for lines and returns slice of rows which have lines in the matrix
// empty slice is returned if no lines found
func findLines(s gotetromino.State) []int {
	rows := []int{}
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
			rows = append(rows, row)
		}
	}
	return rows
}

// setClearedLinesRows returns a new state, comprises of the given state with ClearedLinesRows being set to given rows of cleared lines
func setClearedLinesRows(s gotetromino.State, rows []int) gotetromino.State {
	state := duplicate(s)
	state.ClearedLinesRows = append([]int{}, rows...)
	return state
}

// clearLines returns a new state, which comprises of the given state where the specified rows containing lines are cleared
// rows is expected to have at least length of 1, or it will panic
func clearLines(s gotetromino.State, rows []int) gotetromino.State {
	state := duplicate(s)
	for row := rows[0]; row <= rows[len(rows)-1]; row++ {
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

// incrementLineCount returns a new state comprising of the given state with LineCount incremented with the given no. lines cleared
func incrementLineCount(s gotetromino.State, numLinesCleared int) gotetromino.State {
	state := duplicate(s)
	state.LineCount += numLinesCleared
	return state
}

// clearedLevel return true if the current level of the given state has been cleared
func clearedLevel(s gotetromino.State) bool {
	const linesToClear = 10
	return len(s.ClearedLinesRows) != 0 && s.LineCount%linesToClear == 0 && s.LineCount != 0
}

// setClearedPrevLevel sets the value of the ClearedPrevLevel flag field of the given state
func setClearedPrevLevel(s gotetromino.State, v bool) gotetromino.State {
	state := duplicate(s)
	state.ClearedPrevLevel = v
	return state
}

// levelUp returns a new state comprising of the given state with Level incremented by 1
func levelUp(s gotetromino.State) gotetromino.State {
	state := duplicate(s)
	state.Level += 1
	return state
}

// calcPts returns the points to be added given the current level and no. lines cleared
func calcPts(level int, linesCleared int) int {
	ptsToLinesCleared := []int{40, 100, 300, 1200}
	return (level + 1) * ptsToLinesCleared[linesCleared-1]
}

// incrementScore returns a new state comprising of the given state with Score incremented by the given points
func incrementScore(s gotetromino.State, pts int) gotetromino.State {
	state := duplicate(s)
	state.Score += pts
	return state
}

// tetrominoStartPos returns the starting position of a tetromino
func tetrominoStartPos(tetromino [][]int, matrix [][]int) []int {
	return []int{
		0,
		(len(matrix[0]) - 2 - len(tetromino)) / 2,
	}
}

// TODO: Make all board individual components
// TODO: Fix the spawn pos of tetromino
// TODO: Text alignment of stats & controls
// TODO: Fix clearing of lines
// TODO: Fix rotation such that pieces against the boundary can still be rotated (wall kickback)
// TODO: Refactor UI to not use other UI components as throwaways
// TODO: Finish having next tetromino
// TODO: Finish reset of state
// TODO: Refactor and simplify current store & key event listener to one event loop?
// TODO: Need to fix bug where game freezes after too quick of key presses

// TODO: Make comment terms that relate to the code be of the specific constant/field
