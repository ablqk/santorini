package definitions

type board struct {
	Squares [][]int `json:"squares,omitempty"`
	Pawns   []pawn  `json:"pawns,omitempty"`
}

type pawn struct {
	Type     PawnType `json:"type"`
	PlayerID string   `json:"playerID"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

// PawnType represents the different types of pawn.
type PawnType string

const (
	// PawnWoman is the woman version.
	PawnWoman PawnType = "woman"
	// PawnMan is the man version.
	PawnMan PawnType = "man"
)
