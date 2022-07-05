package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type controlsBoard struct {
	*board
}

func newControlsBoard(sc tcell.Screen) *controlsBoard {
	cb := newBoard(sc)
	cb.SetBorderColour(tcell.ColorGrey)
	cb.SetTitle("Controls")
	cb.SetTitleColour(tcell.ColorGrey)

	controls := []string{
		"esc - quit game",
		"r - restart game",
		fmt.Sprintf("%c - left", tcell.RuneLArrow),
		fmt.Sprintf("%c - right", tcell.RuneRArrow),
		fmt.Sprintf("%c - soft drop", tcell.RuneDArrow),
		"space - hard drop",
		"x - rotate clockwise",
		"z - rotate anti-clockwise",
	}
	cb.SetContentColWidths(1)
	rowHeights := []int{}
	for range controls {
		rowHeights = append(rowHeights, 1)
	}
	cb.SetContentRowHeights(rowHeights...)
	for i := range controls {
		tb := newTextBox(sc)
		tb.SetText(controls[i])
		tb.SetColour(tcell.ColorGrey)
		cb.AddContent(tb, i, 0, 1, 1)
	}
	return &controlsBoard{
		board: cb,
	}
}
