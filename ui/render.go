package ui

import (
	"fmt"
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

const (
	matrixWidth      = 20
	matrixHeight     = 20
	statsBoardHeight = 2
)

func (u *ui) render(s gotetromino.State) {
	// if there were cleared lines from matrix in previous state, animate clearing the lines before rendering matrix of new state
	if len(s.ClearedLinesRows) > 0 {
		u.animateClearingLines(s)
	}
	u.screen.Clear()
	u.renderMatrix(s)
	u.renderTetromino(s)
	u.renderStats(s)
}

func (u *ui) renderMatrix(s gotetromino.State) {
	screenWidth, screenHeight := u.screen.Size()
	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	for row := 0; row < len(s.Matrix); row++ {
		for col := 0; col < len(s.Matrix[row]); col++ {
			st := tcell.StyleDefault
			// TODO: Put Block type all in gotetromino.go instead
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			u.screen.SetContent(matrixX+col, matrixY+row, charForBlock(engine.Block(s.Matrix[row][col])), nil, st)
		}
	}
	u.screen.Show()
}

func (u *ui) renderTetromino(s gotetromino.State) {
	screenWidth, screenHeight := u.screen.Size()
	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	x := s.CurrentTetrominoPos[1]
	y := s.CurrentTetrominoPos[0]
	for row := 0; row < len(s.CurrentTetromino); row++ {
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			// only override rendering what is a space block
			matrixRow := y + row
			matrixCol := x + col
			if matrixRow > len(s.Matrix)-1 || matrixCol > len(s.Matrix[0])-1 {
				continue
			}
			if s.Matrix[matrixRow][matrixCol] == int(engine.Space) {
				st := tcell.StyleDefault
				st = st.Foreground(colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
				u.screen.SetContent(matrixX+matrixCol, matrixY+matrixRow, charForBlock(engine.Block(s.CurrentTetromino[row][col])), nil, st)
			}

		}
	}
	u.screen.Show()
}

func (u *ui) animateClearingLines(s gotetromino.State) {
	screenWidth, screenHeight := u.screen.Size()
	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	for row := s.ClearedLinesRows[0]; row <= s.ClearedLinesRows[len(s.ClearedLinesRows)-1]; row++ {
		for col := 1; col < len(s.Matrix[row])-1; col++ {
			st := tcell.StyleDefault
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			u.screen.SetContent(matrixX+col, matrixY+row, '+', nil, st)
		}
	}
	u.screen.Show()
	time.Sleep(200 * time.Millisecond)
}

func (u *ui) renderStats(s gotetromino.State) {
	screenWidth, screenHeight := u.screen.Size()
	matrixX := (screenWidth - matrixWidth) / 2
	st := tcell.StyleDefault
	levelStr := fmt.Sprintf("Level: %d", s.Level)
	scoreStr := fmt.Sprintf("Score: %d", s.Score)
	stats := fmt.Sprintf("%s  %s", levelStr, scoreStr)
	statsBoardX := matrixX + ((matrixWidth - len(stats)) / 2)
	statsBoardY := (screenHeight - matrixHeight - statsBoardHeight) / 2
	for i, r := range stats {
		u.screen.SetContent(statsBoardX+i, statsBoardY, r, nil, st)
	}
	u.screen.Show()
}
