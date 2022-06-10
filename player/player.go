package player

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type player struct {
	screen tcell.Screen
	action chan gotetromino.Action
}

func New(s tcell.Screen) gotetromino.Player {
	a := make(chan gotetromino.Action)
	p := player{
		action: a,
		screen: s,
	}
    // launch goroutine to listen to key events and send corresponding actions
	go func() {
		for {
			ev := p.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyLeft:
					p.action <- gotetromino.Left
				case tcell.KeyRight:
					p.action <- gotetromino.Right
				default:
					// TODO: Need to listen for key and talk to renderer and stop rendering
					return
				}
			}
		}
	}()
	var player gotetromino.Player = &p
	return player
}

func (p *player) Action() <-chan gotetromino.Action {
	return p.action
}
