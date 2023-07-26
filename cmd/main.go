//	 WordGame API:
//	  version: v1.0
//	  title: WordGame API
//	 Schemes: http, https
//	 Host: localhost:1337
//	 BasePath: /
//		Consumes:
//		- application/json
//	 Produces:
//	 - application/json
//	 swagger:meta
package main

import (
	"github.com/fleetdm/wordgame/cmd/config"
	"github.com/fleetdm/wordgame/internal/api"
	"github.com/fleetdm/wordgame/internal/game"
	"github.com/fleetdm/wordgame/internal/server"
	"github.com/fleetdm/wordgame/internal/store/redis"
	"github.com/fleetdm/wordgame/pkg/words"
	"log"
)

func main() {

	// Load configurations from the environment yml file
	// We could load this from the environment variables and make it more dynamic
	cfg, err := config.LoadConfigs("config.local.yml")
	if err != nil {
		log.Fatal("Failed to load app configurations", err)
	}

	// Load words from file
	// This is not the most efficient way to load a word dataset, assuming that
	// this file could be very large, and we are keeping it in memory.
	// If this is not an explicit requirement, I would definitely consider to load this into redis
	// and get a word when it is needed.
	wordList, err := words.LoadWords("assets/words.txt")
	if err != nil {
		log.Fatal("Failed to load words file", err)
	}

	// Initialize Redis Store
	gameStore := redis.NewRedisStore(cfg.Redis.Address, cfg.Redis.Password, cfg.Redis.DB)

	// Initialize game service with the Redis store and the word list
	gameService := game.NewService(wordList, gameStore)

	// Create new Handler with the game service
	handler := api.NewHandler(gameService)

	// Initialize and start the server
	srv := server.NewServer(handler)

	// Setup routes	for this API
	srv.SetupRoutes()

	// Start the server
	if err := srv.Serve(cfg.Server.Address); err != nil {
		log.Fatal("Failed to start the serve", err)
	}
}
