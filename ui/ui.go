package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type ui struct {
	engine       gotetromino.Engine
	user         gotetromino.User
	renderer     gotetromino.Renderer
	ticker       *time.Ticker
	tickDuration time.Duration
}

func New(e gotetromino.Engine, r gotetromino.Renderer, u gotetromino.User) gotetromino.UI {
	d := 800 * time.Millisecond
	t := time.NewTicker(d)
	var ui gotetromino.UI = &ui{
		engine:       e,
		user:         u,
		renderer:     r,
		ticker:       t,
		tickDuration: d,
	}
	return ui
}

func (ui *ui) Run() {
	actions := ui.user.Action()
	interactions := ui.user.Interaction()
	state := ui.engine.State()
	for {
		newState := ui.engine.State()
		if newState.Level-state.Level > 0 {
			ui.tickDuration = ui.tickDuration - time.Duration(newState.Level*int(ui.tickDuration/10))
			ui.ticker.Reset(ui.tickDuration)
		}
		state = newState
		ui.renderer.Render(state)
		select {
		case i := <-interactions:
			if i == gotetromino.Exit {
				ui.renderer.Stop()
				return
			}
			if i == gotetromino.Restart && ui.engine.State().Over {
				ui.engine.Reset()
			}
		case a := <-actions:
			ui.engine.Step(a)
		case <-ui.ticker.C:
			ui.engine.Step(gotetromino.None)
		}

	}
}
