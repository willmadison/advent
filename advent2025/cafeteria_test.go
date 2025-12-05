package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindFreshIngredients(t *testing.T) {
	given := strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	expectedNumFresh := 3

	freshIngredients, err := advent2025.FindFreshIngredients(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedNumFresh, len(freshIngredients))
}

func TestEnumerateFreshIngredientIds(t *testing.T) {
	given := strings.NewReader(`3-5
10-14
16-20
12-18

1
5
8
11
17
32`)

	expectedNumFreshIds := 14

	numFreshIngredientIds, err := advent2025.EnumerateFreshIngredientIds(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedNumFreshIds, numFreshIngredientIds)
}
