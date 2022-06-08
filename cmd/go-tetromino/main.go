package main

import (
	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/renderer"
)

func main() {
    e := engine.New(20, 10)
    r := renderer.New(e)
    r.Render()
}
