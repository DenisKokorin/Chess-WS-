package main

import (
	chess "main/chessEngine"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GameRoom struct {
	Game  chess.Game
	White *websocket.Conn
	Black *websocket.Conn
	lock  sync.Mutex
}

var games = make(map[string]*GameRoom)
var gamesLock sync.Mutex

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/lobby", getLobby)
	e.GET("/ws", handleWebSocket)
	e.Logger.Fatal(e.Start(":1323"))
}
