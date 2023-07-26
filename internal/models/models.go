package models

type Status string

// Game statuses
// I created this type to represent the status of a game,
//
//	in a second version of the game, we could retrieve the status of a game to the user
const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusWon        Status = "WON"
	StatusLost       Status = "LOST"
)

// Game represents a game of word guessing
type Game struct {
	ID               string
	Current          string
	GuessesRemaining int
}
