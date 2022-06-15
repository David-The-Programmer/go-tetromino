package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

type ui struct {
	engine   gotetromino.Engine
	user     gotetromino.User
	renderer gotetromino.Renderer
	ticker   *time.Ticker
}

func New(e gotetromino.Engine, r gotetromino.Renderer, u gotetromino.User) gotetromino.UI {
	t := time.NewTicker(800 * time.Millisecond)
	var ui gotetromino.UI = &ui{
		engine:   e,
		user:     u,
		renderer: r,
		ticker:   t,
	}
	return ui
}

func (ui *ui) Run() {
	actions := ui.user.Action()
	interactions := ui.user.Interaction()
	for {
		ui.renderer.Render(ui.engine.State())
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
