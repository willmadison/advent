package advent2019_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2019"
)

func TestRunIntCodeMachine(t *testing.T) {
	cases := []struct {
		name     string
		given    io.Reader
		expected []int
	}{
		{
			"1,0,0,0,99",
			strings.NewReader(`1,0,0,0,99`),
			[]int{2, 0, 0, 0, 99},
		},
		{
			"2,3,0,3,99",
			strings.NewReader(`2,3,0,3,99`),
			[]int{2, 3, 0, 6, 99},
		},
		{
			"2,4,4,5,99,0",
			strings.NewReader(`2,4,4,5,99,0`),
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"1,1,1,4,99,5,6,0,99",
			strings.NewReader(`1,1,1,4,99,5,6,0,99`),
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2019.RunIntCodeMachine(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
