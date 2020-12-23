package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDecks(t *testing.T) {
	given := strings.NewReader(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)

	expectedDeck1 := NewDeck(1, []Card{9, 2, 6, 3, 1})
	expectedDeck2 := NewDeck(2, []Card{5, 8, 4, 7, 10})

	actualDeck1, actualDeck2 := ParseDecks(given)

	assert.Equal(t, expectedDeck1, actualDeck1)
	assert.Equal(t, expectedDeck2, actualDeck2)
}

func TestBattle(t *testing.T) {
	player1 := NewDeck(1, []Card{9, 2, 6, 3, 1})
	player2 := NewDeck(2, []Card{5, 8, 4, 7, 10})

	winner := Battle(&player1, &player2)

	expectedWinningScore := 306

	assert.Equal(t, expectedWinningScore, winner.Score())
}

func TestSerialization(t *testing.T) {
	player1 := NewDeck(1, []Card{9, 2, 6, 3, 1})
	player2 := NewDeck(2, []Card{5, 8, 4, 7, 10})

	expected := "1:[9 2 6 3 1]"
	expectedPlayer2 := "2:[5 8 4 7 10]"

	assert.Equal(t, expected, player1.Serialize())

	expectedRoundSerialization := expected + " && " + expectedPlayer2

	assert.Equal(t, expectedRoundSerialization, SerializeRound(player1, player2))
}

func TestRecursiveBattle(t *testing.T) {
	player1 := NewDeck(1, []Card{9, 2, 6, 3, 1})
	player2 := NewDeck(2, []Card{5, 8, 4, 7, 10})

	winner := RecursiveBattle(&player1, &player2)

	expectedWinningScore := 291

	assert.Equal(t, expectedWinningScore, winner.Score())
}
