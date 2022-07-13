package ui

import (
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
		"Quit",
		"Restart",
		"Left",
		"Right",
		"Soft drop",
		"Hard drop",
		"Rotate left",
		"Rotate right",
	}
	keys := []string{
		"esc",
		"r",
		string(tcell.RuneLArrow),
		string(tcell.RuneRArrow),
		string(tcell.RuneDArrow),
		"spacebar",
		"z",
		"x",
	}
	cb.SetContentColWidths(1, 1)
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
	for i := range keys {
		tb := newTextBox(sc)
		tb.SetText(keys[i])
		tb.AlignText(right)
		tb.SetColour(tcell.ColorGrey)
		cb.AddContent(tb, i, 1, 1, 1)
	}
	return &controlsBoard{
		board: cb,
	}
}
