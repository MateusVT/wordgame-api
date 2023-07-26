package identifier_test

import (
	"github.com/fleetdm/wordgame/pkg/identifier"
	"testing"
)

// TestGenerateIdentifier checks if generateIdentifier function
// is generating a non-empty identifier and the identifiers are unique
func TestGenerateIdentifier(t *testing.T) {
	// Generate two identifiers
	id1, err1 := identifier.GenerateIdentifier()
	id2, err2 := identifier.GenerateIdentifier()

	// Both identifiers should not generate an error
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if err2 != nil {
		t.Errorf("Unexpected error: %v", err2)
	}

	// Both identifiers should not be empty
	if id1 == "" {
		t.Errorf("Expected a non-empty string but got empty string")
	}
	if id2 == "" {
		t.Errorf("Expected a non-empty string but got empty string")
	}

	// Both identifiers should be unique
	if id1 == id2 {
		t.Errorf("Expected unique identifiers but got identical identifiers")
	}
}
