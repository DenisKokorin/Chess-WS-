package models

type Pawn struct {
	Name  string
	Color int
	Moves [][]int
}

func NewPawn(n string, c int) *Pawn {
	return &Pawn{
		Name:  n,
		Color: c,
		Moves: [][]int{},
	}
}

func (p *Pawn) GetName() string {
	return p.Name
}

func (p *Pawn) GetColor() int {
	return p.Color
}

func (p *Pawn) GetMoves() [][]int {
	return p.Moves
}

func (p *Pawn) PosibleMoves(board [8][8]Figure, x, y int) [][]int {
	moves := make([][]int, 0)

	var pawnMoves int
	if p.GetColor() == 1 {
		pawnMoves = -1
	} else {
		pawnMoves = 1
	}

	if IsWithinBoard(x+pawnMoves, y) && board[x+pawnMoves][y] == nil {
		moves = append(moves, []int{x + pawnMoves, y})
		if x == 6 && p.GetColor() == 1 || x == 1 && p.GetColor() == 0 {
			if board[x+2*pawnMoves][y] == nil {
				moves = append(moves, []int{x + 2*pawnMoves, y})
			}
		}
	}

	direction := []int{-1, 1}
	for _, dy := range direction {
		if IsWithinBoard(x+pawnMoves, y+dy) {
			target := board[x+pawnMoves][y+dy]
			if target != nil && target.GetColor() != p.GetColor() {
				moves = append(moves, []int{x + pawnMoves, y + dy})
			}
		}
	}

	return moves
}

func (p *Pawn) MovesWhenAttacked(board [8][8]Figure, x, y int, xKing, yKing int) [][]int {
	moves := make([][]int, 0)

	var pawnMoves int
	if p.GetColor() == 1 {
		pawnMoves = -1
	} else {
		pawnMoves = 1
	}

	if IsWithinBoard(x+pawnMoves, y) && board[x+pawnMoves][y] == nil {
		newBoard := board
		newBoard[x+pawnMoves][y] = NewPawn("Pawn", p.GetColor())
		newBoard[x][y] = nil
		king := newBoard[xKing][yKing].(*King)
		if !king.IsAttacked(newBoard, xKing, yKing) {
			moves = append(moves, []int{x + pawnMoves, y})
		}

		if x == 6 && p.GetColor() == 1 || x == 1 && p.GetColor() == 0 {
			if IsWithinBoard(x+2*pawnMoves, y) && board[x+2*pawnMoves][y] == nil {
				newBoard := board
				newBoard[x+2*pawnMoves][y] = NewPawn("Pawn", p.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{x + 2*pawnMoves, y})
				}
			}
		}
	}

	direction := []int{-1, 1}
	for _, dy := range direction {
		if IsWithinBoard(x+pawnMoves, y+dy) {
			target := board[x+pawnMoves][y+dy]
			if target != nil && target.GetColor() != p.GetColor() {
				newBoard := board
				newBoard[x+pawnMoves][y+dy] = NewPawn("Pawn", p.GetColor())
				newBoard[x][y] = nil
				king := newBoard[xKing][yKing].(*King)
				if !king.IsAttacked(newBoard, xKing, yKing) {
					moves = append(moves, []int{x + pawnMoves, y + dy})
				}
			}
		}
	}

	return moves
}
