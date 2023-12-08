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

func TestNewHand(t *testing.T) {
	cards := []advent2023.Card{advent2023.Joker, advent2023.King, advent2023.Queen, advent2023.King, advent2023.King}

	hand := advent2023.NewHand(cards, advent2023.JokersWild)

	assert.Equal(t, advent2023.FourOfAKind, hand.Type)
}
func TestSortCamelCardHandWagersBonusInput(t *testing.T) {
	given := strings.NewReader(`2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JKQKK 21
JJJJ2 41`)

	wagers := advent2023.SortCamelCardWagers(given, advent2023.JokersWild)

	/*
		for _, wager := range wagers {
			fmt.Println(wager.String())
		}
	*/
	expectedTotalWinnings := 7460

	totalWinnings := 0

	for rank, wager := range wagers {
		totalWinnings += (rank + 1) * wager.Bid
	}

	assert.Equal(t, expectedTotalWinnings, totalWinnings)
}
