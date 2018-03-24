package data

import (
	"time"
	"github.com/satori/go.uuid"
)

var games = map[string]Game{}

type Game struct {
	GameID          string
	Finished        *time.Time
	Players         []Player
	nextPlayerIndex int
	Board           Board
}

func (g *Game) Save() {
	games[g.GameID] = *g
}

func CreateGame() (Game, error) {
	id := uuid.Must(uuid.NewV4())

	game := Game{
		GameID:  id.String(),
		Players: makePlayers(),
		Board:   makeBoard(),
	}

	games[id.String()] = game

	return game, nil
}

//func LoadGame(id string) (*Game, error) {
//	game, ok := games[id]
//	if !ok {
//		return nil, errors.New("Cannot find game with id " + id)
//	}
//	return &game, nil
//}
