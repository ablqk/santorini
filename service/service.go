// Package service contains all the service methods.
package service

import (
	"github.com/ablqk/santorini/data"
	"errors"
	"github.com/ablqk/santorini/api"
)

const maxHeight = 4

// NewGame creates a new game and saves it.
func NewGame() (data.Game, error) {
	return data.CreateGame()
}

// FindGame returns the game with the given id.
func FindGame(gameID string) (data.Game, error) {
	game, err := data.FindGame(gameID)
	if err != nil {
		return data.Game{}, err
	}
	return *game, nil
}

// Play lets a player pley a turn.
func Play(gameID string, playerID string, turn api.PlayRequest) (data.Game, error) {
	game, err := FindGame(gameID)
	if err != nil {
		return data.Game{}, err
	}

	if playerID != game.NextPlayer().PlayerID {
		return data.Game{}, errors.New("400: not your turn")
	}

	if turn.Position.X >= 5 {
		return data.Game{}, errors.New("400: square out of bound, x = " + string(turn.Position.X))
	}

	if turn.Position.Y >= 5 {
		return data.Game{}, errors.New("400: square out of bound, y = " + string(turn.Position.Y))
	}

	switch turn.Verb {
	case api.BuildAction:
		game, err = playBuild(game, turn.Position.X, turn.Position.Y)
	}
	if err != nil {
		return data.Game{}, err
	}

	game.Save()
	return game, nil
}

func playBuild(game data.Game, x int, y int) (data.Game, error) {
	if game.Board.Squares[x][y] == maxHeight {
		return data.Game{}, errors.New("400: square height is at maximum (" + string(maxHeight) + ")")
	}
	game.Board.Squares[x][y]++
	return game, nil
}
