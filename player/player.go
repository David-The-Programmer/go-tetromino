package player

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type player struct {
	screen    tcell.Screen
	action    chan gotetromino.Action
	observers []gotetromino.Observer
}

func New(s tcell.Screen) gotetromino.Subject {
	a := make(chan gotetromino.Action, 1)
	p := player{
		screen: s,
		action: a,
	}
	// launch goroutine to listen to key events and send corresponding actions
	go func() {
		for {
			ev := p.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				// TODO: Need to end this routine lol
				case tcell.KeyEsc:
					return
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'x':
						p.action <- gotetromino.RotateCW
						p.NotifyObservers()
					case 'z':
						p.action <- gotetromino.RotateACW
						p.NotifyObservers()
					case ' ':
						p.action <- gotetromino.HardDrop
						p.NotifyObservers()
					}
				case tcell.KeyDown:
					p.action <- gotetromino.SoftDrop
					p.NotifyObservers()
				case tcell.KeyLeft:
					p.action <- gotetromino.Left
					p.NotifyObservers()
				case tcell.KeyRight:
					p.action <- gotetromino.Right
					p.NotifyObservers()
				}
			}
		}
	}()
	var player gotetromino.Subject = &p
	return player
}

func (p *player) Register(o gotetromino.Observer) {
	p.observers = append(p.observers, o)
}

func (p *player) Unregister(o gotetromino.Observer) {
	observerIdx := 0
	for i := range p.observers {
		if p.observers[i] == o {
			observerIdx = i
		}
	}
	retained := []gotetromino.Observer{}
	retained = append(retained, p.observers[:observerIdx]...)
	retained = append(retained, p.observers[observerIdx+1:]...)
	p.observers = retained
}

func (p *player) NotifyObservers() {
	for i := range p.observers {
		p.observers[i].Notify()
	}
}

func (p *player) Action() gotetromino.Action {
	return <-p.action
}
