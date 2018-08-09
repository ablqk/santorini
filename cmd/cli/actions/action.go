package actions

import (
	"github.com/ablqk/santorini/api"
	"bufio"
	"os"
	"fmt"
	strings2 "strings"
	"strconv"
)

type PlayerContext struct {
	GameID   string
	PlayerID string
}

func ReadAction(ctx PlayerContext) (api.PlayRequest, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	return parse(ctx, text)
}

func parse(ctx PlayerContext, str string) (api.PlayRequest, error) {
	str = strings2.TrimRight(str, "\n")
	strings := strings2.Split(str, " ")
	switch strings[0] {
	case "build":
		return parseBuild(ctx, strings)
	case "move":
	default:
	}
	return api.PlayRequest{}, nil
}

func parseBuild(ctx PlayerContext, strings []string) (api.PlayRequest, error) {
	fmt.Println("Building...")
	x, err := strconv.ParseInt(strings[1], 10, 8)
	if err != nil {
		fmt.Println(err)
		return api.PlayRequest{}, err
	}
	y, err := strconv.ParseInt(strings[2], 10, 8)
	if err != nil {
		fmt.Println(err)
		return api.PlayRequest{}, err
	}
	fmt.Println("on position", x, y)

	return api.PlayRequest{
		GameID:   ctx.GameID,
		PlayerID: ctx.PlayerID,
		PlayBody: api.PlayBody{
			Verb: api.BuildAction,
			Position: api.Square{X: uint8(x), Y: uint8(y)},
		}}, nil
}
