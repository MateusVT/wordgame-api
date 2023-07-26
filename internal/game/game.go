package game

import (
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
	"github.com/fleetdm/wordgame/pkg/identifier"
	"github.com/fleetdm/wordgame/pkg/words"
	"strings"
)

const (
	defaultGuessesLimit = 6   // Default amount of guesses allowed it could be configurable
	wordPlaceHolder     = "_" // Default placeholder for the hidden word it could be configurable
)

// ServiceInterface is an interface that any game service should satisfy.
// This is very useful for testing purposes.
type ServiceInterface interface {
	NewGame() (*models.Game, error)
	Guess(string, rune) error
	LoadGame(string) (*models.Game, error)
}

// Service struct which will contain the store and the word list
type Service struct {
	words []string
	store store.GameStore
}

// NewService initializes a new game service.
func NewService(words []string, store store.GameStore) ServiceInterface {
	return &Service{
		words: words,
		store: store,
	}
}

// NewGame starts a new game.
func (s *Service) NewGame() (*models.Game, error) {

	//Generate a unique identifier for the game
	id, err := identifier.GenerateIdentifier()
	if err != nil {
		return nil, err
	}

	// Choose a random word from the word list to be guessed
	word := words.ChooseRandomWord(s.words)
	// Initialize the hidden word placeholder state
	current := strings.Repeat(wordPlaceHolder, len(word))

	// Create a new game object to be stored
	newGame := &store.Game{
		ID:               id,
		Word:             word,
		Status:           models.StatusInProgress,
		Current:          current,
		GuessesRemaining: defaultGuessesLimit,
	}

	err = s.store.SaveGame(newGame)
	if err != nil {
		return nil, err
	}

	return store.GameModelFromStore(newGame), nil
}

// Guess execute a guess action in the game
// We could implement a status return to inform the user about the game status
func (s *Service) Guess(id string, letter rune) error {

	// Load game state from the store
	g, err := s.store.LoadGame(id)
	if err != nil {
		return err
	}

	// Check if game is already over and avoid any further action
	if g.Status != models.StatusInProgress {
		// We could implement a status return to inform the user about the game status
		// For now we just don't apply any action
		return nil
	}

	//Convert guess string to uppercase rune
	letter = rune(strings.ToUpper(string(letter))[0])

	guessWasCorrect := false
	// Iterate over the word to check where the letter is present
	for i, wordLetter := range g.Word {
		if wordLetter == letter {
			g.Current = g.Current[:i] + string(letter) + g.Current[i+1:]
			guessWasCorrect = true
		}
	}

	// Check if the guess was incorrect and there are remaining guesses
	// The second condition shouldn't be necessary, due the status control, but it's a safety check
	if !guessWasCorrect && g.GuessesRemaining > 0 {
		g.GuessesRemaining--
	}

	// Check if the remaining guesses are 0 which means that the game is lost
	if g.GuessesRemaining == 0 && g.Current != g.Word {
		g.Status = models.StatusLost
	}
	// Check if the current state is equal to the word which means that the game is won
	if g.Current == g.Word && g.GuessesRemaining > 0 {
		g.Status = models.StatusWon
	}
	// Update game state in the store
	err = s.store.SaveGame(g)
	if err != nil {
		return err
	}

	return nil
}

// LoadGame get a game given the id
func (s *Service) LoadGame(id string) (*models.Game, error) {
	g, err := s.store.LoadGame(id)

	if err != nil {
		return nil, err
	}
	return store.GameModelFromStore(g), nil
}
