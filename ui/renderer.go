package ui

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type renderer struct {
	*grid
}

func newRenderer(sc tcell.Screen, s gotetromino.State) *renderer {
	g := newGrid()
	g.SetPos(0, 0)
	g.SetDimensions(70, 22)
	g.SetGridRowHeights(1, 1)
	g.SetGridColWidths(19, 22, 29)

	m := newMatrix(sc, s)
	m.SetState(s)
	sb := newStatsBoard(sc)
	sb.SetStats(s)
	cb := newControlsBoard(sc)

	g.AddComponent(sb, 0, 0, 1, 1)
	g.AddComponent(m, 0, 1, 2, 1)
	g.AddComponent(cb, 1, 2, 1, 1)
	return &renderer{
		grid: g,
	}
}

func (r *renderer) SetState(s gotetromino.State) {
	sb := r.GetComponent(0, 0).(*gridComponent).UI.(*statsBoard)
	sb.SetStats(s)
	m := r.GetComponent(0, 1).(*gridComponent).UI.(*matrix)
	m.SetState(s)
}
