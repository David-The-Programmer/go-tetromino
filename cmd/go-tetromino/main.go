package main

import (
	"log"

	"github.com/David-The-Programmer/go-tetromino/event/key"
	"github.com/David-The-Programmer/go-tetromino/store"
	"github.com/David-The-Programmer/go-tetromino/ui/matrix"
	"github.com/David-The-Programmer/go-tetromino/ui/root"

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
	rootComponent := root.New(screen, store, keyEventListener)
	matrixComponent := matrix.New(screen, rootComponent, store, keyEventListener)
	go matrixComponent.Run()
	rootComponent.Run()
}
