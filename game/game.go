package game

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/ui"
)

type game struct {
	engine gotetromino.Engine
	ui     gotetromino.UI
	ticker *time.Ticker
}

func New() gotetromino.Game {
	e := engine.New(20, 20)
	ui := ui.New()
	t := time.NewTicker(200 * time.Millisecond)
	g := game{
		engine: e,
		ui:     ui,
		ticker: t,
	}
	var tetrisGame gotetromino.Game = &g
	return tetrisGame
}

func (g *game) Run() {
	actions := g.ui.Action()
	interactions := g.ui.Interaction()
	for {
		g.ui.Render(g.engine.State())
		select {
		case i := <-interactions:
			if i == gotetromino.Exit {
				g.ui.Stop()
				return
			}
			if i == gotetromino.Restart && g.engine.State().Over {
				g.engine.Reset()
			}
		case a := <-actions:
			g.engine.Step(a)
		case <-g.ticker.C:
			g.engine.Step(gotetromino.SoftDrop)
		}

	}
}
