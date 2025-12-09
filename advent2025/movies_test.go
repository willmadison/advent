package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindMaxiumumAreaRectangle(t *testing.T) {
	given := strings.NewReader(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	expectedArea := 50

	area, err := advent2025.FindMaxiumAreaRectangle(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedArea, area)
}

func TestFindMaxiumumAreaRectangleWithConstraints(t *testing.T) {
	given := strings.NewReader(`7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`)

	expectedArea := 24

	area, err := advent2025.FindMaxiumAreaRectangleWithConstraints(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedArea, area)
}
