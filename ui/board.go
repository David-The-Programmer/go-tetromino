package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type board struct {
	screen       tcell.Screen
	x, y, w, h   int
	title        *textBox
	titleColour  tcell.Color
	border       *border
	borderColour tcell.Color
	content      *grid
}

func newBoard(sc tcell.Screen) *board {
	return &board{
		screen:  sc,
		title:   newTextBox(sc),
		border:  newBorder(sc),
		content: newGrid(),
	}
}

func (b *board) SetPos(x, y int) {
	b.x = x
	b.y = y
}

func (b *board) SetDimensions(w, h int) {
	b.w = w
	b.h = h
}

func (b *board) SetTitle(title string) {
	b.title.SetText(title)
}

func (b *board) SetTitleColour(c tcell.Color) {
	b.title.SetColour(c)
}

func (b *board) SetBorderColour(c tcell.Color) {
	b.border.SetColour(c)
}

func (b *board) SetContentRowHeights(rowHeights ...int) {
	b.content.SetGridRowHeights(rowHeights...)
}

func (b *board) SetContentColWidths(colWidths ...int) {
	b.content.SetGridColWidths(colWidths...)
}

func (b *board) AddContent(component gotetromino.UI, gridRow, gridCol, rowSpan, colSpan int) {
	b.content.AddComponent(component, gridRow, gridCol, rowSpan, colSpan)
}

func (b *board) GetContent(gridRow, gridCol int) gotetromino.UI {
	return b.content.GetComponent(gridRow, gridCol)
}

func (b *board) Render() {
	b.border.SetPos(b.x, b.y)
	b.border.SetDimensions(b.w, b.h)
	b.border.Render()

	b.title.SetPos((b.w-len(b.title.GetText()))/2+b.x, b.y)
	b.title.SetDimensions(len(b.title.GetText()), 1)
	b.title.Render()

	b.content.SetPos(b.x+1, b.y+1)
	b.content.SetDimensions(b.w-2, b.h-2)
	b.content.Render()

}
