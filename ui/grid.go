package ui

import (
	"log"

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
	if g.w == 0 {
		log.Panicln("Width of grid must be set before render")
	}
	if g.h == 0 {
		log.Panicln("Height of grid must be set before render")
	}
	if len(g.colWidths) == 0 {
		log.Panicln("Column widths must be set before render")
	}
	if len(g.rowHeights) == 0 {
		log.Panicln("Row heights must be set before render")
	}

	// calculate absolute widths from given fractional unit of column widths
	sum := 0
	for _, w := range g.colWidths {
		if w <= 0 {
			log.Panicln("Column width cannot be <= 0")
		}
		sum += w
	}
	colAbsWidth := g.w / sum
	// calculate absolute heights from given fractional unit of row heights
	sum = 0
	for _, h := range g.rowHeights {
		if h <= 0 {
			log.Panicln("Row heights cannot be <= 0")
		}
		sum += h
	}
	rowAbsHeight := g.h / sum

	for i := range g.components {
		c := g.components[i]
		xOffset := 0
		for j := 0; j < c.gridCol; j++ {
			xOffset += g.colWidths[j] * colAbsWidth
		}
		yOffset := 0
		for j := 0; j < c.gridRow; j++ {
			yOffset += g.rowHeights[j] * rowAbsHeight
		}
		c.SetPos(g.x+xOffset, g.y+yOffset)

		width := 0
		for j := c.gridCol; j < c.gridCol+c.colSpan; j++ {
			width += g.colWidths[j] * colAbsWidth
		}
		height := 0
		for j := c.gridRow; j < c.gridRow+c.rowSpan; j++ {
			height += g.rowHeights[j] * rowAbsHeight
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

func (g *grid) GetComponent(gridRow, gridCol int) gotetromino.UI {
	for i := range g.components {
		c := g.components[i]
		if c.gridRow == gridRow && c.gridCol == gridCol {
			return c
		}
	}
	return nil
}
