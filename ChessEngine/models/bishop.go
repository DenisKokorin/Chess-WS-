package models

type Bishop struct {
	Name  string
	Color int
	Moves [][]int
}

func NewBishop(n string, c int) *Bishop {
	return &Bishop{
		Name:  n,
		Color: c,
		Moves: [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
	}
}

func (b *Bishop) GetName() string {
	return b.Name
}

func (b *Bishop) GetColor() int {
	return b.Color
}

func (b *Bishop) GetMoves() [][]int {
	return b.Moves
}

func (b *Bishop) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)
	bishopMoves := b.GetMoves()

	for _, pair := range bishopMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				moves = append(moves, []int{nX, nY})
			} else {
				if target.GetColor() != b.GetColor() {
					moves = append(moves, []int{nX, nY})
				}
				break
			}
			nX += dx
			nY += dy
		}
	}

	return moves
}

func (b *Bishop) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)
	bishopMoves := b.GetMoves()

	for _, pair := range bishopMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				newBoard := board
				newBoard[nX][nY] = NewRook("Bishop", b.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{nX, nY})
				}
			} else {
				if target.GetColor() != b.GetColor() {
					newBoard := board
					newBoard[nX][nY] = NewRook("Bishop", b.GetColor())
					newBoard[x][y] = nil
					king := newBoard[xKing][yKing].(*King)
					if !king.IsAttacked(newBoard, xKing, yKing) {
						moves = append(moves, []int{nX, nY})
					}
				}
				break
			}
			nX += dx
			nY += dy
		}
	}

	return moves
}
