package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestDistanceToFurtherPipeFromStart(t *testing.T) {
	given := strings.NewReader(`.....
.S-7.
.|.|.
.L-J.
.....`)

	expected := 4

	actual := advent2023.DistanceToFurthestPipeFromStart(given)

	assert.Equal(t, expected, actual)

	given = strings.NewReader(`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`)

	expected = 8

	actual = advent2023.DistanceToFurthestPipeFromStart(given)

	assert.Equal(t, expected, actual)
}
