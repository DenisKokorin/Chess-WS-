package main

import (
	"encoding/json"
	"log"
	chess "main/chessEngine"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type BoardJSON struct {
	Type          string       `json:"type"`
	Board         [8][8]string `json:"board"`
	PlayerRole    string       `json:"playerRole"`
	CurrentPlayer string       `json:"currentPLayer"`
}

func getLobby(c echo.Context) error {
	keys := make([]string, 0, len(games))
	for key := range games {
		keys = append(keys, key)
	}
	return c.JSON(http.StatusOK, keys)
}

func handleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
		return err
	}

	gameID := c.QueryParam("gameID")
	if gameID == "" {
		conn.WriteJSON(map[string]string{
			"type":  "error",
			"error": "gameID is required",
		})
		return err
	}

	gamesLock.Lock()
	game, exists := games[gameID]
	if !exists {
		game = &GameRoom{
			Game:  *chess.NewGame(),
			White: nil,
			Black: nil,
		}
		games[gameID] = game
	}
	gamesLock.Unlock()

	conn.WriteJSON("connected")

	game.lock.Lock()
	var playerRole string
	if game.White == nil {
		game.White = conn
		playerRole = "white"
	} else if game.Black == nil {
		game.Black = conn
		playerRole = "black"
	} else {
		conn.WriteJSON(map[string]string{
			"type":  "error",
			"error": "game is full",
		})
		return err
	}
	game.lock.Unlock()

	for {
		if game.Black != nil && game.White != nil {
			break
		}
	}

	board := game.Game.GetBoard()
	jsonData := BoardJSON{
		Type:          "init",
		Board:         board,
		PlayerRole:    playerRole,
		CurrentPlayer: game.Game.CurrentP,
	}

	jsonInit, _ := json.Marshal(jsonData)
	conn.WriteMessage(websocket.TextMessage, jsonInit)

	log.Println(string(jsonInit))

	moveChan := make(chan []int)
	boardChan := make(chan [8][8]string)
	overChan := make(chan bool)
	errorChan := make(chan error)

	go game.Game.StartGame(moveChan, boardChan, overChan, errorChan)
	for {
		var msg map[string]string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error: ", err)
			break
		}

		if msg["type"] == "move" {
			currentTurn := game.Game.CurrentP
			if (currentTurn == "white" && playerRole != "white") || (currentTurn == "black" && playerRole != "black") {
				conn.WriteJSON(map[string]string{
					"type":  "error",
					"error": "Not your turn",
				})
				continue
			}

			var move []int
			json.Unmarshal([]byte(msg["move"]), &move)

			moveChan <- move
			select {
			case err := <-errorChan:
				conn.WriteJSON(map[string]string{
					"type":  "error",
					"error": err.Error(),
				})

			case board := <-boardChan:
				boardJson := BoardJSON{
					Type:          "update",
					Board:         board,
					PlayerRole:    playerRole,
					CurrentPlayer: game.Game.CurrentP,
				}
				jsonUpdate, _ := json.Marshal(boardJson)
				if game.White != nil {
					game.White.WriteMessage(websocket.TextMessage, jsonUpdate)
				}
				if game.Black != nil {
					game.Black.WriteMessage(websocket.TextMessage, jsonUpdate)
				}

			case _ = <-overChan:
				conn.WriteJSON(map[string]string{
					"type": "over",
					"msg":  "game over",
				})

				close(overChan)
				close(boardChan)
				close(errorChan)
				close(moveChan)

				game.lock.Lock()
				game.White.Close()
				game.White = nil
				game.Black.Close()
				game.Black = nil
				game.lock.Unlock()

				log.Print("game over")

				if game.White == nil && game.Black == nil {
					gamesLock.Lock()
					delete(games, gameID)
					gamesLock.Unlock()
				}
				break
			}
		}
	}

	return nil
}
