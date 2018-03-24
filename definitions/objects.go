package definitions

type Game struct {
	GameID         string `json:"gameID"`
	NextPlayerID   string `json:"nextPlayerID"`
	SecondPlayerID string `json:"secondPlayerID"`
	Board          Board
}

type Board struct {
	Squares [][]int `json:"squares"`
	Pawns   []Pawn  `json:"pawns"`
}

type Pawn struct {
	Type     PawnType `json:"type"`
	PlayerID string   `json:"playerID"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

type PawnType string

const (
	PawnMan   PawnType = "man"
	PawnWoman PawnType = "woman"
)
