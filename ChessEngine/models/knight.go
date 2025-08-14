package models

type Knight struct {
	Name  string
	Color int
	Moves [][]int
}

func NewKnight(n string, c int) *Knight {
	return &Knight{
		Name:  n,
		Color: c,
		Moves: [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}},
	}
}

func (k *Knight) GetName() string {
	return k.Name
}

func (k *Knight) GetColor() int {
	return k.Color
}

func (k *Knight) GetMoves() [][]int {
	return k.Moves
}

func (k *Knight) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)
	knightMoves := k.GetMoves()

	for _, pair := range knightMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		if IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target != nil {
				if target.GetColor() != k.GetColor() {
					moves = append(moves, []int{nX, nY})
				}
			} else {
				moves = append(moves, []int{nX, nY})
			}
		}
	}
	return moves
}

func (k *Knight) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)
	knightMoves := k.GetMoves()

	for _, pair := range knightMoves {
		dx, dy := pair[0], pair[1]
		nX, nY := x+dx, y+dy
		if IsWithinBoard(nX, nY) {
			target := board[nX][nY]
			if target != nil {
				if target.GetColor() != k.GetColor() {
					newBoard := board
					newBoard[nX][nY] = NewKnight("Knight", k.GetColor())
					newBoard[x][y] = nil
					king := newBoard[xKing][yKing].(*King)
					if !king.IsAttacked(newBoard, xKing, yKing) {
						moves = append(moves, []int{nX, nY})
					}
				}
			} else {
				newBoard := board
				newBoard[nX][nY] = NewKnight("Knight", k.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{nX, nY})
				}
			}
		}
	}
	return moves
}
