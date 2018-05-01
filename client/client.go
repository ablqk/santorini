// Package client contains the functions to call the service.
package client

import (
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/data"
)

// Client is an http client for the Santorini service.
type Client interface {
	New() (api.GameResponse, error)
	Play(request api.PlayRequest) (api.GameResponse, error)
}

// New creates a new client.
func New(port int) (Client, error) {
	return client{}, nil
}

type client struct {
}

func (c client) New() (api.GameResponse, error) {
	game := data.Game{
		GameID:  "123456789",
		Players: []data.Player{{PlayerID: "987654"}, {PlayerID: "354921"}},
		Board:   makeBoard(),
	}

	return api.NewGameResponse(game), nil
}

func (c client) Play(request api.PlayRequest) (api.GameResponse, error) {
	return api.GameResponse{}, nil
}

func makeBoard() data.Board {
	board := make([][]int, 5)
	for i := range board {
		board[i] = make([]int, 5)
		for j := range board[i] {
			board[i][j] = 0
		}
	}
	return data.Board{Squares: board}
}