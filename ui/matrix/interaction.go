package matrix

import (
	"time"

	gotetromino "github.com/David-The-Programmer/go-tetromino"

	"github.com/gdamore/tcell/v2"
)

type matrix struct {
	screen           tcell.Screen
	root             gotetromino.ParentUIComponent
	store            gotetromino.Store
	KeyEventListener gotetromino.KeyEventListener
	interaction      chan gotetromino.Interaction
	action           chan gotetromino.Action
	state            chan gotetromino.State
	ticker           *time.Ticker
	tickDuration     time.Duration
	position         []int
	dimensions       []int
}

func New(sc tcell.Screen, r gotetromino.ParentUIComponent, s gotetromino.Store, k gotetromino.KeyEventListener) gotetromino.ChildUIComponent {
	duration := 800 * time.Millisecond
	ticker := time.NewTicker(duration)
	interaction := make(chan gotetromino.Interaction)
	action := make(chan gotetromino.Action)
	state := make(chan gotetromino.State)
	return &matrix{
		screen:           sc,
		root:             r,
		store:            s,
		KeyEventListener: k,
		interaction:      interaction,
		action:           action,
		state:            state,
		ticker:           ticker,
		tickDuration:     duration,
	}
}

func (m *matrix) Run() {
	m.root.Attach(m)
	m.store.Attach(m)
	m.KeyEventListener.Attach(m)

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
				return
			}
			if i == gotetromino.Restart && m.store.State().Over {
				m.store.Reset()
			}
		case a := <-m.action:
			m.store.Step(a)
		case <-m.ticker.C:
			m.store.Step(gotetromino.None)
		case s := <-m.state:
			render(m.screen, s)
		}

	}

}

func (m *matrix) Pos() []int {
	return m.position
}

func (m *matrix) SetPos(pos []int) {
	m.position = append([]int{}, pos...)
}

func (m *matrix) Dimensions() []int {
	return m.dimensions
}

func (m *matrix) SetDimensions(dimensions []int) {
	m.dimensions = append([]int{}, dimensions...)
}

func (m *matrix) HandleNewState(s gotetromino.State) {
	m.state <- s
}

func (m *matrix) HandleNewKey(k gotetromino.Key) {
	switch k {
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

func (m *matrix) HandleNewInteraction(i gotetromino.Interaction) {
	m.interaction <- i
}
