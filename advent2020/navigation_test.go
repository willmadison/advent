package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/internal/location"
)

func TestParseNavigationInstructions(t *testing.T) {
	given := strings.NewReader(`F10
N3
F7
R90
F11`)

	expected := []NavigationInstruction{
		{Forward, 10},
		{North, 3},
		{Forward, 7},
		{Right, 90},
		{Forward, 11},
	}

	actual := ParseNavigationInstructions(given)

	assert.Equal(t, expected, actual)
}

func TestNavigate(t *testing.T) {
	var ship Ship

	instructions := []NavigationInstruction{
		{Forward, 10},
		{North, 3},
		{Forward, 7},
		{Right, 90},
		{Forward, 11},
	}

	ship.Navigate(instructions)

	var origin location.Point

	assert.Equal(t, 25, origin.ManhattanDistance(ship.Location))

}

func TestNavigateWithWaypoint(t *testing.T) {
	var ship Ship

	instructions := []NavigationInstruction{
		{Forward, 10},
		{North, 3},
		{Forward, 7},
		{Right, 90},
		{Forward, 11},
	}

	ship.NavigateByWaypoint(instructions)

	var origin location.Point

	assert.Equal(t, 286, origin.ManhattanDistance(ship.Location))
}
