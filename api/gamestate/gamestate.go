// Package gamestate contains the data for the getgame endpoint.
package gamestate

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/service"
)

const (
	// Path is he path for this endpoint.
	Path = "/games/{" + api.GameIDParameter + "}"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (api.Response, error) {

	vars := mux.Vars(r)
	gameID := vars[api.GameIDParameter]

	game, err := service.FindGame(gameID)
	if err != nil {
		// TODO return 404
		return api.GameResponse{}, err
	}

	resp := api.NewGameResponse(game)
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
func NewEndpoint() api.Endpoint {
	return endpoint{}
}
