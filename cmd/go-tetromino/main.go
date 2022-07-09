package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/ui"

	"github.com/gdamore/tcell/v2"
)

func main() {
	rand.Seed(time.Now().Unix())
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Panic(err)
	}
	if err = screen.Init(); err != nil {
		log.Panic(err)
	}
	engine := engine.New()
	ui := ui.New(screen, engine)
	ui.Run()
}
