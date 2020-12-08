package advent2020

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseProgram(t *testing.T) {
	given := strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	expected := []Command{
		{NoOp, 0},
		{Accumulate, 1},
		{Jump, 4},
		{Accumulate, 3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, 1},
		{Jump, -4},
		{Accumulate, 6},
	}

	actual := ParseProgram(given)
	assert.Equal(t, expected, actual)
}

func TestDetermineAccumulatorValueBeforeLoop(t *testing.T) {
	given := []Command{
		{NoOp, 0},
		{Accumulate, 1},
		{Jump, 4},
		{Accumulate, 3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, 1},
		{Jump, -4},
		{Accumulate, 6},
	}

	expected := 5

	actual := DetermineAccumulatorValueBeforeLoop(given)
	assert.Equal(t, expected, actual)
}

func TestPatchProgram(t *testing.T) {
	given := []Command{
		{NoOp, 0},
		{Accumulate, 1},
		{Jump, 4},
		{Accumulate, 3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, 1},
		{Jump, -4},
		{Accumulate, 6},
	}

	expected := []Command{
		{NoOp, 0},
		{Accumulate, 1},
		{Jump, 4},
		{Accumulate, 3},
		{Jump, -3},
		{Accumulate, -99},
		{Accumulate, 1},
		{NoOp, -4},
		{Accumulate, 6},
	}

	actual := PatchProgram(given)
	assert.Equal(t, expected, actual)

	expectedAccumulator := 8

	actualAccumulator := Run(actual)
	assert.Equal(t, expectedAccumulator, actualAccumulator)
}

