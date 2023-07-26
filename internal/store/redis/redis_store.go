package redis

import (
	"context"
	"github.com/fleetdm/wordgame/internal/store"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// Store struct keeps the Redis client connection and the context.
type Store struct {
	Client *redis.Client
	Ctx    context.Context
}

// NewRedisStore initializes a new Redis store.
func NewRedisStore(addr string, password string, db int) store.GameStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	// New store with the initialized client and context.
	return &Store{
		Client: rdb,
		Ctx:    ctx,
	}
}

// SaveGame method saves a game instance into the Redis store.
func (rs *Store) SaveGame(game *store.Game) error {
	// Marshal the game into binary.
	gameBytes, err := game.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, "marshalling game failed")
	}

	// Persist the game instance into the Redis store.
	err = rs.Client.Set(rs.Ctx, game.ID, gameBytes, 0).Err()
	if err != nil {
		return errors.Wrap(err, "saving game to redis failed")
	}
	return nil
}

// LoadGame method loads a game instance from the Redis store.
func (rs *Store) LoadGame(gameID string) (*store.Game, error) {
	var game store.Game
	val, err := rs.Client.Get(rs.Ctx, gameID).Result()
	if err == redis.Nil {
		return nil, errors.New("game not found")
	} else if err != nil {
		return nil, errors.Wrap(err, "error retrieving game from redis")
	}

	// Unmarshal the retrieved game instance.
	err = game.UnmarshalBinary([]byte(val))
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling game data")
	}

	// Return the retrieved game instance.
	return &game, nil
}
