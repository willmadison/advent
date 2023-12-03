package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestFindPartNumbers(t *testing.T) {
	given := strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	expected := 8

	partNumbers, gearLocations := advent2023.FindPartNumbers(given)

	assert.Equal(t, expected, len(partNumbers))

	expectedSum := 4361

	var actualSum int

	for _, p := range partNumbers {
		actualSum += p.Number
	}

	assert.Equal(t, expectedSum, actualSum)

	expectedTotalGearRatio := 467835

	actualTotalGearRatio := advent2023.DetermineTotalGearRatio(partNumbers, gearLocations)

	assert.Equal(t, expectedTotalGearRatio, actualTotalGearRatio)
}
