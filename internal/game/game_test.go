package game_test

import (
	"github.com/fleetdm/wordgame/internal/game"
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/store"
	"github.com/fleetdm/wordgame/internal/store/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestService_NewGame tests the NewGame method of the game service
func TestService_NewGame(t *testing.T) {
	mockStore := &mocks.MockGameStore{}
	mockStore.On("SaveGame", mock.Anything).Return(nil) // define the response of the SaveGame method

	s := game.NewService([]string{"word"}, mockStore)

	game, err := s.NewGame()

	assert.NoError(t, err)
	assert.Equal(t, "____", game.Current)
	assert.Equal(t, 6, game.GuessesRemaining)
}

// TestService_Guess tests the Guess method of the game service
// This method must be specially tested because it is the core operation of the game
// If this implementation grows, it would be a good idea to split this test into multiple tests
func TestService_Guess(t *testing.T) {
	// 1. Test game is lost after last incorrect guess
	mockStore := &mocks.MockGameStore{}
	mockStore.On("LoadGame", "test_id").Return(&store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "____",
		GuessesRemaining: 1, // Only one guess remaining
	}, nil)

	mockStore.On("SaveGame", mock.MatchedBy(func(game *store.Game) bool {
		return game.Status == models.StatusLost
	})).Return(nil)

	s := game.NewService([]string{"word"}, mockStore)
	err := s.Guess("test_id", 'Z') // Incorrect guess
	assert.NoError(t, err)

	// 2. Test game is won after correct guess reveal the entire word
	mockStore = &mocks.MockGameStore{}
	mockStore.On("LoadGame", "test_id").Return(&store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "WOR_",
		GuessesRemaining: 1,
	}, nil)

	mockStore.On("SaveGame", mock.MatchedBy(func(game *store.Game) bool {
		return game.Status == models.StatusWon
	})).Return(nil)

	s = game.NewService([]string{"word"}, mockStore)
	err = s.Guess("test_id", 'D') // Correct guess
	assert.NoError(t, err)

	// 3. Test game continues in progress after incorrect guess with remaining guesses
	mockStore = &mocks.MockGameStore{}
	mockStore.On("LoadGame", "test_id").Return(&store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "WOR_",
		GuessesRemaining: 2, // Two guesses remaining
	}, nil)

	mockStore.On("SaveGame", mock.MatchedBy(func(game *store.Game) bool {
		return game.Status == models.StatusInProgress
	})).Return(nil)

	s = game.NewService([]string{"word"}, mockStore)
	err = s.Guess("test_id", 'Z') // Incorrect guess
	assert.NoError(t, err)

	// 4. Test error in saving the game
	mockStore = &mocks.MockGameStore{}
	mockStore.On("LoadGame", "test_id").Return(&store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "WOR_",
		GuessesRemaining: 2,
	}, nil)

	mockStore.On("SaveGame", mock.Anything).Return(errors.New("save error"))

	s = game.NewService([]string{"word"}, mockStore)
	err = s.Guess("test_id", 'Z')
	assert.Error(t, err)
}

// TestService_GetGame tests the LoadGame method of the game service
func TestService_GetGame(t *testing.T) {
	mockStore := &mocks.MockGameStore{} // create an instance of the autogenerated mock
	mockStore.On("LoadGame", "test_id").Return(&store.Game{
		ID:               "test_id",
		Word:             "WORD",
		Status:           models.StatusInProgress,
		Current:          "____",
		GuessesRemaining: 6,
	}, nil) // define the response of the LoadGame method

	s := game.NewService([]string{"word"}, mockStore)

	game, err := s.LoadGame("test_id")

	assert.NoError(t, err)
	assert.Equal(t, "test_id", game.ID)
	assert.Equal(t, "____", game.Current)
	assert.Equal(t, 6, game.GuessesRemaining)
}
