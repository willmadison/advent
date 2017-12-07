package advent2017

import (
	"bufio"
	"sort"
	"strings"
)

type Passphrase string

func (p Passphrase) IsValid() bool {
	uniqueWords := map[string]struct{}{}

	var runeSortedWords []string

	allWords := strings.Fields(string(p))
	for _, word := range allWords {
		runes := []rune(word)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		runeSortedWords = append(runeSortedWords, string(runes))
	}

	for _, word := range runeSortedWords {
		uniqueWords[word] = struct{}{}
	}

	return len(uniqueWords) == len(allWords)
}

func CountValidPassphrases(passphrases string) int {
	var validPassphrases int

	scanner := bufio.NewScanner(strings.NewReader(passphrases))

	for scanner.Scan() {
		passphrase := Passphrase(scanner.Text())

		if passphrase.IsValid() {
			validPassphrases++
		}
	}

	return validPassphrases
}
