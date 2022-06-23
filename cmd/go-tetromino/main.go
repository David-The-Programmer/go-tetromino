package main

import (
	"log"

	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/event/key"
	"github.com/David-The-Programmer/go-tetromino/ui/matrix"

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
	es := engine.NewService(e)
    k := key.New(s)
    k.Listen()
	m := matrix.New(s, es, []int{10, 10})
    m.Subscribe(k)
    m.Subscribe(es)
	m.Run()
}
