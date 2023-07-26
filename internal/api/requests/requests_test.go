package requests_test

import (
	"testing"

	"github.com/fleetdm/wordgame/internal/api/requests"
	"github.com/stretchr/testify/assert"
)

// TestGuessRequest tests the GuessRequest struct
func TestGuessRequest(t *testing.T) {
	t.Run("creates a new guess request", func(t *testing.T) {
		request := requests.GuessRequest{
			ID:    "test_id",
			Guess: "A",
		}

		assert.Equal(t, "test_id", request.ID)
		assert.Equal(t, "A", request.Guess)
	})
}
