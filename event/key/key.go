package key

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type keyEventListener struct {
	screen    tcell.Screen
	event     chan tcell.Event
	key       chan gotetromino.Key
	observers []gotetromino.Observer
}

func New(s tcell.Screen) gotetromino.KeyEventListener {
	var k gotetromino.KeyEventListener = &keyEventListener{
		screen: s,
		key:    make(chan gotetromino.Key, 1),
	}
	return k
}

func (k *keyEventListener) Listen() {
	go func() {
		for {
			ev := k.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventInterrupt:
				return
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEsc:
					k.key <- gotetromino.Esc
					k.NotifyAll()
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'r':
						k.key <- gotetromino.R
						k.NotifyAll()
					case 'x':
						k.key <- gotetromino.X
						k.NotifyAll()
					case 'z':
						k.key <- gotetromino.Z
						k.NotifyAll()
					case ' ':
						k.key <- gotetromino.SpaceBar
						k.NotifyAll()
					}
				case tcell.KeyDown:
					k.key <- gotetromino.DownArrow
					k.NotifyAll()
				case tcell.KeyLeft:
					k.key <- gotetromino.LeftArrow
					k.NotifyAll()
				case tcell.KeyRight:
					k.key <- gotetromino.RightArrow
					k.NotifyAll()
				}
			}
		}
	}()
}

func (k *keyEventListener) Stop() {
	for {
		err := k.screen.PostEvent(tcell.NewEventInterrupt(nil))
		if err == nil {
			return
		}
	}

}

func (k *keyEventListener) Register(o gotetromino.Observer) {
	k.observers = append(k.observers, o)
}

func (k *keyEventListener) Unregister(o gotetromino.Observer) {
	observerIdx := 0
	for i := range k.observers {
		if k.observers[i] == o {
			observerIdx = i
		}
	}
	retained := []gotetromino.Observer{}
	retained = append(retained, k.observers[:observerIdx]...)
	retained = append(retained, k.observers[observerIdx+1:]...)
	k.observers = retained
}

func (k *keyEventListener) NotifyAll() {
	key := <-k.key
	for i := range k.observers {
		k.observers[i].Notify(key)
	}
}
