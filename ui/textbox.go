package ui

import (
	"github.com/gdamore/tcell/v2"
)

type textAlignment int

const (
	left textAlignment = iota
	right
	center
)

type textBox struct {
	screen     tcell.Screen
	x, y, w, h int
	c          tcell.Color
	text       string
	ta         textAlignment
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

func (t *textBox) AlignText(ta textAlignment) {
	t.ta = ta
}

func (t *textBox) GetText() string {
	return t.text
}

func (t *textBox) SetColour(c tcell.Color) {
	t.c = c
}

func (t *textBox) Render() {
	st := tcell.StyleDefault.Foreground(t.c)
	// clear all previous text before rendering new text
	for row := 0; row < t.h; row++ {
		for col := 0; col < t.w; col++ {
			t.screen.SetContent(t.x+col, t.y+row, ' ', nil, st)
		}
	}

	text := []rune(t.text)
	for row := 0; row < t.h; row++ {
		cellX := t.x
		if t.ta == right {
			xOffset := t.w - (len(text) - 1)
			if xOffset > 0 {
				cellX += xOffset
			}
		}
		if t.ta == center {
			xOffset := (t.w - (len(text) - 1)) / 2
			if xOffset > 0 {
				cellX += xOffset
			}

		}
		cellY := t.y + row
		for col := 0; col < t.w; col++ {
			i := row*t.w + col
			if i > len(text)-1 {
				continue
			}
			t.screen.SetContent(cellX, cellY, text[i], nil, st)
			cellX += 1
		}
	}
	t.screen.Show()
}
