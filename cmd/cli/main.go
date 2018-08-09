package main

import (
	"fmt"
	"flag"
	"github.com/ablqk/santorini/client"
	"github.com/ablqk/santorini/api"
	"github.com/ablqk/santorini/cmd/cli/actions"
	"errors"
)

type gameContext struct {
	client client.Client
	player actions.PlayerContext
}

func main() {
	gameID := flag.String("game", "", "The game ID")
	//playerID := flag.String("playerID", "", "The player ID")
	port := flag.Int("port", 3000, "Port for the game server")
	flag.Parse()

	santocl, err := client.New(*port)
	if err != nil {
		fmt.Println("Cannot create client", err)
		return
	}
	gctx := gameContext{client: santocl}

	var game *api.Game

	if *gameID == "" {
		game, err = startNewGame(gctx)
		if err != nil {
			fmt.Println("Cannot create new game", err)
			return
		}
	} else {
		// TODO : retrieve game with this id
		// TODO set player ID recieved in param, check validity
	}

	gctx.player.PlayerID = game.NextPlayerID
	gctx.player.GameID = game.GameID

	play(gctx, game)
}

func startNewGame(gctx gameContext) (*api.Game, error) {
	gameResponse, err := gctx.client.New()
	if err != nil {
		fmt.Println("Error while calling New game:", err)
		return nil, errors.New("error while calling new game: " + err.Error())
	}
	return &gameResponse, nil
}

func printBoard(game api.Game) {
	// clear console
	//fmt.Print("\033c")

	for _, line := range game.Board.Squares {
		for _, height := range line {
			fmt.Print(height, " | ")
		}
		fmt.Println("\n-------------------")
	}
}

func play(gctx gameContext, game *api.Game) {
	for !game.IsFinished {
		printBoard(*game)
		action, err := actions.ReadAction(gctx.player)
		if err != nil {
			fmt.Println("Cannot read action", err)
			return
		}
		resp, err := gctx.client.Play(action)
		fmt.Println(err)
		if err != nil {
			fmt.Println("cannot call play endpoint:", err)
			return
		}
		game = &resp
	}
}
