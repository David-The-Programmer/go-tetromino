package renderer

import (
	"log"
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type renderer struct {
	screen tcell.Screen
}

func New() *renderer {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Panic(err)
	}
	if err = screen.Init(); err != nil {
		log.Panic(err)
	}
	screen.HideCursor()
	screen.DisableMouse()
	return &renderer{
		screen: screen,
	}
}

func (r *renderer) Render(s gotetromino.State) {
	defer r.screen.Fini()
	for {
		r.screen.Clear()
        // render matrix
		for row := 0; row < len(s.Matrix); row++ {
			for col := 0; col < len(s.Matrix[row]); col++ {
				st := tcell.StyleDefault
                st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
				r.screen.SetContent(col, row, charForBlock(engine.Block(s.Matrix[row][col])), nil, st)
			}
		}
        // render current tetromino
        x := s.CurrentTetrominoPos[1]
        y := s.CurrentTetrominoPos[0]
        for row := 0; row < len(s.CurrentTetromino); row++ {
            for col := 0; col < len(s.CurrentTetromino[row]); col++ {
				st := tcell.StyleDefault
                st = st.Foreground(colourForBlock(engine.Block(s.CurrentTetromino[row][col])))
				r.screen.SetContent(x+col, y+row, charForBlock(engine.Block(s.CurrentTetromino[row][col])), nil, st)

            }
        }

		r.screen.Show()
        time.Sleep(time.Second)
        break
	}
}
