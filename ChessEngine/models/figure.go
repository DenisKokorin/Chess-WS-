package models

type Figure interface {
	PosibleMoves(board [8][8]Figure, x, y int) [][]int
	MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int
	GetColor() int
	GetMoves() [][]int
	GetName() string
}

func IsWithinBoard(x, y int) bool {
	return 0 <= x && x <= 7 && 0 <= y && y <= 7
}
