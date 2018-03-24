// Package newgame defines the endpoint to create a new game.
package newgame

import (
	"fmt"
	"net/http"
	"github.com/ablqk/santorini/data"
	"encoding/json"
	"github.com/ablqk/santorini/definitions"
)

const (
	Path = "/games"
	Verb = http.MethodPost
)

func NewHandler() http.Handler {
	return handler{}
}

type handler struct {
}

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

	responseGame := definitions.Game{
		GameID:         game.GameID,
		NextPlayerID:   game.Players[0].PlayerID,
		SecondPlayerID: game.Players[1].PlayerID,
	}

	resp, err := json.Marshal(definitions.NewGameResponse{
		Game: responseGame,
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\n", resp)
	return nil
}
