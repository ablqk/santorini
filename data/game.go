package data

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/ablqk/santorini/lib/errors"
)

var games = map[string]Game{}

// Game represents a game.
type Game struct {
	GameID          string
	Finished        *time.Time
	Players         []Player
	nextPlayerIndex int
	Board           Board
}

// Save a game in db.
func (g *Game) Save() {
	games[g.GameID] = *g
}

//NextPlayer returns the next player.
func (g *Game) NextPlayer() Player {
	return g.Players[g.nextPlayerIndex]
}

// WaitingPlayer returns the other player.
func (g *Game) WaitingPlayer() Player {
	return g.Players[(g.nextPlayerIndex+1)%2]
}

// CreateGame creates a new game and saves it.
func CreateGame() (Game, *errors.HTTPError) {
	id := uuid.Must(uuid.NewV4())

	game := Game{
		GameID:  id.String(),
		Players: makePlayers(),
		Board:   makeBoard(),
	}

	games[id.String()] = game

	return game, nil
}

// FindGame returns an existing game.
func FindGame(id string) (*Game, *errors.HTTPError) {
	game, ok := games[id]
	if !ok {
		return nil, errors.New(errors.NotFoundE, "cannot find game with id " + id)
	}
	return &game, nil
}
