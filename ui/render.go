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
	u.renderer.SetState(s)
	u.renderer.Render()

	if s.ClearedPrevLevel {
		// make animation faster as more levels are cleared
		// TODO: Refactor limiting animation speed
		if u.tickDuration > 80*time.Millisecond {
			u.tickDuration = u.tickDuration - (80 * time.Millisecond)
			u.ticker.Reset(u.tickDuration)
		}
	}
}
