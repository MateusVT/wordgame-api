package api

import (
	"encoding/json"
	"github.com/fleetdm/wordgame/internal/api/requests"
	"github.com/fleetdm/wordgame/internal/api/responses"
	"github.com/fleetdm/wordgame/internal/game"
	"log"
	"net/http"
)

// Handler handles API requests
type Handler struct {
	service game.ServiceInterface
}

// NewHandler creates a new API handler
func NewHandler(s game.ServiceInterface) *Handler {
	return &Handler{service: s}
}

// NewGameHandler handles the POST /new game
// swagger:operation POST /new newGame
// Start a new guessing game
//
// Allows users to start a new guessing game.
// ---
// produces:
// - application/json
// responses:
//
//	'200':
//	  description: Data of the initialized game
//	  schema:
//	    $ref: '#/responses/gameResponse'
//	'500':
//	  description: Internal server error
func (h *Handler) NewGameHandler(w http.ResponseWriter, _ *http.Request) {

	// Call the service to create a new game
	newGame, err := h.service.NewGame()
	if err != nil {
		log.Printf("Error creating new game: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := responses.NewGameResponseFromModel(newGame)
	log.Printf("New game sucessfully created! id = %s", response.ID)
	json.NewEncoder(w).Encode(response)
}

// GuessHandler handles the POST /guess endpoint
// swagger:operation POST /guess makesGuess
// Make a guess in a game
//
// Allows users to make a guess in a game by providing an ID and a guessed letter.
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
//   - in: body
//     name: body
//     description: ID of the game and the guessed letter
//     required: true
//     schema:
//     "$ref": "#/definitions/guessRequest"
//
// responses:
//
//	'200':
//	  description: Game object
//	  schema:
//	    $ref: '#/definitions/gameResponse'
//	'400':
//	  description: Invalid request body
//	'500':
//	  description: Internal server error
func (h *Handler) GuessHandler(w http.ResponseWriter, r *http.Request) {
	var guessRequest requests.GuessRequest

	if err := json.NewDecoder(r.Body).Decode(&guessRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(guessRequest.Guess) != 1 {
		http.Error(w, "Invalid request body, guess should be a single character", http.StatusBadRequest)
		return
	}

	err := h.service.Guess(guessRequest.ID, rune(guessRequest.Guess[0]))
	if err != nil {
		log.Printf("Error making guess, error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game, err := h.service.LoadGame(guessRequest.ID)
	if err != nil {
		log.Printf("Error getting game: %s, error: %s", guessRequest.ID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := responses.NewGameResponseFromModel(game)
	log.Printf("Guess '%s' was made in game %s!", guessRequest.Guess, guessRequest.ID)
	json.NewEncoder(w).Encode(response)
}
