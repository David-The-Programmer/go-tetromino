package store

import (
	"sync"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/engine"
)

func New(numMatrixRows int, numMatrixCols int) gotetromino.Store {
	action := make(chan gotetromino.Action)
	reset := make(chan bool)
	stop := make(chan bool)
	return &store{
		engine: engine.New(numMatrixRows, numMatrixCols),
		action: action,
		reset:  reset,
		stop:   stop,
	}
}

type store struct {
	engine   gotetromino.Engine
	action   chan gotetromino.Action
	reset    chan bool
	stop     chan bool
	handlers []gotetromino.StateEventHandler
	mutex    sync.Mutex
}

func (s *store) Step(a gotetromino.Action) {
	s.action <- a
}

func (s *store) Reset() {
	s.reset <- true
}

func (s *store) State() gotetromino.State {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.engine.State()
}

func (s *store) Attach(h gotetromino.StateEventHandler) {
	s.handlers = append(s.handlers, h)
}

func (s *store) Detach(h gotetromino.StateEventHandler) {
	handlerIdx := 0
	for i := range s.handlers {
		if s.handlers[i] == h {
			handlerIdx = i
		}
	}
	retained := []gotetromino.StateEventHandler{}
	retained = append(retained, s.handlers[:handlerIdx]...)
	retained = append(retained, s.handlers[handlerIdx+1:]...)
	s.handlers = retained
}

func (s *store) Publish() {
	for i := range s.handlers {
		s.handlers[i].HandleNewState(s.State())
	}
}

func (s *store) Start() {
	go func() {
		s.Publish()
		for {
			select {
			case <-s.stop:
				return
			case <-s.reset:
				s.engine.Reset()
				s.Publish()
			case a := <-s.action:
				s.engine.Step(a)
				s.Publish()
			}
		}
	}()
}

func (s *store) Stop() {
	s.stop <- true
}
