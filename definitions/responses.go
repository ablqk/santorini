package definitions

import (
	"github.com/ablqk/santorini/data"
	"encoding/json"
)

// GameResponse dscribes a game and its state.
type GameResponse struct {
	GameID         string `json:"gameID"`
	NextPlayerID   string `json:"nextPlayerID"`
	SecondPlayerID string `json:"secondPlayerID"`
	Board          board  `json:"board"`
}

// Marshal transforms the GameResponse object into its json representation.
func (gr GameResponse) Marshal() ([]byte, error) {
	resp, err := json.Marshal(gr)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewGameResponse creates a GameResponse object.
func NewGameResponse(game data.Game) GameResponse {
	return GameResponse{
		GameID:         game.GameID,
		NextPlayerID:   game.Players[0].PlayerID,
		SecondPlayerID: game.Players[1].PlayerID,
	}
}
