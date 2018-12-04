package advent2018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseClaim(t *testing.T) {
	cases := []struct {
		given    string
		expected Claim
	}{
		{
			"#1 @ 1,3: 4x4",
			Claim{
				ID:       1,
				Location: Point{3, 1},
				Width:    4,
				Height:   4,
			},
		},
		{
			"#2 @ 3,1: 4x4",
			Claim{
				ID:       2,
				Location: Point{1, 3},
				Width:    4,
				Height:   4,
			},
		},
		{
			"#3 @ 5,5: 2x2",
			Claim{
				ID:       3,
				Location: Point{5, 5},
				Width:    2,
				Height:   2,
			},
		},
	}

	for _, tc := range cases {
		actual := ParseClaim(tc.given)
		assert.Equal(t, tc.expected.ID, actual.ID)
		assert.Equal(t, tc.expected.Location, actual.Location)
		assert.Equal(t, tc.expected.Width, actual.Width)
		assert.Equal(t, tc.expected.Height, actual.Height)
	}
}

func TestFindOverlappingArea(t *testing.T) {
	cases := []struct {
		given    Claimset
		expected int
	}{
		{
			NewClaimset(ParseClaim("#1 @ 1,3: 4x4"),
				ParseClaim("#2 @ 3,1: 4x4"),
				ParseClaim("#3 @ 5,5: 2x2")),
			4,
		},
	}

	for _, tc := range cases {
		actual := tc.given.OverlappingRegion()
		assert.Equal(t, tc.expected, actual)
	}
}

func TestFindNonOverlappingClaim(t *testing.T) {
	cases := []struct {
		given    Claimset
		expected int
	}{
		{
			NewClaimset(ParseClaim("#1 @ 1,3: 4x4"),
				ParseClaim("#2 @ 3,1: 4x4"),
				ParseClaim("#3 @ 5,5: 2x2")),
			3,
		},
	}

	for _, tc := range cases {
		actual := tc.given.FindNonOverlappingClaim()
		assert.Equal(t, tc.expected, actual)
	}
}
