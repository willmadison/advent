package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattanDistanceTo(t *testing.T) {
	cases := []struct {
		given    int
		expected int
	}{
		{
			1,
			0,
		},
		{
			12,
			3,
		},
		{
			23,
			2,
		},
		{
			1024,
			31,
		},
	}

	for _, c := range cases {
		actual := ManhattanDistanceTo(c.given)
		assert.Equal(t, c.expected, actual)
	}
}

func TestToCartesianCoordinate(t *testing.T) {
	cases := []struct {
		given    int
		expected Point
	}{
		{
			given:    1,
			expected: Point{0, 0},
		}, {
			given:    3,
			expected: Point{1, 1},
		}, {
			given:    13,
			expected: Point{2, 2},
		}, {
			given:    9,
			expected: Point{1, -1},
		}, {
			given:    25,
			expected: Point{2, -2},
		}, {
			given:    48,
			expected: Point{2, -3},
		},
	}

	for _, c := range cases {
		p := toCartesianCoordinate(c.given)
		assert.Equal(t, c.expected, p)
	}
}

func TestFirstCellLargerThan(t *testing.T) {
	cases := []struct {
		given    int
		expected int
	}{
		{
			1,
			2,
		},
		{
			9,
			10,
		},
		{
			60,
			122,
		},
		{
			50,
			54,
		},
		{
			600,
			747,
		},
	}

	for _, c := range cases {
		actual := FirstCellLargerThan(c.given)
		assert.Equal(t, c.expected, actual)
	}

}
