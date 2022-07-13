package ui

import (
	"github.com/gdamore/tcell/v2"
)

type border struct {
	screen     tcell.Screen
	x, y, w, h int
	c          tcell.Color
}

func newBorder(sc tcell.Screen) *border {
	return &border{
		screen: sc,
	}
}

func (b *border) SetPos(x, y int) {
	b.x = x
	b.y = y
}

func (b *border) SetDimensions(w, h int) {
	b.w = w
	b.h = h
}

func (b *border) SetColour(c tcell.Color) {
	b.c = c
}

func (b *border) Render() {
	const borderThickness = 1
	st := tcell.StyleDefault.Foreground(b.c)
	// left & right border
	for i := b.y + borderThickness; i <= b.y+b.h-(2*borderThickness); i++ {
		b.screen.SetContent(b.x, i, tcell.RuneVLine, nil, st)
		b.screen.SetContent(b.x+b.w-borderThickness, i, tcell.RuneVLine, nil, st)
	}

	// top & bottom border
	for i := b.x + borderThickness; i <= b.x+b.w-(2*borderThickness); i++ {
		b.screen.SetContent(i, b.y, tcell.RuneHLine, nil, st)
		b.screen.SetContent(i, b.y+b.h-borderThickness, tcell.RuneHLine, nil, st)
	}

	// corners
	b.screen.SetContent(b.x, b.y, tcell.RuneULCorner, nil, st)
	b.screen.SetContent(b.x+b.w-borderThickness, b.y, tcell.RuneURCorner, nil, st)
	b.screen.SetContent(b.x, b.y+b.h-borderThickness, tcell.RuneLLCorner, nil, st)
	b.screen.SetContent(b.x+b.w-borderThickness, b.y+b.h-borderThickness, tcell.RuneLRCorner, nil, st)

	b.screen.Show()
}
