package api

import (
	"github.com/ablqk/santorini/data"
	"encoding/json"
)

// Response Represents an API response.
type Response interface {
	Marshal() ([]byte, error)
}

// Game describes a game and its state.
type Game struct {
	GameID         string `json:"gameID"`
	NextPlayerID   string `json:"nextPlayerID"`
	SecondPlayerID string `json:"secondPlayerID"`
	Board          board  `json:"board"`
	IsFinished     bool   `json:"isFinished"`
}

// Marshal transforms the Game object into its json representation.
func (gr Game) Marshal() ([]byte, error) {
	resp, err := json.Marshal(gr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewGameResponse creates a Game object.
func NewGameResponse(game data.Game) Game {
	return Game{
		GameID:         game.GameID,
		NextPlayerID:   game.Players[0].PlayerID,
		SecondPlayerID: game.Players[1].PlayerID,
		Board: board{
			Squares: game.Board.Squares,
		},
	}
}
