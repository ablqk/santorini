// Package gamestate contains the data for the getgame endpoint.
package gamestate

import (
	"net/http"
	"github.com/ablqk/santorini/data"
	"github.com/ablqk/santorini/definitions"
	"github.com/ablqk/santorini/endpoints"
	"github.com/gorilla/mux"
)

const (
	// Path is he path for this endpoint.
	Path = "/games/{" + definitions.GameIDParameter + "}"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (definitions.Response, error) {

	vars := mux.Vars(r)
	gameID := vars[definitions.GameIDParameter]

	game, err := data.FindGame(gameID)
	if err != nil {
		// TODO return 404
		return definitions.GameResponse{}, err
	}

	resp := definitions.NewGameResponse(*game)
	return resp, nil
}

func (e endpoint) NominalResponse() int {
	return http.StatusOK
}

// Path implementation of Endpoint
func (e endpoint) Path() string {
	return Path
}

// Verb implementation of Endpoint
func (e endpoint) Verb() string {
	return http.MethodGet
}

// NewEndpoint creates the handler for this endpoint.
func NewEndpoint() endpoints.Endpoint {
	return endpoint{}
}
