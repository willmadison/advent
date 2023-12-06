package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestFindWinningRaceStrategies(t *testing.T) {
	given := strings.NewReader(`Time:      7  15   30
Distance:  9  40  200`)

	strategies := advent2023.FindWinningRaceStrategies(given)

	expectedNumStrategies := 3

	assert.Equal(t, expectedNumStrategies, len(strategies))

	expectedTotal := 288

	actualTotal := 1

	for _, strategy := range strategies {
		actualTotal *= len(strategy.Winners)
	}

	assert.Equal(t, expectedTotal, actualTotal)
}

func TestFindWinningRaceStrategiesWithGoodKerning(t *testing.T) {
	given := strings.NewReader(`Time:      7  15   30
Distance:  9  40  200`)

	strategies := advent2023.FindWinningRaceStrategies(given, advent2023.Good)

	expectedNumStrategies := 1

	assert.Equal(t, expectedNumStrategies, len(strategies))

	expectedTotal := 71503

	actualTotal := 1

	for _, strategy := range strategies {
		actualTotal *= len(strategy.Winners)
	}

	assert.Equal(t, expectedTotal, actualTotal)
}
