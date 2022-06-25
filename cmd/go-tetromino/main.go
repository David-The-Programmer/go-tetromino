package main

import (
	"log"

	"github.com/David-The-Programmer/go-tetromino/event/key"
	"github.com/David-The-Programmer/go-tetromino/store"
	"github.com/David-The-Programmer/go-tetromino/ui"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Panic(err)
	}
	if err = screen.Init(); err != nil {
		log.Panic(err)
	}
	store := store.New(20, 20)
	keyEventListener := key.NewListener(screen)
	ui := ui.New(screen, store, keyEventListener)
	ui.Run()
}
