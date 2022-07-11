package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type app struct {
	screen       tcell.Screen
	engine       gotetromino.Engine
	ticker       *time.Ticker
	tickDuration time.Duration
	renderer     *renderer
}

func New(sc tcell.Screen, e gotetromino.Engine) gotetromino.App {
	duration := 800 * time.Millisecond
	ticker := time.NewTicker(duration)
	renderer := newRenderer(sc, e.State())
	return &app{
		screen:       sc,
		engine:       e,
		ticker:       ticker,
		tickDuration: duration,
		renderer:     renderer,
	}
}

func (a *app) Run() {
	eventLoop := make(chan tcell.Event)
	exit := make(chan struct{})
	go a.screen.ChannelEvents(eventLoop, exit)
	// render starting state
	a.render(a.engine.State())
	for {
		select {
		case ev := <-eventLoop:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEsc:
					a.screen.Fini()
					close(exit)
					return
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'r':
						if a.engine.State().Over {
							a.engine.Reset()
							a.ticker.Reset(a.tickDuration)
						}
						a.render(a.engine.State())
					case 'x':
						a.engine.Step(gotetromino.RotateCW)
						a.render(a.engine.State())
					case 'z':
						a.engine.Step(gotetromino.RotateACW)
						a.render(a.engine.State())
					case ' ':
						a.engine.Step(gotetromino.HardDrop)
						a.render(a.engine.State())
					}
				case tcell.KeyDown:
					a.engine.Step(gotetromino.SoftDrop)
					a.render(a.engine.State())
				case tcell.KeyLeft:
					a.engine.Step(gotetromino.Left)
					a.render(a.engine.State())
				case tcell.KeyRight:
					a.engine.Step(gotetromino.Right)
					a.render(a.engine.State())
				}
			}
		case <-a.ticker.C:
			if !a.engine.State().Over {
				a.engine.Step(gotetromino.None)
				a.render(a.engine.State())
			}
		}
	}

}

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
