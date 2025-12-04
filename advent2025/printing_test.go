package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindAccessiblePaperRolls(t *testing.T) {
	given := strings.NewReader(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

	numExpectedAccessiblePaperRolls := 13

	accessiblePaperRolls, err := advent2025.FindAccessiblePaperRolls(given)

	assert.Nil(t, err)

	assert.Equal(t, numExpectedAccessiblePaperRolls, len(accessiblePaperRolls))
}

func TestRemoveAllAccessiblePaperRolls(t *testing.T) {
	given := strings.NewReader(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

	numExpectedRollsRemoved := 43

	rollsRemoved, err := advent2025.RemoveAllAccessiblePaperRolls(given)

	assert.Nil(t, err)

	assert.Equal(t, numExpectedRollsRemoved, rollsRemoved)
}
