// Package newgame defines the endpoint to create a new game.
package newgame

import (
	"fmt"
	"net/http"
	"github.com/ablqk/santorini/data"
	"github.com/ablqk/santorini/definitions"
)

const (
	// Path is he path for this endpoint.
	Path = "/games"
	// Verb is he HTTP Verb for this endpoint.
	Verb = http.MethodPost
)

// NewHandler creates the handler for this endpoint.
func NewHandler() http.Handler {
	return handler{}
}

type handler struct {
}

// ServeHTTP serves the request.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := newGame(w, r)
	if err != nil {
		http.Error(w, "error while creating new game", http.StatusInternalServerError)
	}
}

func newGame(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Serving.")

	game, err := data.CreateGame()
	if err != nil {
		return err
	}

	resp, err := definitions.NewGameResponse(game).Marshal()
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", resp)

	// header
	w.Header().Set("Content-Type", "application/json")
	return nil
}
