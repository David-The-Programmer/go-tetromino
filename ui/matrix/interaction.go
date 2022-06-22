package matrix

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

type Matrix struct {
	screen        tcell.Screen
	engineService gotetromino.EngineService
	user          gotetromino.User
	interaction   chan gotetromino.Interaction
	// player gotetromino.Player
	// ticker        *time.Ticker
	// tickDuration  time.Duration
	// position      []int
}

func New(s tcell.Screen, es gotetromino.EngineService, u gotetromino.User) *Matrix {
	s.HideCursor()
	s.DisableMouse()
	// d := 800 * time.Millisecond
	// t := time.NewTicker(d)
	i := make(chan gotetromino.Interaction)
	m := &Matrix{
		screen:        s,
		engineService: es,
		user:          u,
		interaction:   i,
		// ticker:       t,
		// tickDuration: d,
		// position:     append([]int{}, pos...),
	}
	u.Register(m)
	return m
}

// func (m *Matrix) RunAndListen(u gotetromino.User) {
// 	actions := u.Action()
// 	interactions := u.Interaction()
// 	state := m.engine.State()
// 	for {
// 		newState := m.engine.State()
// 		if newState.Level-state.Level > 0 {
// 			m.tickDuration = m.tickDuration - time.Duration(newState.Level*int(m.tickDuration/10))
// 			m.ticker.Reset(m.tickDuration)
// 		}
// 		state = newState
// 		render(m.screen, state)
// 		select {
// 		case i := <-interactions:
// 			if i == gotetromino.Exit {
// 				stop(m.screen)
// 				return
// 			}
// 			if i == gotetromino.Restart && m.engine.State().Over {
// 				m.engine.Reset()
// 			}
// 		case a := <-actions:
// 			m.engine.Step(a)
// 		case <-m.ticker.C:
// 			m.engine.Step(gotetromino.None)
// 		}
//
// 	}
//
// }

func (m *Matrix) Run() {
	for {
		render(m.screen, m.engineService.State())
		select {
		case i := <-m.interaction:
			if i == gotetromino.Exit || i == gotetromino.Restart {
				m.engineService.Stop()
				stop(m.screen)
				return
			}
		}
	}
}

func (m *Matrix) Notify() {
	// get user interaction if notified to get state
	m.interaction <- m.user.Interaction()
}

func render(sc tcell.Screen, s gotetromino.State) {
	// if there were cleared lines from matrix in previous state, animate clearing the lines before rendering matrix of new state
	if len(s.ClearedLinesRows) > 0 {
		animateClearingLines(sc, s)
	}
	renderMatrix(sc, s)
	renderTetromino(sc, s)
}

func renderMatrix(sc tcell.Screen, s gotetromino.State) {
	for row := 0; row < len(s.Matrix); row++ {
		for col := 0; col < len(s.Matrix[row]); col++ {
			st := tcell.StyleDefault
			// TODO: Put Block type all in gotetromino.go instead
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			sc.SetContent(col, row, charForBlock(engine.Block(s.Matrix[row][col])), nil, st)
		}
	}
	sc.Show()
}

func renderTetromino(sc tcell.Screen, s gotetromino.State) {
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
				sc.SetContent(x+col, y+row, charForBlock(engine.Block(s.CurrentTetromino[row][col])), nil, st)
			}

		}
	}
	sc.Show()
}

func animateClearingLines(sc tcell.Screen, s gotetromino.State) {
	for row := s.ClearedLinesRows[0]; row <= s.ClearedLinesRows[len(s.ClearedLinesRows)-1]; row++ {
		for col := 1; col < len(s.Matrix[row])-1; col++ {
			st := tcell.StyleDefault
			st = st.Foreground(colourForBlock(engine.Block(s.Matrix[row][col])))
			sc.SetContent(col, row, '+', nil, st)
		}
	}
	sc.Show()
	time.Sleep(200 * time.Millisecond)
}

func stop(sc tcell.Screen) {
	sc.Fini()
}
