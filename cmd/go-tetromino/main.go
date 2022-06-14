package main

import (
	"log"

	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/renderer"
	"github.com/David-The-Programmer/go-tetromino/ui"
	"github.com/David-The-Programmer/go-tetromino/user"

	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Panic(err)
	}
	if err = s.Init(); err != nil {
		log.Panic(err)
	}
	e := engine.New(20, 20)
	r := renderer.New(s)
	u := user.New(s)
	gameInterface := ui.New(e, r, u)
	gameInterface.Run()
}
