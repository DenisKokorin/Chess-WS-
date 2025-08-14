package models

type Rook struct {
	Name      string
	Color     int
	Moves     [][]int
	FirstMove bool
}

func NewRook(n string, c int) *Rook {
	return &Rook{
		Name:      n,
		Color:     c,
		Moves:     [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		FirstMove: true,
	}
}

func (r *Rook) GetName() string {
	return r.Name
}

func (r *Rook) GetColor() int {
	return r.Color
}

func (r *Rook) GetMoves() [][]int {
	return r.Moves
}

func (r *Rook) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)
	rookMoves := r.GetMoves()

	for _, pair := range rookMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy

		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				moves = append(moves, []int{nX, nY})
			} else {
				if target.GetColor() != r.GetColor() {
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

func (r *Rook) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)
	rookMoves := r.GetMoves()

	for _, pair := range rookMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy

		for IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target == nil {
				newBoard := board
				newBoard[nX][nY] = NewRook("Rook", r.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{nX, nY})
				}
			} else {
				if target.GetColor() != r.GetColor() {
					newBoard := board
					newBoard[nX][nY] = NewRook("Rook", r.GetColor())
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
