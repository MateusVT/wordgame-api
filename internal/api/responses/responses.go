package responses

import (
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
)

// swagger:model gameResponse
type GameResponse struct {
	// The ID of the Game applying a guess
	//
	// in: string
	// example: "5d96bca0-2cf6-11ee-be56-0242ac120002"
	ID string `json:"id"`

	// Current Word current state
	//
	// in: string
	// example: "__PP__"
	Current string `json:"current"`

	// GuessesRemaining Amount of guesses remaining
	//
	// in: int64
	// example: 4
	GuessesRemaining int `json:"guesses_remaining"`
}

func NewGameResponseFromModel(g *models.Game) *GameResponse {
	return &GameResponse{
		ID:               g.ID,
		Current:          g.Current,
		GuessesRemaining: g.GuessesRemaining,
	}
}

func NewGameResponseFromStore(g *store.Game) *GameResponse {
	return &GameResponse{
		ID:               g.ID,
		Current:          g.Current,
		GuessesRemaining: g.GuessesRemaining,
	}
}
