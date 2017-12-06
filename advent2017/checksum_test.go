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
			`5 1 9 5
7 5 3
2 4 6 8`,
			18,
		},
	}

	for _, c := range cases {
		actual := Checksum(c.given)
		assert.Equal(t, c.expected, actual)
	}

}
