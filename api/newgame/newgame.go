// Package newgame defines the endpoint to create a new game.
package newgame

import (
	"net/http"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/service"
)

const (
	// Path is he path for this endpoint.
	Path = "/games"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (api.Response, error) {
	game, err := service.NewGame()
	if err != nil {
		return api.GameResponse{}, err
	}

	resp := api.NewGameResponse(game)
	return resp, nil
}

func (e endpoint) NominalResponse() int {
	return http.StatusCreated
}

// Path implementation of Endpoint
func (e endpoint) Path() string {
	return Path
}

// Verb implementation of Endpoint
func (e endpoint) Verb() string {
	return http.MethodPost
}

// NewEndpoint creates the handler for this endpoint.
func NewEndpoint() api.Endpoint {
	return endpoint{}
}
