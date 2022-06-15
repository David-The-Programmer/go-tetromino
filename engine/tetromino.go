package engine

import (
	"math/rand"
	"time"
)

// Tetrimino matrices follow the NES Rotation Systems (https://strategywiki.org/wiki/Tetris/Rotation_systems)

var iTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(ITetromino), int(ITetromino), int(ITetromino), int(ITetromino)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var jTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space)},
	{int(JTetromino), int(JTetromino), int(JTetromino)},
	{int(Space), int(Space), int(JTetromino)},
}

var lTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space)},
	{int(LTetromino), int(LTetromino), int(LTetromino)},
	{int(LTetromino), int(Space), int(Space)},
}

var oTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(OTetromino), int(OTetromino), int(Space)},
	{int(Space), int(OTetromino), int(OTetromino), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var sTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space)},
	{int(Space), int(STetromino), int(STetromino)},
	{int(STetromino), int(STetromino), int(Space)},
}

var tTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space)},
	{int(TTetromino), int(TTetromino), int(TTetromino)},
	{int(Space), int(TTetromino), int(Space)},
}

var zTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space)},
	{int(ZTetromino), int(ZTetromino), int(Space)},
	{int(Space), int(ZTetromino), int(ZTetromino)},
}

func randTetromino() [][]int {
	t := [][][]int{
		iTetrominoMatrix,
		jTetrominoMatrix,
		lTetrominoMatrix,
		oTetrominoMatrix,
		sTetrominoMatrix,
		tTetrominoMatrix,
		zTetrominoMatrix,
	}
	rand.Seed(time.Now().Unix())
	return t[rand.Intn(len(t))]
}

type direction int

const (
	clockwise direction = iota
	antiClockwise
)

// rotate rotates m, a square matrix, in place by 90 degrees in the specified rotation direction (clockwise or anticlockwise)
func rotate(m [][]int, d direction) {
	// transpose matrix
	for row := 0; row < len(m); row++ {
		for col := row; col < len(m[row]); col++ {
			temp := m[row][col]
			m[row][col] = m[col][row]
			m[col][row] = temp
		}
	}
	if d == clockwise {
		// flip the matrix horizontally if clockwise rotation specified
		for row := 0; row < len(m); row++ {
			for col := 0; col < len(m)/2; col++ {
				temp := m[row][col]
				m[row][col] = m[row][len(m)-1-col]
				m[row][len(m)-1-col] = temp
			}
		}
		return
	}
	if d == antiClockwise {
		// flip the matrix vertically if anticlockwise rotation specified
		for row := 0; row < len(m)/2; row++ {
			for col := 0; col < len(m); col++ {
				temp := m[row][col]
				m[row][col] = m[len(m)-1-row][col]
				m[len(m)-1-row][col] = temp
			}
		}
		return
	}

}
