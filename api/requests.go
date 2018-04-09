package api

type action string

// BuildAction consists in building on one square
const BuildAction action = "build"

// PlayRequest gives the details of the play action.
type PlayRequest struct {
	Verb action `json:"verb"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	Pawn *PawnType `json:"pawn"`
}
