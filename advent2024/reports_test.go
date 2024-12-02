package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestParseReports(t *testing.T) {
	given := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	reports, err := advent2024.ParseReports(given)

	assert.Nil(t, err)
	assert.Equal(t, 6, len(reports))
	assert.Equal(t, 5, len(reports[0]))
}

func TestReportSafety(t *testing.T) {
	r := advent2024.Report{7, 6, 4, 2, 1}
	assert.True(t, r.IsSafe())

	r = advent2024.Report{1, 2, 7, 8, 9}
	assert.False(t, r.IsSafe())
}
