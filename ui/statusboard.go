package ui

import (
	"fmt"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type statusBoard struct {
	*board
}

func newStatusBoard(sc tcell.Screen) *statusBoard {
	sb := newBoard(sc)
	sb.SetBorderColour(tcell.ColorGrey)
	sb.SetTitle("Status")
	sb.SetTitleColour(tcell.ColorGrey)

	sb.SetContentColWidths(1)
	sb.SetContentRowHeights(1, 1, 1)
	numRows := len(sb.GetContentRowHeights())
	for i := 0; i < numRows; i++ {
		text := ""
		if i == numRows-1 {
			text = "Game Started..."
		}
		tb := newTextBox(sc)
		tb.SetText(text)
		tb.AlignText(center)
		tb.SetColour(tcell.ColorGrey)
		sb.AddContent(tb, i, 0, 1, 1)

	}
	return &statusBoard{
		board: sb,
	}
}

func (sb *statusBoard) SetStatus(s gotetromino.State) {
	if s.Over {
		sb.sendStatusMsg("Game Over...")
		return
	}
	// Let player know that game has started if game was restarted after game over
	if s.Reset {
		sb.sendStatusMsg("Game Started...")
	}
	numLinesCleared := len(s.ClearedLinesRows)
	if numLinesCleared > 0 && numLinesCleared < 4 {
		sb.sendStatusMsg(fmt.Sprintf("%d Line(s) Cleared!", numLinesCleared))
	}
	if numLinesCleared == 4 {
		sb.sendStatusMsg("TETRIS!")
	}
	if s.ClearedPrevLevel {
		sb.sendStatusMsg("Level Up!")
	}
}

func (sb *statusBoard) latestStatusMsg() string {
	return sb.GetContent(len(sb.GetContentRowHeights())-1, 0).(*gridComponent).UI.(*textBox).GetText()
}

func (sb *statusBoard) sendStatusMsg(msg string) {
	// bubble up old status messages to make way for the new status message, msg
	numRows := len(sb.GetContentRowHeights())
	for row := 0; row < numRows; row++ {
		if row >= numRows-1 {
			break
		}
		sb.GetContent(row, 0).(*gridComponent).UI.(*textBox).SetText(sb.GetContent(row+1, 0).(*gridComponent).UI.(*textBox).GetText())
	}
	sb.GetContent(numRows-1, 0).(*gridComponent).UI.(*textBox).SetText(msg)
}
