package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestMatchPairs(t *testing.T) {
	given := strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`)

	expected := [][]int{
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 9},
	}

	actual := advent2024.MatchPairs(given)

	assert.ElementsMatch(t, expected, actual)
}

func TestFindTotalDistance(t *testing.T) {
	given := [][]int{
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 9},
	}

	actual := advent2024.FindTotalDistance(given)

	assert.Equal(t, 11, actual)
}

func TestParseLists(t *testing.T) {
	given := strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`)

	leftExpected := []int{3, 4, 2, 1, 3, 3}
	rightExpected := []int{4, 3, 5, 3, 9, 3}

	leftActual, rightActual := advent2024.ParseLists(given)

	assert.ElementsMatch(t, leftExpected, leftActual)
	assert.ElementsMatch(t, rightExpected, rightActual)
}

func TestSimilarityScore(t *testing.T) {
	given := strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`)

	expected := 31

	actual := advent2024.SimilarityScore(given)

	assert.Equal(t, expected, actual)
}
