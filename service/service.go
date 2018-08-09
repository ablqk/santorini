// Package service contains all the service methods.
package service

import (
	"github.com/ablqk/santorini/data"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/lib/errors"
)

// maxHeight is the maximum height for one square.
const maxHeight = 4

// NewGame creates a new game and saves it.
func NewGame() (data.Game, *errors.HTTPError) {
	return data.CreateGame()
}

// FindGame returns the game with the given id.
func FindGame(gameID string) (data.Game, *errors.HTTPError) {
	game, err := data.FindGame(gameID)
	if err != nil {
		return data.Game{}, errors.Wrap(err, "cannot find game")
	}
	return *game, nil
}

// Play lets a player play a turn.
func Play(gameID string, playerID string, turn api.PlayRequest) (data.Game, *errors.HTTPError) {
	game, err := FindGame(gameID)
	if err != nil {
		return data.Game{}, errors.Wrap(err, "cannot find game to play")
	}

	if playerID != game.NextPlayer().PlayerID {
		return data.Game{}, errors.New(errors.BadRequestE, "not your turn")
	}

	if turn.Position.X >= 5 {
		return data.Game{}, errors.New(errors.BadRequestE, "square out of bound, x = "+string(turn.Position.X))
	}

	if turn.Position.Y >= 5 {
		return data.Game{}, errors.New(errors.BadRequestE, "square out of bound, y = "+string(turn.Position.Y))
	}

	switch turn.Verb {
	case api.BuildAction:
		game, err = playBuild(game, turn.Position.X, turn.Position.Y)
		// TODO api.MoveAction
	}
	if err != nil {
		return data.Game{}, err
	}

	game.Save()
	return game, nil
}

func playBuild(game data.Game, x uint8, y uint8) (data.Game, *errors.HTTPError) {
	if game.Board.Squares[x][y] == maxHeight {
		return data.Game{}, errors.New(errors.BadRequestE, "square height is at maximum (%d)", maxHeight)
	}
	game.Board.Squares[x][y]++
	return game, nil
}
