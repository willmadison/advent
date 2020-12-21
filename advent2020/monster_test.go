package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRulesAndMessages(t *testing.T) {
	given := strings.NewReader(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)

	expectedRules := Ruleset(map[int]Rule{
		0: {
			ID: 0,
			Subrules: [][]int{
				{4, 1, 5},
			},
		},
		1: {
			ID: 1,
			Subrules: [][]int{
				{2, 3},
				{3, 2},
			},
		},
		2: {
			ID: 2,
			Subrules: [][]int{
				{4, 4},
				{5, 5},
			},
		},
		3: {
			ID: 3,
			Subrules: [][]int{
				{4, 5},
				{5, 4},
			},
		},
		4: {
			ID:       4,
			Letter:   "a",
			Subrules: [][]int{},
		},
		5: {
			ID:       5,
			Letter:   "b",
			Subrules: [][]int{},
		},
	})

	expectedMessages := []string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}

	actualRules, actualMessages := ParseRulesAndMessages(given)

	assert.Equal(t, expectedRules, actualRules)
	assert.Equal(t, expectedMessages, actualMessages)
}

func TestCountMatches(t *testing.T) {
	ruleset := Ruleset(map[int]Rule{
		0: {
			ID: 0,
			Subrules: [][]int{
				{4, 1, 5},
			},
		},
		1: {
			ID: 1,
			Subrules: [][]int{
				{2, 3},
				{3, 2},
			},
		},
		2: {
			ID: 2,
			Subrules: [][]int{
				{4, 4},
				{5, 5},
			},
		},
		3: {
			ID: 3,
			Subrules: [][]int{
				{4, 5},
				{5, 4},
			},
		},
		4: {
			ID:       4,
			Letter:   "a",
			Subrules: [][]int{},
		},
		5: {
			ID:       5,
			Letter:   "b",
			Subrules: [][]int{},
		},
	})

	messages := []string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}

	expectedMatches := 2
	actualMatches := ruleset.FindMatches(0, messages)

	assert.Equal(t, expectedMatches, actualMatches)
}
