package store

import (
	"encoding/json"
	"github.com/fleetdm/wordgame/internal/models"
)

// Game is a placeholder for the game data.
type Game struct {
	ID               string        `json:"id"`                // Identifier for this game. Use this to make guesses in the game.
	Word             string        `json:"word"`              // Word to be guessed in the game.
	Current          string        `json:"current"`           // The current board state. Always consists of only _ characters at start of game.
	Status           models.Status `json:"status"`            // Status of the game. One of: "IN_PROGRESS", "WON", "LOST".
	GuessesRemaining int           `json:"guesses_remaining"` // How many guesses remain before the player loses.
}

// GameStore is an interface describing the database operations for this application.
// The idea of this interface is to allow for different implementations of the database.
type GameStore interface {
	// SaveGame creates a new game.
	SaveGame(game *Game) error
	// LoadGame retrieves a game by its ID.
	LoadGame(id string) (*Game, error)
}

// MarshalBinary converts the Game object to a byte slice.
func (g *Game) MarshalBinary() (data []byte, err error) {
	return json.Marshal(g)
}

// UnmarshalBinary converts a byte slice to a Game object.
func (g *Game) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, g)
}

// GameModelFromStore converts a Game object from the store to a Game object for the API.
func GameModelFromStore(g *Game) *models.Game {
	return &models.Game{
		ID:               g.ID,
		Current:          g.Current,
		GuessesRemaining: g.GuessesRemaining,
	}
}
