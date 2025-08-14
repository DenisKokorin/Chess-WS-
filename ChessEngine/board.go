package chess

import (
	"errors"
	"main/ChessEngine/models"
)

type Board struct {
	Cells     [8][8]models.Figure
	WhiteKing [2]int
	BlackKing [2]int
}

func NewBoard() Board {
	var board Board

	board.Cells[0] = [8]models.Figure{
		models.NewRook("Rook", 0),
		models.NewKnight("Knight", 0),
		models.NewBishop("Bishop", 0),
		models.NewQueen("Queen", 0),
		models.NewKing("King", 0),
		models.NewBishop("Bishop", 0),
		models.NewKnight("Knight", 0),
		models.NewRook("Rook", 0),
	}

	for i := 0; i < 8; i++ {
		board.Cells[1][i] = models.NewPawn("Pawn", 0)
	}

	board.Cells[7] = [8]models.Figure{
		models.NewRook("Rook", 1),
		models.NewKnight("Knight", 1),
		models.NewBishop("Bishop", 1),
		models.NewQueen("Queen", 1),
		models.NewKing("King", 1),
		models.NewBishop("Bishop", 1),
		models.NewKnight("Knight", 1),
		models.NewRook("Rook", 1),
	}

	for i := 0; i < 8; i++ {
		board.Cells[6][i] = models.NewPawn("Pawn", 1)
	}

	board.WhiteKing = [2]int{7, 4}
	board.BlackKing = [2]int{0, 4}

	return board
}

func (b *Board) Reset() {
	b.Cells[0] = [8]models.Figure{
		models.NewRook("Rook", 0),
		models.NewKnight("Knight", 0),
		models.NewBishop("Bishop", 0),
		models.NewQueen("Queen", 0),
		models.NewKing("King", 0),
		models.NewBishop("Bishop", 0),
		models.NewKnight("Knight", 0),
		models.NewRook("Rook", 0),
	}

	for i := 0; i < 8; i++ {
		b.Cells[1][i] = models.NewPawn("Pawn", 0)
	}

	b.Cells[7] = [8]models.Figure{
		models.NewRook("Rook", 1),
		models.NewKnight("Knight", 1),
		models.NewBishop("Bishop", 1),
		models.NewQueen("Queen", 1),
		models.NewKing("King", 1),
		models.NewBishop("Bishop", 1),
		models.NewKnight("Knight", 1),
		models.NewRook("Rook", 1),
	}

	for i := 0; i < 8; i++ {
		b.Cells[6][i] = models.NewPawn("Pawn", 1)
	}
}

func (b *Board) GetCells() [8][8]models.Figure {
	return b.Cells
}

func (b *Board) PrintBoard() [8][8]string {
	board := b.GetCells()
	var stringBoard [8][8]string
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board[x][y]
			if piece == nil {
				stringBoard[x][y] = ""
			} else {
				stringBoard[x][y] = piece.GetName()
			}
		}
	}
	return stringBoard
}

func (b *Board) Choose(x, y, color int, attack bool) ([][]int, error) {
	piece := b.Cells[x][y]
	if piece == nil {
		return nil, errors.New("invalid piece")
	}

	if piece.GetColor() != color {
		return nil, errors.New("invalid piece")
	}

	var moves [][]int
	if !attack {
		moves = piece.PosibleMoves(b.Cells, x, y)
	} else {
		var xKing, yKing int
		if color == 1 {
			xKing, yKing = b.WhiteKing[0], b.WhiteKing[1]
		} else {
			xKing, yKing = b.BlackKing[0], b.BlackKing[1]
		}
		moves = piece.MovesWhenAttacked(b.Cells, x, y, xKing, yKing)
	}

	return moves, nil
}

func (b *Board) ChangeFigure(x1, y1, x2, y2 int) {
	b.Cells[x2][y2] = b.Cells[x1][y1]
	b.Cells[x1][y1] = nil
}

func (b *Board) MakeTurn(x1, y1, x2, y2 int, color int, attack bool) error {

	moves, err := b.Choose(x1, y1, color, attack)
	if err != nil {
		return err
	}

	if len(moves) == 0 {
		return errors.New("no moves for this piece")
	}

	for _, pair := range moves {
		n, m := pair[0], pair[1]
		if n == x2 && m == y2 {
			b.ChangeFigure(x1, y1, x2, y2)
			b.FindKing()

			return nil
		}
	}

	return errors.New("invalid move")
}

func (b *Board) FindKing() {
	for n, i := range b.Cells {
		for m, j := range i {
			if j != nil {
				if j.GetName() == "King" {
					if j.GetColor() == 1 {
						b.WhiteKing = [2]int{n, m}
					} else {
						b.BlackKing = [2]int{n, m}
					}
				}
			}
		}
	}
}
