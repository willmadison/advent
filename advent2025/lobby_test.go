package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindMaximumJoltages(t *testing.T) {
	given := strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	expectedJoltages := []int64{98, 89, 78, 92}

	joltages, err := advent2025.FindMaximumJoltages(given, 2)

	assert.Nil(t, err)

	assert.ElementsMatch(t, expectedJoltages, joltages)

	given = strings.NewReader(`987654321111111
811111111111119
234234234234278
818181911112111`)

	expectedJoltages = []int64{987654321111, 811111111119, 434234234278, 888911112111}

	joltages, err = advent2025.FindMaximumJoltages(given, 12)

	assert.Nil(t, err)

	assert.ElementsMatch(t, expectedJoltages, joltages)

}
