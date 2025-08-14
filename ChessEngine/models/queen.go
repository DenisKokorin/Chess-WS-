package models

type Queen struct {
	Name  string
	Color int
	Moves [][]int
}

func NewQueen(n string, c int) *Queen {
	return &Queen{
		Name:  n,
		Color: c,
		Moves: [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
	}
}

func (q *Queen) GetName() string {
	return q.Name
}

func (q *Queen) GetColor() int {
	return q.Color
}

func (q *Queen) GetMoves() [][]int {
	return q.Moves
}

func (q *Queen) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)
	queenMoves := q.GetMoves()

	for _, pair := range queenMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy

		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				moves = append(moves, []int{nX, nY})
			} else {
				if target.GetColor() != q.GetColor() {
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

func (q *Queen) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)
	queenMoves := q.GetMoves()

	for _, pair := range queenMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy

		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				newBoard := board
				newBoard[nX][nY] = NewQueen("Queen", q.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{nX, nY})
				}
			} else {
				if target.GetColor() != q.GetColor() {
					newBoard := board
					newBoard[nX][nY] = NewQueen("Queen", q.GetColor())
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
