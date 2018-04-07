// Package newgame defines the endpoint to create a new game.
package newgame

import (
	"net/http"
	"github.com/ablqk/santorini/data"
	"github.com/ablqk/santorini/definitions"
	"github.com/ablqk/santorini/endpoints"
)

const (
	// Path is he path for this endpoint.
	Path = "/games"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (definitions.Response, error) {
	game, err := data.CreateGame()
	if err != nil {
		return definitions.GameResponse{}, err
	}

	resp := definitions.NewGameResponse(game)
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
func NewEndpoint() endpoints.Endpoint {
	return endpoint{}
}
