package models_test

import (
	"testing"

	"github.com/fleetdm/wordgame/internal/models"
)

func TestGameStatuses(t *testing.T) {
	statuses := []models.Status{models.StatusInProgress, models.StatusWon, models.StatusLost}

	for _, status := range statuses {
		if status != models.StatusInProgress && status != models.StatusWon && status != models.StatusLost {
			t.Errorf("unexpected status: got %v", status)
		}
	}
}

func TestGameInitialization(t *testing.T) {
	game := &models.Game{
		ID:               "test_id",
		Current:          "____",
		GuessesRemaining: 6,
	}

	if game.ID != "test_id" {
		t.Errorf("unexpected ID: got %v", game.ID)
	}

	if game.Current != "____" {
		t.Errorf("unexpected Current: got %v", game.Current)
	}

	if game.GuessesRemaining != 6 {
		t.Errorf("unexpected GuessesRemaining: got %v", game.GuessesRemaining)
	}
}
