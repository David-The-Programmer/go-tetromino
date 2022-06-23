package matrix

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type Matrix struct {
	screen        tcell.Screen
	engineService gotetromino.EngineService
	interaction   chan gotetromino.Interaction
	action        chan gotetromino.Action
	state         chan gotetromino.State
	ticker        *time.Ticker
	tickDuration  time.Duration
	position      []int
}

func New(s tcell.Screen, es gotetromino.EngineService, pos []int) *Matrix {
	s.HideCursor()
	s.DisableMouse()
	duration := 800 * time.Millisecond
	ticker := time.NewTicker(duration)
	interaction := make(chan gotetromino.Interaction)
	action := make(chan gotetromino.Action)
	state := make(chan gotetromino.State)
	m := &Matrix{
		screen:        s,
		engineService: es,
		interaction:   interaction,
		action:        action,
		state:         state,
		ticker:        ticker,
		tickDuration:  duration,
		position:      append([]int{}, pos...),
	}
	return m
}

func (m *Matrix) Run() {
	for {
		// newState := m.engine.State()
		// if newState.Level-state.Level > 0 {
		// 	m.tickDuration = m.tickDuration - time.Duration(newState.Level*int(m.tickDuration/10))
		// 	m.ticker.Reset(m.tickDuration)
		// }
		// state = newState
		select {
		case i := <-m.interaction:
			if i == gotetromino.Exit {
				stop(m.screen)
				return
			}
			if i == gotetromino.Restart && m.engineService.State().Over {
				m.engineService.Reset()
			}
		case a := <-m.action:
			m.engineService.Step(a)
		case <-m.ticker.C:
			m.engineService.Step(gotetromino.None)
		case s := <-m.state:
			render(m.screen, s)
		}

	}

}

func (m *Matrix) Subscribe(s gotetromino.Subject) {
    s.Register(m)
}

func (m *Matrix) Notify(v any) {
	switch t := v.(type) {
	case gotetromino.Key:
        m.keyEventHandler(t)
	case gotetromino.State:
		m.state <- t
	}
}

func (m *Matrix) keyEventHandler(k gotetromino.Key) {
    switch k {
    case gotetromino.Esc:
        m.interaction <- gotetromino.Exit
    case gotetromino.R:
        m.interaction <- gotetromino.Restart
    case gotetromino.X:
        m.action <- gotetromino.RotateCW
    case gotetromino.Z:
        m.action <- gotetromino.RotateACW
    case gotetromino.SpaceBar:
        m.action <- gotetromino.HardDrop
    case gotetromino.DownArrow:
        m.action <- gotetromino.SoftDrop
    case gotetromino.LeftArrow:
        m.action <- gotetromino.Left
    case gotetromino.RightArrow:
        m.action <- gotetromino.Right
    }
}
