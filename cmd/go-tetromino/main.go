package main

import (
	"github.com/David-The-Programmer/go-tetromino/engine"
	"github.com/David-The-Programmer/go-tetromino/renderer"
)

func main() {
    r := renderer.New()
    e := engine.New(20, 10)
    r.Render(e.State())
}
