package api

import "github.com/ablqk/santorini/lib/errors"

// Santorini represents the santorini API.
type Santorini interface {
	New() (Game, *errors.HTTPError)
	FindGame(gameID string) (Game, *errors.HTTPError)
	Play(request PlayRequest) (Game, *errors.HTTPError)
}
