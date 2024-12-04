package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestCountXmases(t *testing.T) {
	given := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 18

	count, err := advent2024.CountXmases(given)

	assert.Nil(t, err)
	assert.Equal(t, expected, count)
}

func TestCountXs(t *testing.T) {
	given := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 9

	count, err := advent2024.CountXs(given)

	assert.Nil(t, err)
	assert.Equal(t, expected, count)
}
