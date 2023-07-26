package server_test

import (
	"bytes"
	"github.com/fleetdm/wordgame/internal/api"
	"github.com/fleetdm/wordgame/internal/game/mocks"
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/fleetdm/wordgame/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestServer tests the server
func TestServer(t *testing.T) {
	mockGameService := &mocks.MockGameService{}
	mockGameService.On("NewGame").Return(&models.Game{
		ID:               "test_id",
		Current:          "____",
		GuessesRemaining: 6,
	}, nil)

	handler := api.NewHandler(mockGameService)
	srv := server.NewServer(handler)

	// Set up the test server for new game handler
	server := httptest.NewServer(http.HandlerFunc(srv.H.NewGameHandler))
	defer server.Close()

	t.Run("test new game handler", func(t *testing.T) {
		resp, err := http.Post(server.URL, "application/json", nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode, "they should be equal")

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read body: %v", err)
		}

		expectedBody := `{"id":"test_id","current":"____","guesses_remaining":6}`
		assert.JSONEq(t, expectedBody, string(body))

	})

	// Set up the test server for guess handler
	server = httptest.NewServer(http.HandlerFunc(srv.H.GuessHandler))
	defer server.Close()

	t.Run("test guess handler", func(t *testing.T) {

		mockGameService.On("Guess", "test_id", 'w').Return(nil)
		mockGameService.On("LoadGame", "test_id").Return(&models.Game{
			ID:               "test_id",
			Current:          "____",
			GuessesRemaining: 6,
		}, nil)

		// Convert the string body to a bytes.Buffer
		body := bytes.NewBuffer([]byte(`{"id":"test_id","guess":"w"}`))

		resp, err := http.Post(server.URL, "application/json", body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, resp.StatusCode, "they should be equal")

		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read body: %v", err)
		}

		expectedBody := `{"id":"test_id","current":"____","guesses_remaining":6}`
		assert.JSONEq(t, expectedBody, string(respBody))
	})

}

// TestServe tests the Serve function
// This test is a bit more complicated because we need to start the server in a goroutine
func TestServe(t *testing.T) {
	//TODO This test can be a little bit tricky
	// because it's a blocking call that runs indefinitely until an error occurs
	// Maybe if we have a pipeline that runs tests the implementation for this could cause problems
}

// TestSetupRoutes tests the SetupRoutes function
func TestSetupRoutes(t *testing.T) {
	mockGameService := &mocks.MockGameService{}
	mockGameService.On("NewGame").Return(&models.Game{
		ID:               "test_id",
		Current:          "____",
		GuessesRemaining: 6,
	}, nil)
	mockGameService.On("Guess", mock.Anything, mock.Anything).Return(nil)
	mockGameService.On("LoadGame", "test_id").Return(&models.Game{
		ID:               "test_id",
		Current:          "____",
		GuessesRemaining: 6,
	}, nil)

	handler := api.NewHandler(mockGameService)
	srv := server.NewServer(handler)

	srv.SetupRoutes()

	t.Run("test new game route", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/new", nil)
		rr := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("test guess route", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/guess", io.NopCloser(strings.NewReader(`{"id":"test_id","guess":"a"}`)))
		rr := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

// TestHandlerMiddleware tests the HandlerMiddleware function
func TestHandlerMiddleware(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test handler"))
	})

	// Wrap testHandler with the middleware
	handler := server.HandlerMiddleware(testHandler)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/new", nil)

	handler.ServeHTTP(recorder, req)

	// Check if the Content-Type header was correctly set
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
}
