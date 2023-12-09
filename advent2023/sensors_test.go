package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestExtrapolateSensorNextReadings(t *testing.T) {
	given := strings.NewReader(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

	expectedPreviousSum := 2
	expectedNextSum := 114

	readings := advent2023.ExtrapolateSenorReadings(given)

	var previousSum int
	var nextSum int

	for _, reading := range readings {
		previousSum += reading[0]
		nextSum += reading[len(reading)-1]
	}

	assert.Equal(t, expectedPreviousSum, previousSum)
	assert.Equal(t, expectedNextSum, nextSum)
}
