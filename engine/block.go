package engine

type Block int

const (
	Space Block = iota + 100
	ITetromino
	JTetromino
	LTetromino
	OTetromino
	STetromino
	TTetromino
	ZTetromino
)
