package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/internal/location"
)

func TestParseConwayCube(t *testing.T) {
	t.Skip()
	given := strings.NewReader(`.#.
..#
###`)

	cells := []CubeCell{
		{location.Point{X: 0, Y: 0}, Inactive},
		{location.Point{X: 1, Y: 0}, Active},
		{location.Point{X: 2, Y: 0}, Inactive},
		{location.Point{X: 0, Y: 1}, Inactive},
		{location.Point{X: 1, Y: 1}, Inactive},
		{location.Point{X: 2, Y: 1}, Active},
		{location.Point{X: 0, Y: 2}, Active},
		{location.Point{X: 1, Y: 2}, Active},
		{location.Point{X: 2, Y: 2}, Active},
	}

	expected := Cube{
		CellsByLocation: map[location.Point]CubeCell{},
	}

	for _, cell := range cells {
		func(c CubeCell) {
			expected.CellsByLocation[c.Coordinate] = c
		}(cell)
	}

	actual := ParseConwayCube(given)

	assert.Equal(t, expected, actual)
}

func TestFindNeighboringLocations(t *testing.T) {
	t.Skip()
	cell := CubeCell{location.Point{X: 0, Y: 0, Z: 0}, Inactive}

	expectedNeighbors := []location.Point{
		{X: -1, Y: -1, Z: -1},
		{X: 0, Y: -1, Z: -1},
		{X: 1, Y: -1, Z: -1},
		{X: -1, Y: 0, Z: -1},
		{X: 0, Y: 0, Z: -1},
		{X: 1, Y: 0, Z: -1},
		{X: -1, Y: 1, Z: -1},
		{X: 0, Y: 1, Z: -1},
		{X: 1, Y: 1, Z: -1},
		{X: -1, Y: -1, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: 1, Y: -1, Z: 0},
		{X: -1, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 0},
		{X: -1, Y: 1, Z: 0},
		{X: 0, Y: 1, Z: 0},
		{X: 1, Y: 1, Z: 0},
		{X: -1, Y: -1, Z: 1},
		{X: 0, Y: -1, Z: 1},
		{X: 1, Y: -1, Z: 1},
		{X: -1, Y: 0, Z: 1},
		{X: 0, Y: 0, Z: 1},
		{X: 1, Y: 0, Z: 1},
		{X: -1, Y: 1, Z: 1},
		{X: 0, Y: 1, Z: 1},
		{X: 1, Y: 1, Z: 1},
	}

	actualNeighbors := cell.NeighborLocations()
	assert.ElementsMatch(t, expectedNeighbors, actualNeighbors)
}

func TestRunCubeCycle(t *testing.T) {
	t.Skip()
	cells := []CubeCell{
		{location.Point{X: 0, Y: 0}, Inactive},
		{location.Point{X: 1, Y: 0}, Active},
		{location.Point{X: 2, Y: 0}, Inactive},
		{location.Point{X: 0, Y: 1}, Inactive},
		{location.Point{X: 1, Y: 1}, Inactive},
		{location.Point{X: 2, Y: 1}, Active},
		{location.Point{X: 0, Y: 2}, Active},
		{location.Point{X: 1, Y: 2}, Active},
		{location.Point{X: 2, Y: 2}, Active},
	}

	cube := Cube{
		CellsByLocation: map[location.Point]CubeCell{},
	}

	for _, cell := range cells {
		func(c CubeCell) {
			cube.CellsByLocation[c.Coordinate] = c
		}(cell)
	}

	cube.RunCycle()

	expectedNumActiveCells := 11

	actualNumActiveCells := cube.CellCountByStatus(Active)

	assert.Equal(t, expectedNumActiveCells, actualNumActiveCells)
}
