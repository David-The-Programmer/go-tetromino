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
	u.screen.Clear()
	u.renderMatrix(s)
	u.renderGhostTetromino(s)
	u.renderTetromino(s)
	u.renderStats(s)
	if s.ClearedPrevLevel {
		// make animation faster as more levels are cleared
		// TODO: Refactor limiting animation speed
		if u.tickDuration > 80*time.Millisecond {
			u.tickDuration = u.tickDuration - (80 * time.Millisecond)
			u.ticker.Reset(u.tickDuration)
		}
	}
}

func (u *ui) renderMatrix(s gotetromino.State) {
	numRows := len(s.Matrix)
	numCols := len(s.Matrix[0])

	screenWidth, screenHeight := u.screen.Size()

	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	u.renderBorder(matrixX-1, matrixY-1, matrixWidth+1, matrixHeight+1, tcell.ColorGrey)

	blockX := matrixX
	blockY := matrixY
	blockWidth := matrixWidth / numCols
	blockHeight := matrixHeight / numRows

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			// TODO: Put Block type all in gotetromino.go instead
			if engine.Block(s.Matrix[row][col]) != engine.Space {
				u.renderBlock(blockX, blockY, blockWidth, blockHeight, colourForBlock(engine.Block(s.Matrix[row][col])))
			}
			blockX += blockWidth
		}
		blockX = matrixX
		blockY += blockHeight
	}
	u.screen.Show()
}

func (u *ui) renderTetromino(s gotetromino.State) {
	numRows := len(s.Matrix)
	numCols := len(s.Matrix[0])

	screenWidth, screenHeight := u.screen.Size()

	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	blockX := 0
	blockY := 0
	blockWidth := matrixWidth / numCols
	blockHeight := matrixHeight / numRows

	x := s.CurrentTetrominoPos[1]
	y := s.CurrentTetrominoPos[0]
	for row := 0; row < len(s.CurrentTetromino); row++ {
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			// only override rendering what is a space block
			matrixRow := y + row
			matrixCol := x + col
			if matrixRow < 0 || matrixCol < 0 || matrixRow > len(s.Matrix)-1 || matrixCol > len(s.Matrix[0])-1 {
				continue
			}
			if s.Matrix[matrixRow][matrixCol] == int(engine.Space) && engine.Block(s.CurrentTetromino[row][col]) != engine.Space {
				blockX = matrixX + (matrixCol * blockWidth)
				blockY = matrixY + (matrixRow * blockHeight)
				u.renderBlock(blockX, blockY, blockWidth, blockHeight, colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
			}

		}
	}
	u.screen.Show()
}

func (u *ui) renderGhostTetromino(s gotetromino.State) {
	numRows := len(s.Matrix)
	numCols := len(s.Matrix[0])

	screenWidth, screenHeight := u.screen.Size()

	matrixX := (screenWidth - matrixWidth) / 2
	matrixY := ((screenHeight - matrixHeight - statsBoardHeight) / 2) + statsBoardHeight

	blockX := 0
	blockY := 0
	blockWidth := matrixWidth / numCols
	blockHeight := matrixHeight / numRows

	x := s.GhostTetrominoPos[1]
	y := s.GhostTetrominoPos[0]
	for row := 0; row < len(s.CurrentTetromino); row++ {
		for col := 0; col < len(s.CurrentTetromino[row]); col++ {
			// only override rendering what is a space block
			matrixRow := y + row
			matrixCol := x + col
			if matrixRow < 0 || matrixCol < 0 || matrixRow > len(s.Matrix)-1 || matrixCol > len(s.Matrix[0])-1 {
				continue
			}
			if s.Matrix[matrixRow][matrixCol] == int(engine.Space) && engine.Block(s.CurrentTetromino[row][col]) != engine.Space {
				blockX = matrixX + (matrixCol * blockWidth)
				blockY = matrixY + (matrixRow * blockHeight)
				u.renderBlock(blockX, blockY, blockWidth, blockHeight, tcell.ColorGray)
			}

		}
	}
	u.screen.Show()
}

func (u *ui) renderStats(s gotetromino.State) {
	screenWidth, screenHeight := u.screen.Size()
	matrixX := (screenWidth - matrixWidth) / 2
	st := tcell.StyleDefault
	levelStr := fmt.Sprintf("LEVEL: %d", s.Level)
	scoreStr := fmt.Sprintf("SCORE: %d", s.Score)
	stats := fmt.Sprintf("%s  %s", levelStr, scoreStr)
	statsBoardX := matrixX + ((matrixWidth - len(stats)) / 2)
	statsBoardY := (screenHeight - matrixHeight - statsBoardHeight) / 2
	for i, r := range stats {
		u.screen.SetContent(statsBoardX+i, statsBoardY, r, nil, st)
	}
	u.screen.Show()
}

func (u *ui) renderBlock(x, y, w, h int, c tcell.Color) {
	cellX := x
	cellY := y
	st := tcell.StyleDefault.Foreground(c).Background(c)
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			u.screen.SetContent(cellX, cellY, tcell.RuneBlock, nil, st)
			cellX += 1
		}
		cellY += 1
	}
}

func (u *ui) renderBorder(x, y, w, h int, c tcell.Color) {
	st := tcell.StyleDefault.Foreground(c)
	// left & right border
	for i := y + 1; i <= y+h-1; i++ {
		u.screen.SetContent(x, i, tcell.RuneVLine, nil, st)
		u.screen.SetContent(x+w, i, tcell.RuneVLine, nil, st)
	}

	// top & bottom border
	for i := x + 1; i <= x+w-1; i++ {
		u.screen.SetContent(i, y, tcell.RuneHLine, nil, st)
		u.screen.SetContent(i, y+h, tcell.RuneHLine, nil, st)
	}

	// corners
	u.screen.SetContent(x, y, tcell.RuneULCorner, nil, st)
	u.screen.SetContent(x+w, y, tcell.RuneURCorner, nil, st)
	u.screen.SetContent(x, y+h, tcell.RuneLLCorner, nil, st)
	u.screen.SetContent(x+w, y+h, tcell.RuneLRCorner, nil, st)

}
