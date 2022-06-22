package user

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type user struct {
	screen      tcell.Screen
	interaction chan gotetromino.Interaction
	observers   []gotetromino.Observer
}

func New(s tcell.Screen) gotetromino.User {
	// Give channel capacity of 1 to prevent blocking when sending to channel
	i := make(chan gotetromino.Interaction, 1)
	u := user{
		screen:      s,
		interaction: i,
	}
	// launch goroutine to listen to key events and send corresponding actions/interactions
	go func() {
		for {
			ev := u.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEsc:
					u.interaction <- gotetromino.Exit
					u.NotifyObservers()
					return
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'r':
						u.interaction <- gotetromino.Restart
						u.NotifyObservers()
					}
				}
			}
		}
	}()
	var user gotetromino.User = &u
	return user
}

func (u *user) Register(o gotetromino.Observer) {
	u.observers = append(u.observers, o)
}

func (u *user) Unregister(o gotetromino.Observer) {
	observerIdx := 0
	for i := range u.observers {
		if u.observers[i] == o {
			observerIdx = i
		}
	}
	retained := []gotetromino.Observer{}
	retained = append(retained, u.observers[:observerIdx]...)
	retained = append(retained, u.observers[observerIdx+1:]...)
	u.observers = retained
}

func (u *user) NotifyObservers() {
	for i := range u.observers {
		u.observers[i].Notify()
	}
}

func (u *user) Interaction() gotetromino.Interaction {
	return <-u.interaction
}
