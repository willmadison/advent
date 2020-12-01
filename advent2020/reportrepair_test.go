package advent2020_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2020"
)

func TestReportRepair(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`1721
979
366
299
675
1456`),
			514579,
		},
	}

	for _, tc := range cases {
		actual := advent2020.RepairReport(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestReportRepairTriplet(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`1721
979
366
299
675
1456`),
			241861950,
		},
	}

	for _, tc := range cases {
		actual := advent2020.RepairReportTriplet(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}
