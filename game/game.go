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
	e := engine.New(20, 20)
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
	userInteractions := g.ui.Interaction()
	state := gotetromino.State{}
	for {
		select {
		case i := <-userInteractions:
			// stop engine and quit UI if user signals to exit
			if i == gotetromino.Exit {
				g.engine.Stop()
				g.ui.Stop()
				return
			}
			if i == gotetromino.Restart && state.Over {
				g.engine.Reset()
			}
		case state = <-stateChanges:
			// render whenever a change in state occurs
			g.ui.Render(state)
		}
	}
	/*
		    select {
		    case <- userInteractions:
		        // exit & restart
		    case <- userActions:
			    Left
		        g.engine.Step(a Action) State

			    Right
		        g.engine.Step(a Action) State

			    Drop
		        g.engine.Step(a Action) State

			    Rotate
		        g.engine.Step(a Action) State
		    case <- ticker.C
		        NoAction
		        g.engine.Step(a Action) State
		    }
	*/
}
