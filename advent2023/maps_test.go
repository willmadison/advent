package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestDetermineTripLength(t *testing.T) {
	given := strings.NewReader(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`)

	expected := 2

	actual := advent2023.DetermineTripLength(given)

	assert.Equal(t, expected, actual)

	given = strings.NewReader(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)

	expected = 6

	actual = advent2023.DetermineTripLength(given)

	assert.Equal(t, expected, actual)
}

func TestDetermineGhostlyTripLength(t *testing.T) {
	given := strings.NewReader(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)

	expected := 6

	actual := advent2023.DetermineGhostlyTripLength(given)

	assert.Equal(t, expected, actual)
}
