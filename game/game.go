package game

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/renderer"
)

type game struct {
	engine   gotetromino.Engine
	renderer gotetromino.Renderer
}

func New() gotetromino.Game {
	e := engine.New(20, 10)
	r := renderer.New()
	g := game{
		engine:   e,
		renderer: r,
	}
	var tetrisGame gotetromino.Game = &g
	return tetrisGame
}

func (g *game) Run() {
	action := make(chan gotetromino.Action)
	stateChanges := g.engine.Start(action)
	for {
		state := <-stateChanges
		g.renderer.Render(state)
		if state.Over {
			break
		}
	}
}
