package renderer

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type renderer struct {
	screen tcell.Screen
}

func New(s tcell.Screen) gotetromino.Renderer {
	s.HideCursor()
	s.DisableMouse()
	return &renderer{
		screen: s,
	}
}

func (r *renderer) Render(s gotetromino.State) {
	// animate clearing rows with lines
	if len(s.ClearedLinesRows) > 0 {
		for row := s.ClearedLinesRows[0]; row <= s.ClearedLinesRows[len(s.ClearedLinesRows)-1]; row++ {
			for col := 1; col < len(s.Matrix[row])-1; col++ {
				st := tcell.StyleDefault
				st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
				r.screen.SetContent(col, row, '+', nil, st)
			}
		}
		r.screen.Show()
		time.Sleep(200 * time.Millisecond)
	}

	// render matrix
	for row := 0; row < len(s.Matrix); row++ {
		for col := 0; col < len(s.Matrix[row]); col++ {
			st := tcell.StyleDefault
			// TODO: Put Block type all in gotetromino.go instead
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			r.screen.SetContent(col, row, charForBlock(engine.Block(s.Matrix[row][col])), nil, st)
		}
	}
	// render current tetromino
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
				r.screen.SetContent(x+col, y+row, charForBlock(engine.Block(s.CurrentTetromino[row][col])), nil, st)
			}

		}
	}
	r.screen.Show()
}

func (r *renderer) Stop() {
	r.screen.Fini()
}
