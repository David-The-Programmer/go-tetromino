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
	scWidth, scHeight := sc.Size()
	rWidth, rHeight := 61, 22
	g.SetPos((scWidth-rWidth)/2, (scHeight-rHeight)/2)
	g.SetDimensions(rWidth, rHeight)
	g.SetGridColWidths(22, 10, 29)
	g.SetGridRowHeights(6, 6, 10)

	m := newMatrix(sc, s)
	m.SetState(s)
	sb := newStatsBoard(sc)
	sb.SetStats(s)
	cb := newControlsBoard(sc)
	nb := newNextBoard(sc, s)
	nb.SetNextTetromino(s)
	gsb := newStatusBoard(sc)
	gsb.SetStatus(s)

	g.AddComponent(m, 0, 0, 3, 1)
	g.AddComponent(nb, 0, 1, 1, 1)
	g.AddComponent(gsb, 0, 2, 1, 1)
	g.AddComponent(sb, 1, 1, 1, 2)
	g.AddComponent(cb, 2, 1, 1, 2)
	return &renderer{
		grid: g,
	}
}

func (r *renderer) SetState(s gotetromino.State) {
	m := r.GetComponent(0, 0).(*gridComponent).UI.(*matrix)
	m.SetState(s)
	nb := r.GetComponent(0, 1).(*gridComponent).UI.(*nextBoard)
	nb.SetNextTetromino(s)
	sb := r.GetComponent(1, 1).(*gridComponent).UI.(*statsBoard)
	sb.SetStats(s)
	gsb := r.GetComponent(0, 2).(*gridComponent).UI.(*statusBoard)
	gsb.SetStatus(s)
}
