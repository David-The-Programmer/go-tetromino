package key

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type keyEventListener struct {
	screen   tcell.Screen
	key      chan gotetromino.Key
	handlers []gotetromino.KeyEventHandler
}

func NewListener(s tcell.Screen) gotetromino.KeyEventListener {
	return &keyEventListener{
		screen: s,
		key:    make(chan gotetromino.Key, 1),
	}
}

func (k *keyEventListener) Start() {
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
					k.Publish()
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'r':
						k.key <- gotetromino.R
						k.Publish()
					case 'x':
						k.key <- gotetromino.X
						k.Publish()
					case 'z':
						k.key <- gotetromino.Z
						k.Publish()
					case ' ':
						k.key <- gotetromino.SpaceBar
						k.Publish()
					}
				case tcell.KeyDown:
					k.key <- gotetromino.DownArrow
					k.Publish()
				case tcell.KeyLeft:
					k.key <- gotetromino.LeftArrow
					k.Publish()
				case tcell.KeyRight:
					k.key <- gotetromino.RightArrow
					k.Publish()
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

func (k *keyEventListener) Attach(h gotetromino.KeyEventHandler) {
	k.handlers = append(k.handlers, h)
}

func (k *keyEventListener) Detach(h gotetromino.KeyEventHandler) {
	handlerIdx := 0
	for i := range k.handlers {
		if k.handlers[i] == h {
			handlerIdx = i
		}
	}
	retained := []gotetromino.KeyEventHandler{}
	retained = append(retained, k.handlers[:handlerIdx]...)
	retained = append(retained, k.handlers[handlerIdx+1:]...)
	k.handlers = retained
}

func (k *keyEventListener) Publish() {
	key := <-k.key
	for i := range k.handlers {
		k.handlers[i].HandleNewKey(key)
	}
}
