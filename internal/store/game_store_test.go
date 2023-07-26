package store_test

import (
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
	"github.com/fleetdm/wordgame/internal/store/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMockGameStore tests the mock implementation of the GameStore interface
func TestMockGameStore(t *testing.T) {
	mockStore := new(mocks.MockGameStore)
	gameStore := mockStore

	game := &store.Game{ID: "test", Word: "word", Current: "current", Status: models.StatusInProgress, GuessesRemaining: 3}

	// Mock the SaveGame method
	mockStore.On("SaveGame", game).Return(nil)
	err := gameStore.SaveGame(game)
	assert.NoError(t, err)
	mockStore.AssertCalled(t, "SaveGame", game)

	// Mock the LoadGame method
	mockStore.On("LoadGame", game.ID).Return(game, nil)
	gotGame, err := gameStore.LoadGame(game.ID)
	assert.NoError(t, err)
	assert.Equal(t, game, gotGame)
	mockStore.AssertCalled(t, "LoadGame", game.ID)
}

// TestGameStore_MarshalBinary tests the MarshalBinary method of the GameStore interface
func TestGameStore_MarshalBinary(t *testing.T) {
	// Initialize a new game
	game := store.Game{
		ID:               "1",
		Current:          "_",
		Status:           models.StatusInProgress,
		GuessesRemaining: 5,
	}

	// Marshal the game into a byte slice
	data, err := game.MarshalBinary()
	assert.NoError(t, err, "Expected MarshalBinary to complete without error")
	assert.NotNil(t, data, "Expected data to be not nil")

}

// TestGameStore_UnmarshalBinary tests the UnmarshalBinary method of the GameStore interface
func TestGameStore_UnmarshalBinary(t *testing.T) {
	// Initialize a JSON representation of a Game
	data := []byte(`{"id":"1","word":"test","status":"IN_PROGRESS","current":"_____","guesses_remaining":5}`)

	game := store.Game{}
	err := game.UnmarshalBinary(data)

	// The Game struct should now hold the unmarshalled data
	assert.NoError(t, err, "Expected UnmarshalBinary to complete without error")
	assert.Equal(t, "1", game.ID, "Expected ID to be '1'")
	assert.Equal(t, "test", game.Word, "Expected Word to be 'test'")
	assert.Equal(t, "_____", game.Current, "Expected Current to be '_____'")
	assert.Equal(t, models.StatusInProgress, game.Status, "Expected Current to be IN_PROGRESS")
	assert.Equal(t, 5, game.GuessesRemaining, "Expected GuessesRemaining to be 5")
}
