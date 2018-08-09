// Package client contains the functions to call the service.
package client

import (
	"github.com/ablqk/santorini/api"
	"net/http"
	"github.com/ablqk/santorini/api/newgame"
	"strings"
	"encoding/json"
	"github.com/ablqk/santorini/api/gamestate"
	"github.com/ablqk/santorini/api/play"
	"bytes"
	"fmt"
	"github.com/ablqk/santorini/lib/errors"
)

type client struct {
	GameID   string
	PlayerID string
}

// New calls the New Game endpoint.
func (c client) New() (api.Game, *errors.HTTPError) {
	req, err := http.NewRequest(http.MethodPost, c.buildURL(newgame.Path), strings.NewReader(""))
	if err != nil {
		return api.Game{}, errors.New(0, err.Error())
	}
	return callApiDecodeGame(req)
}

// FindGame calls the Find Game endpoint.
func (c client) FindGame(gameID string) (api.Game, *errors.HTTPError) {
	req, err := http.NewRequest(http.MethodGet, c.buildURL(gamestate.Path), strings.NewReader(gameID))
	if err != nil {
		return api.Game{}, errors.NewFromType(0, "cannot create request")
	}
	return callApiDecodeGame(req)
}

// Play calls the Play endpoint.
func (c client) Play(request api.PlayRequest) (api.Game, *errors.HTTPError) {
	b, err := json.Marshal(request)
	if err != nil {
		return api.Game{}, errors.NewFromType(501, "cannot marshall play request")
	}

	path := strings.Replace(play.Path, "{"+api.GameIDParameter+"}", request.GameID, 1)
	path = strings.Replace(path, "{"+api.PlayerIDParameter+"}", request.PlayerID, 1)
	fmt.Println("DEBUG: path is", path)

	req, err := http.NewRequest(http.MethodPost, c.buildURL(path), bytes.NewReader(b))
	if err != nil {
		return api.Game{}, errors.NewFromType(501, "cannot create request")
	}

	return callApiDecodeGame(req)
}

func (c client) buildURL(path string) string {
	// TODO dynamic
	return "http://localhost:3000" + path
}

func callApiDecodeGame(req *http.Request) (api.Game, *errors.HTTPError) {
	// call http request
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return api.Game{}, errors.NewFromType(response.StatusCode, err.Error())
	}

	// decode response
	decoder := json.NewDecoder(response.Body)
	var game api.Game
	err = decoder.Decode(&game)
	if err != nil {
		return api.Game{}, errors.NewFromType(501, "cannot decode response: " + err.Error())
	}
	defer response.Body.Close()

	return game, nil
}
