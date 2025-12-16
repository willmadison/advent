package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindFullyContainedCleanupRangePairs(t *testing.T) {
	given := strings.NewReader(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	expectedFullyContainedPairs := 2

	fullyContainedPairs, err := advent2022.FindFullyContainedCleanupRangePairs(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedFullyContainedPairs, fullyContainedPairs)
}

func TestFindOverlappingCleanupRangePairs(t *testing.T) {
	given := strings.NewReader(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

	expectedOverlappingPairs := 4

	overlappingPairs, err := advent2022.FindOverlappingCleanupRangePairs(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedOverlappingPairs, overlappingPairs)
}
