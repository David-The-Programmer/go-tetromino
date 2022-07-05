package ui

import (
	"fmt"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type statsBoard struct {
	*board
}

func newStatsBoard(sc tcell.Screen) *statsBoard {
	sb := newBoard(sc)
	sb.SetBorderColour(tcell.ColorGrey)
	sb.SetTitle("Stats")
	sb.SetTitleColour(tcell.ColorGrey)

	score := newTextBox(sc)
	score.SetColour(tcell.ColorGrey)

	level := newTextBox(sc)
	level.SetColour(tcell.ColorGrey)

	lines := newTextBox(sc)
	lines.SetColour(tcell.ColorGrey)

	sb.SetContentColWidths(1)
	sb.SetContentRowHeights(1, 1, 1)
	sb.AddContent(score, 0, 0, 1, 1)
	sb.AddContent(level, 1, 0, 1, 1)
	sb.AddContent(lines, 2, 0, 1, 1)

	return &statsBoard{
		board: sb,
	}
}

func (sb *statsBoard) SetStats(s gotetromino.State) {
	score := sb.GetContent(0, 0).(*gridComponent).UI.(*textBox)
	score.SetText(fmt.Sprintf("SCORE: %d", s.Score))

	level := sb.GetContent(1, 0).(*gridComponent).UI.(*textBox)
	level.SetText(fmt.Sprintf("LEVEL: %d", s.Level))

	lines := sb.GetContent(2, 0).(*gridComponent).UI.(*textBox)
	lines.SetText(fmt.Sprintf("LINES: %d", s.LineCount))
}
