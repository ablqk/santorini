// Package play defines the endpoint to play a turn.
package play

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/service"
	"encoding/json"
)

const (
	// Path is he path for this endpoint.
	Path = "/games/{" + api.GameIDParameter + "}/play/{" + api.PlayerIDParameter + "}"
)

type endpoint struct {
}

// Serve serves the request.
func (e endpoint) Serve(r *http.Request) (api.Response, error) {
	vars := mux.Vars(r)
	gameID := vars[api.GameIDParameter]
	playerID := vars[api.PlayerIDParameter]

	turn := api.PlayRequest{}
	err := json.NewDecoder(r.Body).Decode(&turn)
	if err != nil {
		return api.GameResponse{}, err
	}
	defer r.Body.Close()


	game, err := service.Play(gameID, playerID, turn)
	if err != nil {
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
	return http.MethodPost
}

// NewEndpoint creates the handler for this endpoint.
func NewEndpoint() api.Endpoint {
	return endpoint{}
}
