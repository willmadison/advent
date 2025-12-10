package advent2025_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestMinimizeButtonPressesForMachineInitialization(t *testing.T) {
	cases := []struct {
		scenario        string
		r               io.Reader
		expectedPresses int
	}{
		{
			scenario:        "Two Presses (Machine 1)",
			r:               strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}`),
			expectedPresses: 2,
		},
		{
			scenario:        "Three Presses (Machine 2)",
			r:               strings.NewReader(`[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}`),
			expectedPresses: 3,
		},
		{
			scenario:        "Two Presses (Machine 3)",
			r:               strings.NewReader(`[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`),
			expectedPresses: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.scenario, func(t *testing.T) {
			presses, err := advent2025.MinimizeButtonPressesForMachineInitialization(tc.r)

			assert.Nil(t, err)

			assert.Equal(t, tc.expectedPresses, presses)
		})
	}

	given := strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`)

	expectedMinimumPresses := 7

	minimumPresses, err := advent2025.MinimizeButtonPressesForMachineInitialization(given)

	assert.Nil(t, err)
	assert.Equal(t, expectedMinimumPresses, minimumPresses)
}

func TestMinimizeButtonPressesForProperMachineJoltage(t *testing.T) {
	cases := []struct {
		scenario        string
		r               io.Reader
		expectedPresses int
	}{
		{
			scenario:        "Ten Presses (Machine 1)",
			r:               strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}`),
			expectedPresses: 10,
		},
		{
			scenario:        "Twelve Presses (Machine 2)",
			r:               strings.NewReader(`[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}`),
			expectedPresses: 12,
		},
		{
			scenario:        "Eleven Presses (Machine 3)",
			r:               strings.NewReader(`[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`),
			expectedPresses: 11,
		},
	}

	for _, tc := range cases {
		t.Run(tc.scenario, func(t *testing.T) {
			presses, err := advent2025.MinimizeButtonPressesForProperMachineJoltage(tc.r)

			assert.Nil(t, err)

			assert.Equal(t, tc.expectedPresses, presses)
		})
	}

	given := strings.NewReader(`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`)

	expectedMinimumPresses := 33

	minimumPresses, err := advent2025.MinimizeButtonPressesForProperMachineJoltage(given)

	assert.Nil(t, err)
	assert.Equal(t, expectedMinimumPresses, minimumPresses)
}
