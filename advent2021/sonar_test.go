package advent2021_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2021"
)

func TestSonarDepthIncreaseRate(t *testing.T) {
	given := strings.NewReader(`199
200
208
210
200
207
240
269
260
263`)

	expected := 7

	actual := advent2021.SonarDepthIncreaseRate(given)
	assert.Equal(t, expected, actual)
}

func TestSonarDepthIncreaseRateSlidingWindow(t *testing.T) {
	given := strings.NewReader(`199
200
208
210
200
207
240
269
260
263`)

	expected := 5

	actual := advent2021.SonarDepthIncreaseRate(given, advent2021.SlidingWindowStrategy)
	assert.Equal(t, expected, actual)
}
