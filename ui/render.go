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

	// left boundary
	leftBoundaryX := matrixX - 1
	for i := matrixY; i <= matrixY+matrixHeight; i++ {
		leftBoundaryY := i
		st := tcell.StyleDefault
		st = st.Foreground(tcell.ColorGray)
		u.renderBlock(leftBoundaryX, leftBoundaryY, 1, 1, tcell.RuneVLine, st)
	}

	// right boundary
	rightBoundaryX := matrixX + matrixWidth
	for i := matrixY; i <= matrixY+matrixHeight; i++ {
		rightBoundaryY := i
		st := tcell.StyleDefault
		st = st.Foreground(tcell.ColorGray)
		u.renderBlock(rightBoundaryX, rightBoundaryY, 1, 1, tcell.RuneVLine, st)
	}

	// bottom boundary
	bottomBoundaryY := matrixY + matrixHeight
	for i := matrixX - 1; i <= matrixX+matrixWidth; i++ {
		ch := tcell.RuneHLine
		if i == matrixX-1 {
			ch = tcell.RuneLLCorner
		}
		if i == matrixX+matrixWidth {
			ch = tcell.RuneLRCorner
		}
		bottomBoundaryX := i
		st := tcell.StyleDefault
		st = st.Foreground(tcell.ColorGray)
		u.renderBlock(bottomBoundaryX, bottomBoundaryY, 1, 1, ch, st)

	}

	blockX := matrixX
	blockY := matrixY
	blockWidth := matrixWidth / numCols
	blockHeight := matrixHeight / numRows

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			st := tcell.StyleDefault
			// TODO: Put Block type all in gotetromino.go instead
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			// make sure all tetrominos do not have unwanted lines
			if engine.Block(s.Matrix[row][col]) != engine.Space {
				st = st.Background(colourForBlock(engine.Block(s.Matrix[row][col])))
			}
			u.renderBlock(blockX, blockY, blockWidth, blockHeight, charForBlock(engine.Block(s.Matrix[row][col])), st)
			blockX += blockWidth
		}
		blockX = matrixX
		blockY += blockHeight
	}
	u.screen.Show()
}

func (u *ui) renderBlock(x, y, w, h int, ch rune, st tcell.Style) {
	cellX := x
	cellY := y
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			u.screen.SetContent(cellX, cellY, ch, nil, st)
			cellX += 1
		}
		cellY += 1
	}
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
			if s.Matrix[matrixRow][matrixCol] == int(engine.Space) {
				blockX = matrixX + (matrixCol * blockWidth)
				blockY = matrixY + (matrixRow * blockHeight)
				st := tcell.StyleDefault
				st = st.Foreground(colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
				st = st.Background(colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
				u.renderBlock(blockX, blockY, blockWidth, blockHeight, charForBlock(engine.Block(s.CurrentTetromino[row][col])), st)
			}

		}
	}
	u.screen.Show()
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
