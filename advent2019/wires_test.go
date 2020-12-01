package advent2019_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent"
	"github.com/willmadison/advent/advent2019"
)

func TestFindNearestIntersection(t *testing.T) {
	cases := []struct {
		name             string
		given            io.Reader
		expectedDistance int
	}{
		{
			"First test case...",
			strings.NewReader(`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R8312`),
			159,
		},
		{
			"Second test case...",
			strings.NewReader(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`),
			135,
		},
	}

	var origin advent.Point

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			intersection := advent2019.FindNearestIntersection(tc.given)
			assert.Equal(t, tc.expectedDistance, origin.ManhattanDistance(intersection))
		})

	}
}

func TestFindMinimalTotalSteps(t *testing.T) {
	cases := []struct {
		name               string
		given              io.Reader
		expectedTotalSteps int
	}{
		{
			"First test case...",
			strings.NewReader(`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R8312`),
			610,
		},
		{
			"Second test case...",
			strings.NewReader(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`),
			410,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			steps := advent2019.FindMinimalTotalSteps(tc.given)
			assert.Equal(t, tc.expectedTotalSteps, steps)
		})

	}
}
