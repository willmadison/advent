package advent2018_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2018"
)

func TestFindLargestFiniteArea(t *testing.T) {
	cases := []struct {
		name     string
		given    io.Reader
		expected int
	}{
		{
			"Basic case...",
			strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`),
			17,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2018.FindLargestFiniteArea(tc.given)
			assert.Equal(t, tc.expected, actual)
		})

	}
}

func TestFindRegionAreaMinimizedByConstraint(t *testing.T) {
	cases := []struct {
		name  string
		given struct {
			input      io.Reader
			constraint int
		}
		expected int
	}{
		{
			"Basic case...",
			struct {
				input      io.Reader
				constraint int
			}{
				strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`),
				32,
			},
			16,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2018.FindRegionAreaMinimizedByConstraint(tc.given.input, tc.given.constraint)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
