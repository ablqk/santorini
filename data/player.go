package data

import "github.com/satori/go.uuid"

// Player defines the data relative to one player.
type Player struct {
	PlayerID string
	Man      Pawn
	Woman    Pawn
}

func makePlayers() []Player {
	return []Player{newPlayer(), newPlayer()}
}

func newPlayer() Player {
	return Player{
		PlayerID: uuid.Must(uuid.NewV4()).String(),
	}
}

// Pawn describes the position of one pawn.
type Pawn struct {
	X int
	Y int
}
