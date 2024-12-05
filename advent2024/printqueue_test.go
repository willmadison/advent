package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestParseRulesAndPages(t *testing.T) {
	given := strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	ruleset, pagelists, err := advent2024.ParseRulesAndPages(given)

	assert.Nil(t, err)

	expectedNumberOfPagelists := 6

	assert.Equal(t, expectedNumberOfPagelists, len(pagelists))
	assert.True(t, len(ruleset) <= 21)

	expectedPrintable := 3

	var actualPrintable int

	for _, pages := range pagelists {
		if advent2024.CanPrint(ruleset, pages) {
			actualPrintable++
		}
	}

	assert.Equal(t, expectedPrintable, actualPrintable)

	expectedSumOfMiddlePages := 143

	var actualSumOfMiddlePages int

	for _, pages := range pagelists {
		if advent2024.CanPrint(ruleset, pages) {
			actualSumOfMiddlePages += pages[len(pages)/2]
		}
	}

	assert.Equal(t, expectedSumOfMiddlePages, actualSumOfMiddlePages)
}

func TestPrintCorrection(t *testing.T) {
	given := strings.NewReader(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

	ruleset, pagelists, err := advent2024.ParseRulesAndPages(given)

	assert.Nil(t, err)

	invalidPrintRequest := pagelists[3]

	expectedCorrectListing := []int{97, 75, 47, 61, 53}

	actualCorrectedListing := advent2024.CorrectPrintRequest(ruleset, invalidPrintRequest)

	assert.Equal(t, expectedCorrectListing, actualCorrectedListing)
}
