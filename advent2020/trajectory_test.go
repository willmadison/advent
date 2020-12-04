package advent2020

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/internal/location"
)

func TestParseTrajectoryMap(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected TrajectoryMap
	}{
		{
			strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`),
			TrajectoryMap{
				Trees: map[location.Coordinate]struct{}{
					{0, 2}:   {},
					{0, 3}:   {},
					{1, 0}:   {},
					{1, 4}:   {},
					{1, 4}:   {},
					{1, 8}:   {},
					{2, 1}:   {},
					{2, 6}:   {},
					{2, 9}:   {},
					{3, 2}:   {},
					{3, 4}:   {},
					{3, 8}:   {},
					{3, 10}:  {},
					{4, 1}:   {},
					{4, 5}:   {},
					{4, 6}:   {},
					{4, 9}:   {},
					{5, 2}:   {},
					{5, 4}:   {},
					{5, 5}:   {},
					{6, 1}:   {},
					{6, 3}:   {},
					{6, 5}:   {},
					{6, 10}:  {},
					{7, 1}:   {},
					{7, 10}:  {},
					{8, 0}:   {},
					{8, 2}:   {},
					{8, 3}:   {},
					{8, 7}:   {},
					{9, 0}:   {},
					{9, 4}:   {},
					{9, 5}:   {},
					{9, 10}:  {},
					{10, 1}:  {},
					{10, 4}:  {},
					{10, 8}:  {},
					{10, 10}: {},
				},
				Rows:    11,
				Columns: 11,
			},
		},
	}

	for _, tc := range cases {
		actual := NewTrajectoryMap(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCountEncounteredTrees(t *testing.T) {
	cases := []struct {
		given struct {
			trajectoryMap TrajectoryMap
			slope         location.Slope
		}
		expected int
	}{
		{
			struct {
				trajectoryMap TrajectoryMap
				slope         location.Slope
			}{
				trajectoryMap: TrajectoryMap{
					Trees: map[location.Coordinate]struct{}{
						{0, 2}:   {},
						{0, 3}:   {},
						{1, 0}:   {},
						{1, 4}:   {},
						{1, 4}:   {},
						{1, 8}:   {},
						{2, 1}:   {},
						{2, 6}:   {},
						{2, 9}:   {},
						{3, 2}:   {},
						{3, 4}:   {},
						{3, 8}:   {},
						{3, 10}:  {},
						{4, 1}:   {},
						{4, 5}:   {},
						{4, 6}:   {},
						{4, 9}:   {},
						{5, 2}:   {},
						{5, 4}:   {},
						{5, 5}:   {},
						{6, 1}:   {},
						{6, 3}:   {},
						{6, 5}:   {},
						{6, 10}:  {},
						{7, 1}:   {},
						{7, 10}:  {},
						{8, 0}:   {},
						{8, 2}:   {},
						{8, 3}:   {},
						{8, 7}:   {},
						{9, 0}:   {},
						{9, 4}:   {},
						{9, 5}:   {},
						{9, 10}:  {},
						{10, 1}:  {},
						{10, 4}:  {},
						{10, 8}:  {},
						{10, 10}: {},
					},
					Rows:    11,
					Columns: 11,
				},
				slope: location.Slope{1, 3},
			},
			7,
		},
	}

	for _, tc := range cases {
		actual := CountEncounteredTrees(tc.given.trajectoryMap, tc.given.slope)
		assert.Equal(t, tc.expected, actual)
	}
}
