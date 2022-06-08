package renderer

import (
	"github.com/David-The-Programmer/go-tetromino/engine"

	"github.com/gdamore/tcell/v2"
)

var blockCodeToColour = map[engine.Block]tcell.Color{
	engine.Space:      tcell.ColorBlack,
	engine.Boundary:   tcell.ColorWhite,
	engine.ITetromino: tcell.ColorDarkCyan,
	engine.JTetromino: tcell.ColorBlue,
	engine.LTetromino: tcell.ColorOrange,
	engine.OTetromino: tcell.ColorYellow,
	engine.STetromino: tcell.ColorGreen,
	engine.TTetromino: tcell.ColorPurple,
	engine.ZTetromino: tcell.ColorRed,
}

func ColourForBlock(b engine.Block) tcell.Color {
    if _, ok := blockCodeToColour[b]; !ok {
        return tcell.ColorDefault
    }
	return blockCodeToColour[b]
}

// TODO: Put map to map between tetris block codes (individual squares making up everything) to character (rune)
