package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSchedule(t *testing.T) {
	given := strings.NewReader(`939
7,13,x,x,59,x,31,19`)

	expected := Itinerary{
		EarliestDeparture: 939,
		Busses: []Bus{
			{7, 0},
			{13, 1},
			{59, 4},
			{31, 6},
			{19, 7},
		},
	}

	actual := ParseShuttleItinerary(given)

	assert.Equal(t, expected, actual)
}

func TestFindEarliestBus(t *testing.T) {
	given := Itinerary{
		EarliestDeparture: 939,
		Busses: []Bus{
			{7, 0},
			{13, 1},
			{59, 4},
			{31, 6},
			{19, 7},
		},
	}

	expectedBus := 59
	expectedWaitTime := 5

	actualBus, actualWaitTime := FindEarliestBus(given)

	assert.Equal(t, expectedBus, actualBus)
	assert.Equal(t, expectedWaitTime, actualWaitTime)
}

func TestFindEarliestTimestampWithDepartureCadence(t *testing.T) {
	given := Itinerary{
		EarliestDeparture: 939,
		Busses: []Bus{
			{7, 0},
			{13, 1},
			{59, 4},
			{31, 6},
			{19, 7},
		},
	}

	expected := 1068781
	actual := FindEarliestTimestampWithDepartureCadence(given)

	assert.Equal(t, expected, actual)
}
