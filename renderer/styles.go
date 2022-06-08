package renderer

import (
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

var blockCodeToColour = map[engine.Block]tcell.Color{
	engine.Space:      tcell.ColorDefault,
	engine.Boundary:   tcell.ColorWhite,
	engine.ITetromino: tcell.ColorDarkCyan,
	engine.JTetromino: tcell.ColorBlue,
	engine.LTetromino: tcell.ColorOrange,
	engine.OTetromino: tcell.ColorYellow,
	engine.STetromino: tcell.ColorGreen,
	engine.TTetromino: tcell.ColorPurple,
	engine.ZTetromino: tcell.ColorRed,
}

func colourForBlock(b engine.Block) tcell.Color {
    if _, ok := blockCodeToColour[b]; !ok {
        return tcell.ColorDefault
    }
	return blockCodeToColour[b]
}

var blockCodeToChar = map[engine.Block]rune{
	engine.Space:      ' ',
	engine.Boundary:   'x',
	engine.ITetromino: 'I',
	engine.JTetromino: 'J',
	engine.LTetromino: 'L',
	engine.OTetromino: 'O',
	engine.STetromino: 'S',
	engine.TTetromino: 'T',
	engine.ZTetromino: 'Z',
}

func charForBlock(b engine.Block) rune {
    if _, ok := blockCodeToChar[b]; !ok {
        return ' '
    }
	return blockCodeToChar[b]
}
