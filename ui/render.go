package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

func (a *app) render(s gotetromino.State) {
	a.renderer.SetState(s)
	a.renderer.Render()

	if s.ClearedPrevLevel {
		// make animation faster as more levels are cleared
		// TODO: Refactor limiting animation speed
		if a.tickDuration > 80*time.Millisecond {
			a.tickDuration = a.tickDuration - (80 * time.Millisecond)
			a.ticker.Reset(a.tickDuration)
		}
	}
}
