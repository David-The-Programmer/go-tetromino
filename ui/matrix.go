package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type matrix struct {
	x, y, w, h int
	screen     tcell.Screen
	state      gotetromino.State
}

func newMatrix(sc tcell.Screen) *matrix {
	return &matrix{
		screen: sc,
	}
}

func (m *matrix) SetPos(x, y int) {
	m.x = x
	m.y = y
}

func (m *matrix) SetDimensions(w, h int) {
	m.w = w
	m.h = h
}

func (m *matrix) SetState(s gotetromino.State) {
	m.state = s
}

func (m *matrix) Render() {
	// render border around matrix
	mb := newBorder(m.screen)
	mb.SetPos(m.x, m.y)
	mb.SetDimensions(m.w, m.h)
	mb.SetColour(tcell.ColorGrey)
	mb.Render()

	// render all locked blocks in the matrix
	numRows := len(m.state.Matrix)
	numCols := len(m.state.Matrix[0])
	blockX := m.x + 1
	blockY := m.y + 1
	blockWidth := (m.w - 2) / numCols
	blockHeight := (m.h - 2) / numRows
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			// TODO: Put Block type all in gotetromino.go instead
			if engine.Block(m.state.Matrix[row][col]) != engine.Space {
				b := newBlock(m.screen)
				b.SetDimensions(blockWidth, blockHeight)
				b.SetPos(blockX, blockY)
				b.SetColour(colourForBlock(engine.Block(m.state.Matrix[row][col])))
				b.Render()
			}
			blockX += blockWidth
		}
		blockX = m.x + 1
		blockY += blockHeight
	}

	// render current tetromino
	for row := 0; row < len(m.state.CurrentTetromino); row++ {
		for col := 0; col < len(m.state.CurrentTetromino[row]); col++ {
			matrixRow := m.state.CurrentTetrominoPos[0] + row
			matrixCol := m.state.CurrentTetrominoPos[1] + col
			if matrixRow < 0 || matrixRow > len(m.state.Matrix)-1 {
				continue
			}
			if matrixCol < 0 || matrixCol > len(m.state.Matrix[0])-1 {
				continue
			}
			if engine.Block(m.state.CurrentTetromino[row][col]) == engine.Space {
				continue
			}
			tBlockX := m.x + 1 + (matrixCol * blockWidth)
			tBlockY := m.y + 1 + (matrixRow * blockHeight)
			b := newBlock(m.screen)
			b.SetDimensions(blockWidth, blockHeight)
			b.SetPos(tBlockX, tBlockY)
			b.SetColour(colourForBlock(engine.Block(m.state.CurrentTetromino[row][col])))
			b.Render()
		}
	}

	// render ghost tetromino
	for row := 0; row < len(m.state.CurrentTetromino); row++ {
		for col := 0; col < len(m.state.CurrentTetromino[row]); col++ {
			matrixRow := m.state.GhostTetrominoPos[0] + row
			matrixCol := m.state.GhostTetrominoPos[1] + col
			if matrixRow < 0 || matrixRow > len(m.state.Matrix)-1 {
				continue
			}
			if matrixCol < 0 || matrixCol > len(m.state.Matrix[0])-1 {
				continue
			}
			if engine.Block(m.state.CurrentTetromino[row][col]) == engine.Space {
				continue
			}
			tBlockX := m.x + 1 + (matrixCol * blockWidth)
			tBlockY := m.y + 1 + (matrixRow * blockHeight)
			b := newBlock(m.screen)
			b.SetDimensions(blockWidth, blockHeight)
			b.SetPos(tBlockX, tBlockY)
			b.SetColour(tcell.ColorGray)
			b.Render()
		}
	}

}
