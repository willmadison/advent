package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestTrebuchet(t *testing.T) {
	given := strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	expected := 142

	actual := advent2023.ProcessCalibrationDocument(given)

	assert.Equal(t, expected, actual)
}

func TestSophisticatedTrebuchet(t *testing.T) {
	given := strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	expected := 281

	actual := advent2023.ProcessCalibrationDocument(given)

	assert.Equal(t, expected, actual)
}
