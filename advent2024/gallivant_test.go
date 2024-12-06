package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
	"github.com/willmadison/advent/internal/location"
)

func TestParseMap(t *testing.T) {
	given := strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

	obstacles, guard, dimensions, err := advent2024.ParseMap(given)

	expectedOrientation := location.North
	expectedLocation := location.Coordinate{Row: 6, Col: 4}
	expectedDimensions := location.Coordinate{Row: 10, Col: 10}

	assert.Nil(t, err)
	assert.Equal(t, 8, len(obstacles))
	assert.Equal(t, expectedLocation, guard.Location)
	assert.Equal(t, expectedOrientation, guard.Orientation)
	assert.Equal(t, expectedDimensions, dimensions)
}

func TestSimulatePatrol(t *testing.T) {
	given := strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

	obstacles, guard, dimensions, err := advent2024.ParseMap(given)

	assert.Nil(t, err)

	patrolledLocations, err := advent2024.SimulatePatrol(guard, obstacles, dimensions)

	assert.Nil(t, err)

	expectedNumberOfLocations := 41

	assert.Equal(t, expectedNumberOfLocations, len(patrolledLocations))
}

func TestIntroduceObstacles(t *testing.T) {
	given := strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

	obstacles, guard, dimensions, err := advent2024.ParseMap(given)

	assert.Nil(t, err)

	expectedNumberOfLoops := 6

	actualNumberOfLoops := advent2024.IntroduceObstacles(guard, obstacles, dimensions)

	assert.Equal(t, expectedNumberOfLoops, actualNumberOfLoops)
}
