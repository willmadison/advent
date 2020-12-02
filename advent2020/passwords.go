package advent2020

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type PasswordDBEntry struct {
	Policy   PasswordPolicy
	Password string
}

type PolicyType int

const (
	_ PolicyType = iota
	Legacy
	Current
)

func (p PasswordDBEntry) IsValid(policyType ...PolicyType) bool {
	policyEnforced := Legacy

	if len(policyType) > 0 {
		policyEnforced = policyType[0]
	}

	switch policyEnforced {
	case Legacy:
		return p.isValidByCharacterCount()
	case Current:
		return p.isValidByLocale()
	default:
		return false
	}

}

func (p PasswordDBEntry) isValidByCharacterCount() bool {
	characterCounts := map[rune]int{}

	for _, c := range p.Password {
		characterCounts[rune(c)]++
	}

	counts := characterCounts[p.Policy.Character]

	if counts < p.Policy.MinOccurrences || counts > p.Policy.MaxOccurrences {
		return false
	}

	return true
}

func (p PasswordDBEntry) isValidByLocale() bool {
	locationsByCharacter := map[rune]map[int]struct{}{}

	for location, c := range p.Password {
		if _, present := locationsByCharacter[rune(c)]; !present {
			locationsByCharacter[rune(c)] = map[int]struct{}{}
		}

		locationsByCharacter[rune(c)][location+1] = struct{}{}
	}

	firstLocale := p.Policy.MinOccurrences
	secondLocale := p.Policy.MaxOccurrences

	_, inFirstLocale := locationsByCharacter[p.Policy.Character][firstLocale]
	_, inSecondLocale := locationsByCharacter[p.Policy.Character][secondLocale]

	return (inFirstLocale || inSecondLocale) && !(inFirstLocale && inSecondLocale)
}

type PasswordPolicy struct {
	Character                      rune
	MinOccurrences, MaxOccurrences int
}

func CountValidPasswords(r io.Reader) int {
	entries := parsePasswordDBEntries(r)

	var validEntries int

	for _, entry := range entries {
		if entry.IsValid() {
			validEntries++
		}
	}

	return validEntries
}

func CountValidPasswordsUpdatedPolicy(r io.Reader) int {
	entries := parsePasswordDBEntries(r)

	var validEntries int

	for _, entry := range entries {
		if entry.IsValid(Current) {
			validEntries++
		}
	}

	return validEntries
}

func parsePasswordDBEntries(r io.Reader) []PasswordDBEntry {
	var entries []PasswordDBEntry

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawEntry := scanner.Text()
		entry := parseEntry(rawEntry)
		entries = append(entries, entry)
	}

	return entries
}

var entryRegEx = regexp.MustCompile(`^(\d+)\-(\d+)\s(\S):\s(\S+)$`)

func parseEntry(rawEntry string) PasswordDBEntry {
	matches := entryRegEx.FindStringSubmatch(rawEntry)

	rawMinimum := matches[1]
	rawMaximum := matches[2]

	rawCharacter := matches[3]

	minimum, _ := strconv.Atoi(rawMinimum)
	maximum, _ := strconv.Atoi(rawMaximum)

	policy := PasswordPolicy{rune(rawCharacter[0]), minimum, maximum}

	password := matches[4]

	return PasswordDBEntry{policy, password}
}
