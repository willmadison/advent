package advent2018

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibrate(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`+1
-2
+3
+1`),
			3,
		},
		{
			strings.NewReader(`+1
+1
+1`),
			3,
		},
		{
			strings.NewReader(`+1
+1
-2`),
			0,
		},
	}

	for _, tc := range cases {
		actual := Calibrate(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCalibrateDuplication(t *testing.T) {
	cases := []struct {
		given    string
		expected int
	}{
		{
			`+1
-2
+3
+1`,
			2,
		},
		{
			`+1
-1`,
			0,
		},
		{
			`+3
+3
+4
-2
-4`,
			10,
		},
		{
			`-6
+3
+8
+5
-6`,
			5,
		},
	}

	for _, tc := range cases {
		actual := CalibrateDuplication(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}
