package ui

import (
	// "fmt"
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
	g.SetPos(4, 3)
	g.SetDimensions(u.screen.Size())
	g.SetGridColWidths(22)
	g.SetGridRowHeights(22)
	m := newMatrix(u.screen)
	m.SetState(s)
	g.AddComponent(m, 0, 0, 1, 1)
	g.Render()
	// u.renderStats(s)
	if s.ClearedPrevLevel {
		// make animation faster as more levels are cleared
		// TODO: Refactor limiting animation speed
		if u.tickDuration > 80*time.Millisecond {
			u.tickDuration = u.tickDuration - (80 * time.Millisecond)
			u.ticker.Reset(u.tickDuration)
		}
	}
}

// func (u *ui) renderStats(s gotetromino.State) {
// 	screenWidth, screenHeight := u.screen.Size()
// 	level := fmt.Sprintf("LEVEL: %d", s.Level)
// 	score := fmt.Sprintf("SCORE: %d", s.Score)
// 	statsBoardX := (screenWidth - statsBoardWidth) / 2
// 	statsBoardY := (screenHeight - matrixHeight - statsBoardHeight) / 2
// 	u.renderBoard(statsBoardX, statsBoardY, statsBoardWidth, statsBoardHeight, "STATS", tcell.ColorGray, tcell.ColorGray)
// 	u.renderText(statsBoardX+(statsBoardWidth-len(level))/2, statsBoardY+1, level, tcell.ColorGray)
// 	u.renderText(statsBoardX+(statsBoardWidth-len(score))/2, statsBoardY+2, score, tcell.ColorGray)
// }
//
// func (u *ui) renderBoard(x, y, w, h int, title string, borderColour, titleColour tcell.Color) {
// 	u.renderBorder(x, y, w, h, borderColour)
// 	u.renderText(x+((w-len(title))/2), y, title, titleColour)
// }
//
// func (u *ui) renderText(x, y int, text string, c tcell.Color) {
// 	st := tcell.StyleDefault.Foreground(c)
// 	for i, r := range text {
// 		u.screen.SetContent(x+i, y, r, nil, st)
// 	}
// 	u.screen.Show()
// }
//
// func (u *ui) renderBorder(x, y, w, h int, c tcell.Color) {
// 	st := tcell.StyleDefault.Foreground(c)
// 	// left & right border
// 	for i := y + borderThickness; i <= y+h-(2*borderThickness); i++ {
// 		u.screen.SetContent(x, i, tcell.RuneVLine, nil, st)
// 		u.screen.SetContent(x+w-borderThickness, i, tcell.RuneVLine, nil, st)
// 	}
//
// 	// top & bottom border
// 	for i := x + borderThickness; i <= x+w-(2*borderThickness); i++ {
// 		u.screen.SetContent(i, y, tcell.RuneHLine, nil, st)
// 		u.screen.SetContent(i, y+h-borderThickness, tcell.RuneHLine, nil, st)
// 	}
//
// 	// corners
// 	u.screen.SetContent(x, y, tcell.RuneULCorner, nil, st)
// 	u.screen.SetContent(x+w-borderThickness, y, tcell.RuneURCorner, nil, st)
// 	u.screen.SetContent(x, y+h-borderThickness, tcell.RuneLLCorner, nil, st)
// 	u.screen.SetContent(x+w-borderThickness, y+h-borderThickness, tcell.RuneLRCorner, nil, st)
//
// 	u.screen.Show()
//
// }
