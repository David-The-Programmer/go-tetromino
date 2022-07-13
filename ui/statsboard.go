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

	scoreTitle := newTextBox(sc)
	scoreTitle.SetColour(tcell.ColorGrey)
	scoreTitle.SetText("Score")
	score := newTextBox(sc)
	score.SetColour(tcell.ColorGrey)
	score.SetText("0")
	score.AlignText(right)

	levelTitle := newTextBox(sc)
	levelTitle.SetColour(tcell.ColorGrey)
	levelTitle.SetText("Level")
	level := newTextBox(sc)
	level.SetColour(tcell.ColorGrey)
	level.SetText("0")
	level.AlignText(right)

	linesTitle := newTextBox(sc)
	linesTitle.SetColour(tcell.ColorGrey)
	linesTitle.SetText("Lines")
	lines := newTextBox(sc)
	lines.SetColour(tcell.ColorGrey)
	lines.SetText("0")
	lines.AlignText(right)

	sb.SetContentColWidths(1, 1)
	sb.SetContentRowHeights(1, 1, 1)
	sb.AddContent(scoreTitle, 0, 0, 1, 1)
	sb.AddContent(score, 0, 1, 1, 1)
	sb.AddContent(levelTitle, 1, 0, 1, 1)
	sb.AddContent(level, 1, 1, 1, 1)
	sb.AddContent(linesTitle, 2, 0, 1, 1)
	sb.AddContent(lines, 2, 1, 1, 1)

	return &statsBoard{
		board: sb,
	}
}

func (sb *statsBoard) SetStats(s gotetromino.State) {
	score := sb.GetContent(0, 1).(*gridComponent).UI.(*textBox)
	score.SetText(fmt.Sprintf("%d", s.Score))

	level := sb.GetContent(1, 1).(*gridComponent).UI.(*textBox)
	level.SetText(fmt.Sprintf("%d", s.Level))

	lines := sb.GetContent(2, 1).(*gridComponent).UI.(*textBox)
	lines.SetText(fmt.Sprintf("%d", s.LineCount))
}
