package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

// const (
// 	matrixWidth       = 20
// 	matrixHeight      = 20
// 	matrixBoardWidth  = matrixWidth + 2*borderThickness
// 	matrixBoardHeight = matrixHeight + 2*borderThickness
// 	statsBoardWidth   = matrixBoardWidth
// 	statsBoardHeight  = 2 + 2*borderThickness
// )

func (u *ui) render(s gotetromino.State) {
	u.screen.Clear()

	g := newGrid()
	g.SetPos(0, 0)
	g.SetDimensions(70, 22)
	g.SetGridRowHeights(1, 1)
	g.SetGridColWidths(19, 22, 29)

	m := newMatrix(u.screen)
	m.SetState(s)

	sb := newStatsBoard(u.screen)
	sb.SetStats(s)

	cb := newControlsBoard(u.screen)

	g.AddComponent(sb, 0, 0, 1, 1)
	g.AddComponent(m, 0, 1, 2, 1)
	g.AddComponent(cb, 1, 2, 1, 1)
	g.Render()

	if s.ClearedPrevLevel {
		// make animation faster as more levels are cleared
		// TODO: Refactor limiting animation speed
		if u.tickDuration > 80*time.Millisecond {
			u.tickDuration = u.tickDuration - (80 * time.Millisecond)
			u.ticker.Reset(u.tickDuration)
		}
	}
}
