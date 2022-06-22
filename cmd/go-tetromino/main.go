package main

import (
	"log"

	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/ui/matrix"
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
	es := engine.NewService(e)
	u := user.New(s)
	m := matrix.New(s, es, u)
	m.Run()
}
