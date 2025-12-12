package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestCountGiftRegionFits(t *testing.T) {
	given := strings.NewReader(`0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`)

	expectedFits := 2

	fits, err := advent2025.CountGiftRegionFits(given)
	assert.Nil(t, err)

	assert.Equal(t, expectedFits, fits)
}
