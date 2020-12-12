package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/internal/location"
)

func TestParseSeatingArrangment(t *testing.T) {
	given := strings.NewReader(`L.LL.LL.LL
L.L.L.L.LL`)

	expected := SeatingArrangement([]Seat{
		{
			location.Coordinate{0, 0},
			Empty,
		},
		{
			location.Coordinate{0, 2},
			Empty,
		},
		{
			location.Coordinate{0, 3},
			Empty,
		},
		{
			location.Coordinate{0, 5},
			Empty,
		},
		{
			location.Coordinate{0, 6},
			Empty,
		},
		{
			location.Coordinate{0, 8},
			Empty,
		},
		{
			location.Coordinate{0, 9},
			Empty,
		},
		{
			location.Coordinate{1, 0},
			Empty,
		},
		{
			location.Coordinate{1, 2},
			Empty,
		},
		{
			location.Coordinate{1, 4},
			Empty,
		},
		{
			location.Coordinate{1, 6},
			Empty,
		},
		{
			location.Coordinate{1, 8},
			Empty,
		},
		{
			location.Coordinate{1, 9},
			Empty,
		},
	})

	actual := ParseSeatingArrangement(given)
	assert.ElementsMatch(t, expected, actual)
}

func TestSeatingCycle(t *testing.T) {
	given := SeatingArrangement([]Seat{
		{
			location.Coordinate{0, 0},
			Empty,
		},
		{
			location.Coordinate{0, 2},
			Empty,
		},
		{
			location.Coordinate{0, 3},
			Empty,
		},
		{
			location.Coordinate{0, 5},
			Empty,
		},
		{
			location.Coordinate{0, 6},
			Empty,
		},
		{
			location.Coordinate{0, 8},
			Empty,
		},
		{
			location.Coordinate{0, 9},
			Empty,
		},
		{
			location.Coordinate{1, 0},
			Empty,
		},
		{
			location.Coordinate{1, 2},
			Empty,
		},
		{
			location.Coordinate{1, 4},
			Empty,
		},
		{
			location.Coordinate{1, 6},
			Empty,
		},
		{
			location.Coordinate{1, 8},
			Empty,
		},
		{
			location.Coordinate{1, 9},
			Empty,
		},
	})

	expected := SeatingArrangement([]Seat{
		{
			location.Coordinate{0, 0},
			Occupied,
		},
		{
			location.Coordinate{0, 2},
			Occupied,
		},
		{
			location.Coordinate{0, 3},
			Occupied,
		},
		{
			location.Coordinate{0, 5},
			Occupied,
		},
		{
			location.Coordinate{0, 6},
			Occupied,
		},
		{
			location.Coordinate{0, 8},
			Occupied,
		},
		{
			location.Coordinate{0, 9},
			Occupied,
		},
		{
			location.Coordinate{1, 0},
			Occupied,
		},
		{
			location.Coordinate{1, 2},
			Occupied,
		},
		{
			location.Coordinate{1, 4},
			Occupied,
		},
		{
			location.Coordinate{1, 6},
			Occupied,
		},
		{
			location.Coordinate{1, 8},
			Occupied,
		},
		{
			location.Coordinate{1, 9},
			Occupied,
		},
	})

	actual := given.RunSeatingCycle(4)
	assert.ElementsMatch(t, expected, actual)

	last := actual
	actual = given.RunSeatingCycle(4)

	assert.True(t, last.Equals(actual))
	assert.Equal(t, 13, len(actual.SeatsByState(Occupied)))
}

func TestExampleInput(t *testing.T) {
	arrangement := ParseSeatingArrangement(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))

	var lastSeatingArrangement SeatingArrangement

	for !lastSeatingArrangement.Equals(arrangement) {
		arrangement, lastSeatingArrangement = arrangement.RunSeatingCycle(5, FirstVisible), arrangement
	}

	assert.Equal(t, 26, len(arrangement.SeatsByState(Occupied)))
}
