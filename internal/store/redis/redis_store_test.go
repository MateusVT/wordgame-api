package redis_test

import (
	"context"
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
	"github.com/fleetdm/wordgame/internal/store/redis"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestStore_SaveGame tests the SaveGame method of the redis Store
func TestStore_SaveGame(t *testing.T) {
	db, mock := redismock.NewClientMock()

	gameToSave := &store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "____",
		GuessesRemaining: 6,
	}
	gameBytes, _ := gameToSave.MarshalBinary()

	mock.ExpectSet("test_id", gameBytes, 0).SetVal("OK")

	store := &redis.Store{
		Client: db,
		Ctx:    context.Background(),
	}

	err := store.SaveGame(gameToSave)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestStore_LoadGame tests the LoadGame method of the redis Store
func TestStore_LoadGame(t *testing.T) {
	db, mock := redismock.NewClientMock()

	gameToLoad := &store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "____",
		GuessesRemaining: 6,
	}
	gameBytes, _ := gameToLoad.MarshalBinary()

	mock.ExpectGet("test_id").SetVal(string(gameBytes))

	store := &redis.Store{
		Client: db,
		Ctx:    context.Background(),
	}

	game, err := store.LoadGame("test_id")

	assert.NoError(t, err)
	assert.Equal(t, "test_id", game.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
