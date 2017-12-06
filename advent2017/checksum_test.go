package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	cases := []struct {
		given    string
		expected int
	}{
		{
			`5 10 9 7
7 6 3
9 4 6 8`,
			6,
		},
	}

	for _, c := range cases {
		actual := Checksum(c.given)
		assert.Equal(t, c.expected, actual)
	}

}
