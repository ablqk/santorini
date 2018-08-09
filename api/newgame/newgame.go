// Package newgame defines the endpoint to create a new game.
package newgame

import (
	"net/http"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/service"
	"github.com/ablqk/santorini/lib/errors"
)

const (
	// Path is he path for this endpoint.
	Path = "/games"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (api.Response, *errors.HTTPError) {
	game, err := service.NewGame()
	if err != nil {
		return api.Game{}, errors.Wrap(err, "cannot create new game")
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
