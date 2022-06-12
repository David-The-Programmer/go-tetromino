package user

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type user struct {
	screen      tcell.Screen
	action      chan gotetromino.Action
	interaction chan gotetromino.Interaction
}

func New(s tcell.Screen) gotetromino.User {
	a := make(chan gotetromino.Action)
	i := make(chan gotetromino.Interaction)
	u := user{
		screen:      s,
		action:      a,
		interaction: i,
	}
	// launch goroutine to listen to key events and send corresponding actions/interactions
	go func() {
		for {
			ev := u.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyDown:
					u.action <- gotetromino.Drop
				case tcell.KeyLeft:
					u.action <- gotetromino.Left
				case tcell.KeyRight:
					u.action <- gotetromino.Right
				case tcell.KeyEsc:
					u.interaction <- gotetromino.Exit
					return
				}
			}
		}
	}()
	var user gotetromino.User = &u
	return user
}

func (u *user) Action() <-chan gotetromino.Action {
	return u.action
}

func (u *user) Interaction() <-chan gotetromino.Interaction {
	return u.interaction
}
