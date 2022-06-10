package ui

import (
	"log"

	gotetromino "github.com/David-The-Programmer/go-tetromino"
	"github.com/David-The-Programmer/go-tetromino/player"
	"github.com/David-The-Programmer/go-tetromino/renderer"
	"github.com/gdamore/tcell/v2"
)

type ui struct {
	screen tcell.Screen
    renderer gotetromino.Renderer
    player gotetromino.Player
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
    p := player.New(s)
    var ui gotetromino.UI = &ui{
        screen: s,
        renderer: r,
        player: p,
    }
    return ui
}

func (ui *ui) Render(s gotetromino.State) {
    ui.renderer.Render(s)
}

func (ui *ui) Action() <-chan gotetromino.Action {
    return ui.player.Action()
}

func (ui *ui) Stop() {
    ui.renderer.Stop()
}
