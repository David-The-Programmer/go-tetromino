package ui

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type ui struct {
	screen           tcell.Screen
	store            gotetromino.Store
	keyEventListener gotetromino.KeyEventListener
	interaction      chan gotetromino.Interaction
	action           chan gotetromino.Action
	state            chan gotetromino.State
	ticker           *time.Ticker
	tickDuration     time.Duration
}

func New(sc tcell.Screen, s gotetromino.Store, k gotetromino.KeyEventListener) gotetromino.UI {
	duration := 800 * time.Millisecond
	ticker := time.NewTicker(duration)
	interaction := make(chan gotetromino.Interaction)
	action := make(chan gotetromino.Action)
	state := make(chan gotetromino.State)
	return &ui{
		screen:           sc,
		store:            s,
		keyEventListener: k,
		interaction:      interaction,
		action:           action,
		state:            state,
		ticker:           ticker,
		tickDuration:     duration,
	}
}

func (u *ui) Run() {
	u.store.Attach(u)
	u.keyEventListener.Attach(u)

	u.store.Start()
	u.keyEventListener.Start()

	for {
		// newState := u.engine.State()
		// if newState.Level-state.Level > 0 {
		// 	u.tickDuration = u.tickDuration - time.Duration(newState.Level*int(u.tickDuration/10))
		// 	u.ticker.Reset(u.tickDuration)
		// }
		// state = newState
		select {
		case i := <-u.interaction:
			if i == gotetromino.Exit {
				u.keyEventListener.Stop()
				u.store.Stop()
				u.screen.Fini()
				return
			}
			if i == gotetromino.Restart && u.store.State().Over {
				u.store.Reset()
			}
		case a := <-u.action:
			u.store.Step(a)
		case <-u.ticker.C:
			u.store.Step(gotetromino.None)
		case s := <-u.state:
			u.render(s)
		}
	}

}

func (u *ui) HandleNewKey(k gotetromino.Key) {
	switch k {
	case gotetromino.Esc:
		u.interaction <- gotetromino.Exit
	case gotetromino.R:
		u.interaction <- gotetromino.Restart
	case gotetromino.X:
		u.action <- gotetromino.RotateCW
	case gotetromino.Z:
		u.action <- gotetromino.RotateACW
	case gotetromino.SpaceBar:
		u.action <- gotetromino.HardDrop
	case gotetromino.DownArrow:
		u.action <- gotetromino.SoftDrop
	case gotetromino.LeftArrow:
		u.action <- gotetromino.Left
	case gotetromino.RightArrow:
		u.action <- gotetromino.Right
	}
}

func (u *ui) HandleNewState(s gotetromino.State) {
	u.state <- s
}
