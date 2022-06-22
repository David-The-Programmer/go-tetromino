package engine

import (
	"sync"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
)

func NewService(e gotetromino.Engine) gotetromino.EngineService {
	action := make(chan gotetromino.Action)
	reset := make(chan bool)
	stop := make(chan bool)
	es := engineService{
		engine: e,
		action: action,
		reset:  reset,
		stop:   stop,
	}
	go func() {
		for {
			select {
			case <-es.stop:
				return
			case <-es.reset:
				es.engine.Reset()
				es.NotifyObservers()
			case a := <-es.action:
				es.engine.Step(a)
				es.NotifyObservers()
			}
		}
	}()
	var service gotetromino.EngineService = &es
	return service
}

type engineService struct {
	engine    gotetromino.Engine
	action    chan gotetromino.Action
	reset     chan bool
	stop      chan bool
	observers []gotetromino.Observer
	mutex     sync.Mutex
}

func (es *engineService) Step(a gotetromino.Action) {
	es.action <- a
}

func (es *engineService) Reset() {
	es.reset <- true
}

func (es *engineService) State() gotetromino.State {
	es.mutex.Lock()
	state := es.engine.State()
	es.mutex.Unlock()
	return state
}

func (es *engineService) Stop() {
	es.stop <- true
}

func (es *engineService) Register(o gotetromino.Observer) {
	es.observers = append(es.observers, o)
}

func (es *engineService) Unregister(o gotetromino.Observer) {
	observerIdx := 0
	for i := range es.observers {
		if es.observers[i] == o {
			observerIdx = i
		}
	}
	retained := []gotetromino.Observer{}
	retained = append(retained, es.observers[:observerIdx]...)
	retained = append(retained, es.observers[observerIdx+1:]...)
	es.observers = retained
}

func (es *engineService) NotifyObservers() {
	for i := range es.observers {
		es.observers[i].Notify()
	}
}
