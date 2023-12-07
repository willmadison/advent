package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestSortCamelCardHandWagers(t *testing.T) {
	given := strings.NewReader(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	expectedNumWagers := 5

	wagers := advent2023.SortCamelCardWagers(given)

	assert.Equal(t, expectedNumWagers, len(wagers))

	expectedTotalWinnings := 6440

	totalWinnings := 0

	for rank, wager := range wagers {
		totalWinnings += (rank + 1) * wager.Bid
	}

	assert.Equal(t, expectedTotalWinnings, totalWinnings)
}

func TestSortCamelCardHandWagersJokersWild(t *testing.T) {
	given := strings.NewReader(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	expectedNumWagers := 5

	wagers := advent2023.SortCamelCardWagers(given, advent2023.JokersWild)

	assert.Equal(t, expectedNumWagers, len(wagers))

	expectedTotalWinnings := 5905

	totalWinnings := 0

	for rank, wager := range wagers {
		totalWinnings += (rank + 1) * wager.Bid
	}

	assert.Equal(t, expectedTotalWinnings, totalWinnings)
}
