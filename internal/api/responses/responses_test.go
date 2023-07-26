package responses_test

import (
	"github.com/fleetdm/wordgame/internal/api/responses"
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewGameResponseFromModel tests the NewGameResponseFromModel function.
func TestNewGameResponseFromModel(t *testing.T) {
	game := &models.Game{
		ID:               "1",
		Current:          "_e__",
		GuessesRemaining: 3,
	}

	response := responses.NewGameResponseFromModel(game)

	assert.Equal(t, game.ID, response.ID)
	assert.Equal(t, game.Current, response.Current)
	assert.Equal(t, game.GuessesRemaining, response.GuessesRemaining)
}

// TestNewGameResponseFromStore tests the NewGameResponseFromStore function.
func TestNewGameResponseFromStore(t *testing.T) {
	game := &store.Game{
		ID:               "1",
		Current:          "_e__",
		GuessesRemaining: 3,
	}

	response := responses.NewGameResponseFromStore(game)

	assert.Equal(t, game.ID, response.ID)
	assert.Equal(t, game.Current, response.Current)
	assert.Equal(t, game.GuessesRemaining, response.GuessesRemaining)
}
