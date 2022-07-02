package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type gridComponent struct {
	gotetromino.UI
	gridRow int
	gridCol int
	rowSpan int
	colSpan int
}

type grid struct {
	x, y, w, h            int
	components            []*gridComponent
	rowHeights, colWidths []int
}

func newGrid() *grid {
	return &grid{}
}

func (g *grid) SetPos(x, y int) {
	g.x = x
	g.y = y
}

func (g *grid) SetDimensions(w, h int) {
	g.w = w
	g.h = h
}

func (g *grid) Render() {
	for i := range g.components {
		c := g.components[i]
		xOffset := 0
		for j := 0; j < c.gridCol; j++ {
			xOffset += g.colWidths[j]
		}
		yOffset := 0
		for j := 0; j < c.gridRow; j++ {
			yOffset += g.rowHeights[j]
		}
		c.SetPos(g.x+xOffset, g.y+yOffset)

		width := 0
		for j := c.gridCol; j < c.gridCol+c.colSpan; j++ {
			width += g.colWidths[j]
		}
		height := 0
		for j := c.gridRow; j < c.gridRow+c.rowSpan; j++ {
			height += g.rowHeights[j]
		}
		c.SetDimensions(width, height)

		c.Render()
	}
}

func (g *grid) SetGridRowHeights(rowHeights ...int) {
	g.rowHeights = rowHeights
}

func (g *grid) SetGridColWidths(colWidths ...int) {
	g.colWidths = colWidths
}

func (g *grid) AddComponent(component gotetromino.UI, gridRow, gridCol, rowSpan, colSpan int) {
	g.components = append(g.components, &gridComponent{
		component,
		gridRow,
		gridCol,
		rowSpan,
		colSpan,
	})
}
