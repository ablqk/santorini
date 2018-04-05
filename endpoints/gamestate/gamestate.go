// Package gamestate contains the data for the getgame endpoint.
package gamestate

import (
	"fmt"
	"net/http"
	"github.com/ablqk/santorini/data"
	"github.com/ablqk/santorini/definitions"
	"github.com/gorilla/mux"
)

const (
	// Path is he path for this endpoint.
	Path = "/games/{" + definitions.GameIDParameter + "}"
	// Verb is he HTTP Verb for this endpoint.
	Verb = http.MethodGet
)

// NewHandler creates the handler for this endpoint.
func NewHandler() http.Handler {
	return handler{}
}

type handler struct {
}

// ServeHTTP serves the request.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars[definitions.GameIDParameter]

	err := gameState(w, gameID)
	if err != nil {
		http.Error(w, "error while retrieving game "+err.Error(), http.StatusInternalServerError)
	}

}

// gameState returns the state of the game at this moment.
func gameState(w http.ResponseWriter, gameID string) error {
	fmt.Println("Serving.")

	game, err := data.FindGame(gameID)
	if err != nil {
		return err
	}

	resp, err := definitions.NewGameResponse(*game).Marshal()
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", resp)

	w.Header().Set("Content-Type", "application/json")
	return nil
}
