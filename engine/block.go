package engine

type Block int

const (
	Space Block = iota + 100
	Boundary
    ITetromino
    JTetromino
    LTetromino
    OTetromino
    STetromino
    TTetromino
    ZTetromino
)
