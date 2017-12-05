package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverseCaptcha(t *testing.T) {
	cases := []struct {
		given    string
		expected int
	}{
		{
			"1122",
			3,
		},
		{
			"1111",
			4,
		},
		{
			"1234",
			0,
		},
		{
			"91212129",
			9,
		},
	}

	for _, c := range cases {
		actual := InverseCaptcha(c.given)
		assert.Equal(t, c.expected, actual)
	}

}
