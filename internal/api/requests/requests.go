package requests

// GuessRequest contains the ID of the game and the guessed letter
// swagger:model guessRequest
type GuessRequest struct {
	// The ID of the Game applying a guess
	//
	// in: string
	// required: true
	// example: "5d96bca0-2cf6-11ee-be56-0242ac120002"
	ID string `json:"id"`

	// The guessing letter
	//
	// in: string
	// required: true
	// example: "A"
	Guess string `json:"guess"`
}
