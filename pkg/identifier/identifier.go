package identifier

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GenerateIdentifier generates a new random UUID used to identify an
// in-progress game.
func GenerateIdentifier() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "generate game ID")
	}

	return id.String(), nil
}
