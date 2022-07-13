package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type matrix struct {
	*board
}

func newMatrix(sc tcell.Screen, s gotetromino.State) *matrix {
	b := newBoard(sc)
	b.SetBorderColour(tcell.ColorGrey)
	b.SetTitle("Tetris")
	b.SetTitleColour(tcell.ColorGrey)
	matrixRows := len(s.Matrix)
	matrixCols := len(s.Matrix[0])
	rowHeights := []int{}
	colWidths := []int{}
	for row := 0; row < matrixRows; row++ {
		rowHeights = append(rowHeights, 1)
	}
	b.SetContentRowHeights(rowHeights...)
	for col := 0; col < matrixCols; col++ {
		colWidths = append(colWidths, 1)
	}
	b.SetContentColWidths(colWidths...)
	for row := 0; row < matrixRows; row++ {
		for col := 0; col < matrixCols; col++ {
			blk := newBlock(sc)
			b.AddContent(blk, row, col, 1, 1)
		}
	}
	return &matrix{
		board: b,
	}
}

func (m *matrix) SetState(s gotetromino.State) {
	for row := 0; row < len(s.Matrix); row++ {
		for col := 0; col < len(s.Matrix[row]); col++ {
			blk := m.GetContent(row, col).(*gridComponent).UI.(*block)
			if engine.Block(s.Matrix[row][col]) == engine.Space {
				blk.Hide(true)
				continue
			}
			blk.Hide(false)
			blk.SetColour(colourForBlock(engine.Block(s.Matrix[row][col])))
		}
	}
	for row := 0; row < len(s.CurrentTetromino); row++ {
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			matrixRow := s.GhostTetrominoPos[0] + row
			matrixCol := s.GhostTetrominoPos[1] + col
			if matrixRow < 0 || matrixRow > len(s.Matrix)-1 {
				continue
			}
			if matrixCol < 0 || matrixCol > len(s.Matrix[0])-1 {
				continue
			}
			if engine.Block(s.Matrix[matrixRow][matrixCol]) != engine.Space {
				continue
			}
			blk := m.GetContent(matrixRow, matrixCol).(*gridComponent).UI.(*block)
			if engine.Block(s.CurrentTetromino[row][col]) == engine.Space {
				continue
			}
			blk.Hide(false)
			blk.SetColour(tcell.ColorGrey)
		}
	}
	for row := 0; row < len(s.CurrentTetromino); row++ {
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			matrixRow := s.CurrentTetrominoPos[0] + row
			matrixCol := s.CurrentTetrominoPos[1] + col
			if matrixRow < 0 || matrixRow > len(s.Matrix)-1 {
				continue
			}
			if matrixCol < 0 || matrixCol > len(s.Matrix[0])-1 {
				continue
			}
			if engine.Block(s.Matrix[matrixRow][matrixCol]) != engine.Space {
				continue
			}
			blk := m.GetContent(matrixRow, matrixCol).(*gridComponent).UI.(*block)
			if engine.Block(s.CurrentTetromino[row][col]) == engine.Space {
				continue
			}
			blk.Hide(false)
			blk.SetColour(colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
		}
	}
}
