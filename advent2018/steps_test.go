package advent2018_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2018"
)

func TestDetermineInstructionOrder(t *testing.T) {
	cases := []struct {
		name     string
		given    io.Reader
		expected string
	}{
		{
			"Basic case...",
			strings.NewReader(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`),
			"CABDFE",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2018.DetermineStepOrder(tc.given)
			assert.Equal(t, tc.expected, actual)
		})

	}
}

func TestDetermineInstructionSLA(t *testing.T) {
	cases := []struct {
		name                   string
		given                  io.Reader
		numHelpers             int
		stepCompletionOverHead int
		expected               int
	}{
		{
			"1 helper case...",
			strings.NewReader(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`),
			1,
			0,
			15,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2018.DetermineInstructionSLA(tc.given, tc.numHelpers, tc.stepCompletionOverHead)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
