package models

type King struct {
	Name      string
	Code      int
	Moves     [][]int
	FirstMove bool
}

func NewKing(n string, c int) *King {
	return &King{
		Name:      n,
		Code:      c,
		Moves:     [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}},
		FirstMove: true,
	}
}

func (k *King) GetName() string {
	return k.Name
}

func (k *King) GetColor() int {
	return k.Code
}

func (k *King) GetMoves() [][]int {
	return k.Moves
}

func (k *King) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)
	kingMoves := k.GetMoves()

	for _, pair := range kingMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		if IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil || target.GetColor() != k.GetColor() {
				moves = append(moves, []int{nX, nY})
			}
		}
	}

	return moves
}

func (k *King) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)
	kingMoves := k.GetMoves()

	for _, pair := range kingMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := xKing+dx, yKing+dy
		if IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil || target.GetColor() != k.GetColor() && target.GetName() != "King" {
				if !k.IsAttacked(board, nX, nY) {
					moves = append(moves, []int{nX, nY})
				}
			}
		}
	}

	return moves
}

func (k *King) IsAttacked(board [8][8]Figure, x, y int) bool {
	queenDirection := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, pair := range queenDirection {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		for IsWithinBoard(nX, nY) {
			piece := board[nX][nY]
			if piece != nil {
				if piece.GetColor() != k.GetColor() && (piece.GetName() == "Rook" || piece.GetName() == "Queen" || piece.GetName() == "Bishop") {
					return true
				}
				break
			}
			nX += dx
			nY += dy
		}
	}

	knightDirection := [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}
	for _, pair := range knightDirection {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		if IsWithinBoard(nX, nY) {
			piece := board[nX][nY]
			if piece != nil && piece.GetColor() != k.GetColor() && piece.GetName() == "Knight" {
				return true
			}
		}
	}

	var pawnDirection int
	if k.GetColor() == 1 {
		pawnDirection = -1
	} else {
		pawnDirection = 1
	}
	for _, dy := range []int{-1, 1} {
		nX, nY := x+pawnDirection, y+dy
		if IsWithinBoard(nX, nY) {
			piece := board[nX][nY]
			if piece != nil && piece.GetColor() != k.GetColor() && piece.GetName() == "Pawn" {
				return true
			}
		}
	}

	return false
}
