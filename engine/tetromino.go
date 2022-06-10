package engine

import (
	"math/rand"
	"time"
)

var iTetrominoMatrix = [][]int{
	{int(ITetromino), int(ITetromino), int(ITetromino), int(ITetromino)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var jTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(JTetromino)},
	{int(Space), int(Space), int(Space), int(JTetromino)},
	{int(Space), int(Space), int(JTetromino), int(JTetromino)},
}

var lTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space), int(Space)},
	{int(LTetromino), int(Space), int(Space), int(Space)},
	{int(LTetromino), int(Space), int(Space), int(Space)},
	{int(LTetromino), int(LTetromino), int(Space), int(Space)},
}

var oTetrominoMatrix = [][]int{
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(OTetromino), int(OTetromino), int(Space)},
	{int(Space), int(OTetromino), int(OTetromino), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var sTetrominoMatrix = [][]int{
	{int(Space), int(STetromino), int(STetromino), int(Space)},
	{int(STetromino), int(STetromino), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var tTetrominoMatrix = [][]int{
	{int(TTetromino), int(TTetromino), int(TTetromino), int(Space)},
	{int(Space), int(TTetromino), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
}

var zTetrominoMatrix = [][]int{
	{int(ZTetromino), int(ZTetromino), int(Space), int(Space)},
	{int(Space), int(ZTetromino), int(ZTetromino), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
	{int(Space), int(Space), int(Space), int(Space)},
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
