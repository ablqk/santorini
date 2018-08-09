package api

type action string

// BuildAction consists in building on one square
const BuildAction action = "build"

// PlayRequest gives the details of the play action.
type PlayRequest struct {
	GameID   string
	PlayerID string
	PlayBody
}

// PlayBody is the body of the play request.
type PlayBody struct {
	Verb     action    `json:"verb"`
	Position Square    `json:"position"`
	Pawn     *PawnType `json:"pawn"`
}

// Square defines the position of one square on the board.
type Square struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}
