package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSeatLocation(t *testing.T) {
	cases := []struct{
		given string
		expectedRow, expectedColumn int
	}{
			{
				"FBFBBFFRLR",
				44,
				5,
			},
			{
				"BFFFBBFRRR",
				70,
				7,
			},
			{
				"FFFBBBFRRR",
				14,
				7,
			},
			{
				"BBFFBBFRLL",
				102,
				4,
			},
	}

	for _, tc := range cases {
		actualRow, actualColumn := FindSeatLocation(tc.given)
		assert.Equal(t, tc.expectedRow, actualRow)
		assert.Equal(t, tc.expectedColumn, actualColumn)
	}

}

func TestFindMaxSeatID(t *testing.T) {
	given := strings.NewReader(`FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`)
	expected := 820

	actual := FindMaxSeatID(given)
	assert.Equal(t, expected, actual)
}