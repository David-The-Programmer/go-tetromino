package ui

import (
	"log"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/renderer"
	"github.com/David-The-Programmer/go-tetromino/user"

	"github.com/gdamore/tcell/v2"
)

type ui struct {
	renderer gotetromino.Renderer
	user     gotetromino.User
}

func New() gotetromino.UI {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Panic(err)
	}
	if err = s.Init(); err != nil {
		log.Panic(err)
	}
	r := renderer.New(s)
	u := user.New(s)
	var ui gotetromino.UI = &ui{
		renderer: r,
		user:     u,
	}
	return ui
}

func (ui *ui) Render(s gotetromino.State) {
	ui.renderer.Render(s)
}

func (ui *ui) Stop() {
	ui.renderer.Stop()
}

func (ui *ui) Action() <-chan gotetromino.Action {
	return ui.user.Action()
}

func (ui *ui) Interaction() <-chan gotetromino.Interaction {
	return ui.user.Interaction()
}
