package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type nextBoard struct {
	*board
}

func newNextBoard(sc tcell.Screen, s gotetromino.State) *nextBoard {
	b := newBoard(sc)
	b.SetBorderColour(tcell.ColorGrey)
	b.SetTitle("Next")
	b.SetTitleColour(tcell.ColorGrey)
	largestTetromino := s.Bag[0]
	for i := range s.Bag {
		if len(largestTetromino) < len(s.Bag[i]) {
			largestTetromino = s.Bag[i]
		}
	}
	rows := len(largestTetromino)
	cols := len(largestTetromino[0])
	rowHeights := []int{}
	colWidths := []int{}
	for row := 0; row < rows; row++ {
		rowHeights = append(rowHeights, 1)
	}
	b.SetContentRowHeights(rowHeights...)
	for col := 0; col < cols; col++ {
		colWidths = append(colWidths, 1)
	}
	b.SetContentColWidths(colWidths...)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			blk := newBlock(sc)
			blk.Hide(true)
			b.AddContent(blk, row, col, 1, 1)
		}
	}
	return &nextBoard{
		board: b,
	}
}

func (n *nextBoard) SetNextTetromino(s gotetromino.State) {
	rows := len(n.GetContentRowHeights())
	cols := len(n.GetContentColWidths())
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			blk := n.GetContent(row, col).(*gridComponent).UI.(*block)
			blk.Hide(true)
		}
	}
	nextTetromino := s.Bag[0]
	for row := 0; row < len(nextTetromino); row++ {
		for col := 0; col < len(nextTetromino[row]); col++ {
			blk := n.GetContent(row, col).(*gridComponent).UI.(*block)
			if engine.Block(nextTetromino[row][col]) == engine.Space {
				blk.Hide(true)
				continue
			}
			blk.Hide(false)
			blk.SetColour(colourForBlock(engine.Block(nextTetromino[row][col])))
		}
	}
}
