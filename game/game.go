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

// TODO: Update run method to allow user to restart game after game over
func (g *game) Run() {
	stateChanges := g.engine.Start(g.ui.Action())
	userInteractions := g.ui.Interaction()
	for {
		select {
		case i := <-userInteractions:
            // stop engine and quit UI if user signals to exit
			if i == gotetromino.Exit {
				g.engine.Stop()
				g.ui.Stop()
                return
			}
		case state := <-stateChanges:
            // render whenever a change in state occurs
			g.ui.Render(state)
		}
	}
}
