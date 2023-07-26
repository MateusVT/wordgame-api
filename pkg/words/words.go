package words

import (
	"bufio"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Thanks to https://github.com/dwyl/english-words for the word list.

var wordsRegexp = regexp.MustCompile("^[A-Z]+$")

// LoadWords loads the word dictionary from the provided file path.
func LoadWords(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open word file")
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Normalize and filter words.
		word := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		if wordsRegexp.MatchString(word) {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "scan words")
	}

	return words, nil
}

// ChooseRandomWord pick a random word from the provided word list
func ChooseRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}
