package ui

import (
	// "fmt"

	"fmt"
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
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

	sb := newBoard(u.screen)
	sb.SetBorderColour(tcell.ColorGrey)
	sb.SetTitle("STATS")
	sb.SetTitleColour(tcell.ColorGrey)
	score := newTextBox(u.screen)
	score.SetText(fmt.Sprintf("SCORE: %d", s.Score))
	score.SetColour(tcell.ColorGrey)

	level := newTextBox(u.screen)
	level.SetText(fmt.Sprintf("LEVEL: %d", s.Level))
	level.SetColour(tcell.ColorGrey)

	lines := newTextBox(u.screen)
	lines.SetText(fmt.Sprintf("LINES: %d", s.LineCount))
	lines.SetColour(tcell.ColorGrey)

	sb.SetContentColWidths(1, 2, 1)
	sb.SetContentRowHeights(1, 1, 1, 1)
	sb.AddContent(score, 1, 1, 1, 1)
	sb.AddContent(level, 2, 1, 1, 1)
	sb.AddContent(lines, 3, 1, 1, 1)

	cb := newBoard(u.screen)
	cb.SetBorderColour(tcell.ColorGrey)
	cb.SetTitle("CONTROLS")
	cb.SetTitleColour(tcell.ColorGrey)
	keymaps := []string{
		"esc - quit",
		fmt.Sprintf("%c - left", tcell.RuneLArrow),
		fmt.Sprintf("%c - right", tcell.RuneRArrow),
		fmt.Sprintf("%c - soft drop", tcell.RuneDArrow),
		"space - hard drop",
		"x - rotate clockwise",
		"z - rotate anti-clockwise",
	}
	cb.SetContentColWidths(1, 25, 1)
	cb.SetContentRowHeights(1, 1, 1, 1, 1, 1, 1, 1)
	for i := range keymaps {
		km := newTextBox(u.screen)
		km.SetText(keymaps[i])
		km.SetColour(tcell.ColorGrey)
		cb.AddContent(km, i+1, 1, 1, 1)
	}

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
