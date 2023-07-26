package mocks

import (
	"github.com/fleetdm/wordgame/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockGameService is a mock implementation of the GameService interface
type MockGameService struct {
	mock.Mock
}

func (m *MockGameService) NewGame() (*models.Game, error) {
	args := m.Called()
	return args.Get(0).(*models.Game), args.Error(1)
}

func (m *MockGameService) Guess(id string, letter rune) error {
	args := m.Called(id, letter)
	return args.Error(0)
}

func (m *MockGameService) LoadGame(id string) (*models.Game, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Game), args.Error(1)
}
