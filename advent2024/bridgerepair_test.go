package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestParseCalibrationEquations(t *testing.T) {
	given := strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations, err := advent2024.ParseCalibrationEquations(given)

	assert.Nil(t, err)

	expectedNumberOfEquations := 9

	assert.Equal(t, expectedNumberOfEquations, len(equations))

	expectedExpectedValue := 190
	expectedNumberOfOperands := 2

	assert.Equal(t, expectedExpectedValue, equations[0].ExpectedValue)
	assert.True(t, expectedNumberOfOperands > 0)
	assert.Equal(t, expectedNumberOfOperands, len(equations[0].Operands))
}

func TestOperatorAlternatives(t *testing.T) {
	given := strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations, err := advent2024.ParseCalibrationEquations(given)

	assert.Nil(t, err)

	equation := equations[0]

	alternatives := equation.OperatorAlternatives()

	assert.Equal(t, 2, len(alternatives))
}

func TestCanBeMadeTrue(t *testing.T) {
	given := strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations, err := advent2024.ParseCalibrationEquations(given)

	assert.Nil(t, err)

	expectedCouldBeMadeTrue := 3

	var actuallyCouldBeMadeTrue int

	for _, equation := range equations {
		if equation.CouldBeMadeTrue() {
			actuallyCouldBeMadeTrue++
		}

	}

	assert.Equal(t, expectedCouldBeMadeTrue, actuallyCouldBeMadeTrue)
}

func TestCanBeMadeTrueExpandedOperatorSet(t *testing.T) {
	given := strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	equations, err := advent2024.ParseCalibrationEquations(given)

	assert.Nil(t, err)

	expectedCouldBeMadeTrue := 6

	var actuallyCouldBeMadeTrue int

	for _, equation := range equations {
		if equation.CouldBeMadeTrue(advent2024.AllOperators...) {
			actuallyCouldBeMadeTrue++
		}

	}

	assert.Equal(t, expectedCouldBeMadeTrue, actuallyCouldBeMadeTrue)
}
