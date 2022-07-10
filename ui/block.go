package ui

import (
	"github.com/gdamore/tcell/v2"
)

type block struct {
	screen     tcell.Screen
	x, y, w, h int
	c          tcell.Color
	hide       bool
}

func newBlock(sc tcell.Screen) *block {
	return &block{
		screen: sc,
	}
}

func (b *block) SetPos(x, y int) {
	b.x = x
	b.y = y
}

func (b *block) SetDimensions(w, h int) {
	b.w = w
	b.h = h
}

func (b *block) SetColour(c tcell.Color) {
	b.c = c
}

func (b *block) Hide(v bool) {
	b.hide = v
}

func (b *block) Render() {
	cellX := b.x
	cellY := b.y
	st := tcell.StyleDefault.Foreground(b.c)
	for j := 0; j < b.h; j++ {
		for i := 0; i < b.w; i++ {
			char := ' '
			if !b.hide {
				char = tcell.RuneBlock
				st = st.Background(b.c)
			}
			b.screen.SetContent(cellX, cellY, char, nil, st)
			cellX += 1
		}
		cellX = b.x
		cellY += 1
	}
	b.screen.Show()
}
