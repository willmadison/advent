package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindJoltDifferences(t *testing.T) {
	given := strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`)

	expected := map[int]int{
		1: 7,
		3: 5,
	}

	actual := FindJoltDifferences(given)
	assert.Equal(t, expected, actual)
}

func TestCountDistinctPossibleArrangements(t *testing.T) {
	given := strings.NewReader(`16
10
15
5
1
11
7
19
6
12
4`)

	expected := 8

	actual := CountDistinctPossibleArrangements(given)
	assert.Equal(t, expected, actual)
}
