package api_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fleetdm/wordgame/internal/api"
	"github.com/fleetdm/wordgame/internal/game/mocks"
	"github.com/fleetdm/wordgame/internal/models"
)

// TestNewGameHandler tests the NewGameHandler function
func TestNewGameHandler(t *testing.T) {
	t.Run("returns 200 and game data when new game created successfully", func(t *testing.T) {
		mockGameService := &mocks.MockGameService{}
		mockGameService.On("NewGame").Return(&models.Game{
			ID:               "test_id",
			Current:          "____",
			GuessesRemaining: 6,
		}, nil)

		req, _ := http.NewRequest(http.MethodPost, "/new", nil)
		rec := httptest.NewRecorder()
		handler := api.NewHandler(mockGameService)

		handler.NewGameHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v; got %v", http.StatusOK, rec.Code)
		}

		expectedBody := `{"id":"test_id","current":"____","guesses_remaining":6}`
		assert.JSONEq(t, expectedBody, rec.Body.String())

	})

	t.Run("returns 500 when new game creation fails", func(t *testing.T) {
		mockGameService := &mocks.MockGameService{}
		mockGameService.On("NewGame").Return(&models.Game{}, errors.New("creation error"))

		req, _ := http.NewRequest(http.MethodPost, "/new", nil)
		rec := httptest.NewRecorder()
		handler := api.NewHandler(mockGameService)

		handler.NewGameHandler(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected %v; got %v", http.StatusInternalServerError, rec.Code)
		}
	})
}

// TestGuessHandler tests the GuessHandler function
func TestGuessHandler(t *testing.T) {
	t.Run("returns 200 and game data when guess is successful", func(t *testing.T) {
		mockGameService := &mocks.MockGameService{}
		mockGameService.On("Guess", "test_id", 'w').Return(nil)
		mockGameService.On("LoadGame", "test_id").Return(&models.Game{}, nil)

		req, _ := http.NewRequest(http.MethodPost, "/guess", strings.NewReader(`{"id":"test_id","guess":"w"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		handler := api.NewHandler(mockGameService)

		handler.GuessHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v; got %v", http.StatusOK, rec.Code)
		}
	})

	t.Run("returns 500 when guess fails", func(t *testing.T) {
		mockGameService := &mocks.MockGameService{}
		mockGameService.On("Guess", "test_id", 'w').Return(errors.New("guess error"))

		mockGameService.On("LoadGame", "test_id").Return(nil, errors.New("load game error"))

		req, _ := http.NewRequest(http.MethodPost, "/guess", strings.NewReader(`{"id":"test_id","guess":"w"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		handler := api.NewHandler(mockGameService)

		handler.GuessHandler(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected %v; got %v", http.StatusInternalServerError, rec.Code)
		}
	})
}
