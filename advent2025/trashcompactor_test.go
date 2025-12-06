package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestCheckWorksheet(t *testing.T) {
	given := strings.NewReader(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `)

	expectedTotal := int64(4277556)

	subtotals, err := advent2025.CheckWorksheet(given, advent2025.Human)
	assert.Nil(t, err)

	var total int64

	for _, v := range subtotals {
		total += v
	}

	assert.Equal(t, expectedTotal, total)

	given = strings.NewReader(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `)

	expectedTotal = int64(3263827)

	subtotals, err = advent2025.CheckWorksheet(given, advent2025.Cephalopod)
	assert.Nil(t, err)

	total = int64(0)

	for _, v := range subtotals {
		total += v
	}

	assert.Equal(t, expectedTotal, total)
}
