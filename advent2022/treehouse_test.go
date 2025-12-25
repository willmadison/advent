package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestCountVisibleTrees(t *testing.T) {
	given := strings.NewReader(`30373
25512
65332
33549
35390`)

	expectedVisible := 21

	visible, err := advent2022.CountVisibleTrees(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedVisible, visible)
}

func TestGetMaxScenicScore(t *testing.T) {
	given := strings.NewReader(`30373
25512
65332
33549
35390`)

	expectedScenicScore := 8

	scenicScore, err := advent2022.GetMaxScenicScore(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedScenicScore, scenicScore)
}
