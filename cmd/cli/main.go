package main

import (
	"fmt"
	"flag"
	"github.com/ablqk/santorini/client"
	"github.com/ablqk/santorini/api"
)

type context struct {
	client client.Client
}

func main() {
	fmt.Print("\033c")

	gameID := flag.String("game", "", "The game ID")
	//playerID := flag.String("playerID", "", "The player ID")
	port := flag.Int("port", 3000, "Port for the game server")
	flag.Parse()

	santocl, err := client.New(*port)
	if err != nil {
		fmt.Println("Cannot create client", err)
		return
	}
	ctx := context{santocl}

	var game api.GameResponse

	if *gameID == "" {
		game, err = startNewGame(ctx)
		if err != nil {
			fmt.Println("Cannot create new game", err)
			return
		}
		gameID = &game.GameID
		//playerID = &game.NextPlayerID
		fmt.Println("Play with opponent :", game.SecondPlayerID)
	} else {
		// TODO
	}

	printBoard(game)
}

func startNewGame(ctx context) (api.GameResponse, error) {
	gameResponse, err := ctx.client.New()
	if err != nil {
		fmt.Println("Error while calling New game", err)
		return api.GameResponse{}, err
	}
	return gameResponse, nil
}

func printBoard(game api.GameResponse) {
	fmt.Println()
	for _, line := range game.Board.Squares {
		for _, height := range line {
			fmt.Print(height, " | ")
		}
		fmt.Println("\n-------------------")
	}
}
