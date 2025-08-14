package chess

import (
	"main/ChessEngine/models"
)

type Game struct {
	ChessBoard Board
	CurrentP   string
	Over       bool
}

func NewGame() *Game {
	return &Game{
		ChessBoard: NewBoard(),
		CurrentP:   "white",
		Over:       false,
	}
}

func (g *Game) GetBoard() [8][8]string {
	return g.ChessBoard.PrintBoard()
}

func (g *Game) SwitchPlayer() {
	if g.CurrentP == "white" {
		g.CurrentP = "black"
	} else {
		g.CurrentP = "white"
	}
}

func (g *Game) KingCanEscape(color int) bool {
	var xKing, yKing int
	if color == 1 {
		xKing, yKing = g.ChessBoard.WhiteKing[0], g.ChessBoard.WhiteKing[1]
	} else {
		xKing, yKing = g.ChessBoard.BlackKing[0], g.ChessBoard.BlackKing[1]
	}

	for x, i := range g.ChessBoard.Cells {
		for y, j := range i {
			if j != nil && j.GetColor() == color {
				moves := j.MovesWhenAttacked(g.ChessBoard.Cells, x, y, xKing, yKing)
				if len(moves) != 0 {
					return true
				}
			}
		}
	}

	return false
}

func (g *Game) CheckAttack(color int) bool {
	var xKing, yKing int
	if color == 1 {
		xKing, yKing = g.ChessBoard.WhiteKing[0], g.ChessBoard.WhiteKing[1]
	} else {
		xKing, yKing = g.ChessBoard.BlackKing[0], g.ChessBoard.BlackKing[1]
	}

	king := g.ChessBoard.Cells[xKing][yKing].(*models.King)
	return king.IsAttacked(g.ChessBoard.Cells, xKing, yKing)
}

func (g *Game) CheckMate(color int) bool {
	if !g.CheckAttack(color) {
		return false
	}

	if g.KingCanEscape(color) {
		return false
	}

	return true
}

func (g *Game) StartGame(moveChan chan []int, boardChan chan [8][8]string, overChan chan bool, errorChan chan error) {
	for {
		select {
		case move := <-moveChan:
			flagErr := false
			x1, y1, x2, y2 := move[0], move[1], move[2], move[3]
			if g.CurrentP == "white" {
				if g.CheckAttack(1) {
					err := g.ChessBoard.MakeTurn(x1, y1, x2, y2, 1, true)
					if err != nil {
						errorChan <- err
						flagErr = true
					}
				} else {
					err := g.ChessBoard.MakeTurn(x1, y1, x2, y2, 1, false)
					if err != nil {
						errorChan <- err
						flagErr = true
					}
				}
				if g.CheckMate(0) {
					overChan <- true
					return
				}

			} else {
				if g.CheckAttack(0) {
					err := g.ChessBoard.MakeTurn(x1, y1, x2, y2, 0, true)
					if err != nil {
						errorChan <- err
						flagErr = true
					}
				} else {
					err := g.ChessBoard.MakeTurn(x1, y1, x2, y2, 0, false)
					if err != nil {
						errorChan <- err
						flagErr = true
					}
				}
				if g.CheckMate(1) {
					overChan <- true
					return
				}
			}

			if !flagErr {
				g.SwitchPlayer()
				boardChan <- g.GetBoard()
			}
		}
	}
}
