package words_test

import (
	"github.com/fleetdm/wordgame/pkg/words"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadWords(t *testing.T) {
	// Determine the current working directory of this test file.
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// Append the relative path to the assets/words.txt file.
	path := filepath.Join(dir, "..", "..", "assets", "words.txt")

	words, err := words.LoadWords(path)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(words) == 0 {
		t.Errorf("No words were loaded")
	}

	// Write more tests based on your specific requirements.
}

func TestLoadWordsNonexistentFile(t *testing.T) {
	_, err := words.LoadWords("nonexistent.txt")

	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
