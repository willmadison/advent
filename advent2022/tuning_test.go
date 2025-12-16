package advent2022_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindFirstMarkerIndex(t *testing.T) {
	cases := []struct {
		given               io.Reader
		windowSize          int
		expectedMarkerIndex int
	}{
		{given: strings.NewReader(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`), windowSize: 4, expectedMarkerIndex: 7},
		{given: strings.NewReader(`mjqjpqmgbljsphdztnvjfqwrcgsmlb`), windowSize: 14, expectedMarkerIndex: 19},
		{given: strings.NewReader(`bvwbjplbgvbhsrlpgdmjqwftvncz`), windowSize: 4, expectedMarkerIndex: 5},
		{given: strings.NewReader(`bvwbjplbgvbhsrlpgdmjqwftvncz`), windowSize: 14, expectedMarkerIndex: 23},
		{given: strings.NewReader(`nppdvjthqldpwncqszvftbrmjlhg`), windowSize: 4, expectedMarkerIndex: 6},
		{given: strings.NewReader(`nppdvjthqldpwncqszvftbrmjlhg`), windowSize: 14, expectedMarkerIndex: 23},
		{given: strings.NewReader(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`), windowSize: 4, expectedMarkerIndex: 10},
		{given: strings.NewReader(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`), windowSize: 14, expectedMarkerIndex: 29},
		{given: strings.NewReader(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`), windowSize: 4, expectedMarkerIndex: 11},
		{given: strings.NewReader(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`), windowSize: 14, expectedMarkerIndex: 26},
	}

	for _, tc := range cases {
		markerIndex, err := advent2022.FindFirstMarkerIndex(tc.given, tc.windowSize)
		assert.Nil(t, err)

		assert.Equal(t, tc.expectedMarkerIndex, markerIndex)
	}
}
