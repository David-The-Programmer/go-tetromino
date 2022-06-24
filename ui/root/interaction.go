package root

import (
	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type root struct {
	screen           tcell.Screen
	store            gotetromino.Store
	keyEventListener gotetromino.KeyEventListener
	children         []gotetromino.ChildUIComponent
	stop             chan bool
	position         []int
	dimensions       []int
}

func New(sc tcell.Screen, s gotetromino.Store, k gotetromino.KeyEventListener) gotetromino.ParentUIComponent {
	stop := make(chan bool)
	return &root{
		screen:           sc,
		store:            s,
		keyEventListener: k,
		stop:             stop,
	}
}

func (r *root) Run() {
	r.store.Start()
	r.keyEventListener.Start()
	r.screen.HideCursor()
	r.screen.DisableMouse()
	r.keyEventListener.Attach(r)
	for {
		select {
		case <-r.stop:
			r.Publish()
			r.store.Stop()
			r.keyEventListener.Stop()
			r.screen.Fini()
			return
		}
	}

}

func (r *root) Pos() []int {
	return r.position
}

func (r *root) SetPos(pos []int) {
	r.position = append([]int{}, pos...)
}

func (r *root) Dimensions() []int {
	return r.dimensions
}

func (r *root) SetDimensions(dimensions []int) {
	r.dimensions = append([]int{}, dimensions...)
}

func (r *root) Attach(c gotetromino.ChildUIComponent) {
	r.children = append(r.children, c)
}

func (r *root) Detach(c gotetromino.ChildUIComponent) {
	childIdx := 0
	for i := range r.children {
		if r.children[i] == c {
			childIdx = i
		}
	}
	retained := []gotetromino.ChildUIComponent{}
	retained = append(retained, r.children[:childIdx]...)
	retained = append(retained, r.children[childIdx+1:]...)
	r.children = retained
}

func (r *root) Publish() {
	for i := range r.children {
		r.children[i].HandleNewInteraction(gotetromino.Exit)
	}
}

func (r *root) HandleNewKey(k gotetromino.Key) {
	switch k {
	case gotetromino.Esc:
		r.stop <- true
	}
}
