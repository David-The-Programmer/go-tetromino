package game

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/ui"
)

type game struct {
	engine gotetromino.Engine
	ui     gotetromino.UI
}

func New() gotetromino.Game {
	e := engine.New(20, 10)
	ui := ui.New()
	g := game{
		engine: e,
		ui:     ui,
	}
	var tetrisGame gotetromino.Game = &g
	return tetrisGame
}

func (g *game) Run() {
	stateChanges := g.engine.Start(g.ui.Action())
	for {
		state := <-stateChanges
		g.ui.Render(state)
		if state.Over {
			g.ui.Stop()
			break
		}
	}
}
