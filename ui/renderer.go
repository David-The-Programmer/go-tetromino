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
	g.SetDimensions(78, 22)
	g.SetGridRowHeights(5, 1, 6, 10)
	g.SetGridColWidths(27, 22, 10, 19)

	m := newMatrix(sc, s)
	m.SetState(s)
	sb := newStatsBoard(sc)
	sb.SetStats(s)
	cb := newControlsBoard(sc)
	nb := newNextBoard(sc, s)
	nb.SetNextTetromino(s)

	g.AddComponent(sb, 0, 0, 1, 1)
	g.AddComponent(m, 0, 1, 4, 1)
	g.AddComponent(nb, 0, 2, 2, 1)
	g.AddComponent(cb, 3, 2, 1, 2)
	return &renderer{
		grid: g,
	}
}

func (r *renderer) SetState(s gotetromino.State) {
	sb := r.GetComponent(0, 0).(*gridComponent).UI.(*statsBoard)
	sb.SetStats(s)
	m := r.GetComponent(0, 1).(*gridComponent).UI.(*matrix)
	m.SetState(s)
	nb := r.GetComponent(0, 2).(*gridComponent).UI.(*nextBoard)
	nb.SetNextTetromino(s)
}
