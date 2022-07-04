package ui

import (
	"github.com/gdamore/tcell/v2"
)

type textBox struct {
	screen     tcell.Screen
	x, y, w, h int
	c          tcell.Color
	text       string
}

func newTextBox(sc tcell.Screen) *textBox {
	return &textBox{
		screen: sc,
	}
}

func (t *textBox) SetPos(x, y int) {
	t.x = x
	t.y = y
}

func (t *textBox) SetDimensions(w, h int) {
	t.w = w
	t.h = h
}

func (t *textBox) SetText(text string) {
	t.text = text
}

func (t *textBox) GetText() string {
	return t.text
}

func (t *textBox) SetColour(c tcell.Color) {
	t.c = c
}

func (t *textBox) Render() {
	cellX := t.x
	cellY := t.y
	st := tcell.StyleDefault.Foreground(t.c)

	// text would be left aligned for now
	// TODO: Make different alignments for text
	text := []rune(t.text)
	for row := 0; row < t.h; row++ {
		for col := 0; col < t.w; col++ {
			i := row*t.w + col
			if i > len(text)-1 {
				continue
			}
			t.screen.SetContent(cellX, cellY, text[i], nil, st)
			cellX += 1
		}
		cellX = t.x
		cellY += 1
	}
	t.screen.Show()
}
