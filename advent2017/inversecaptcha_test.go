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
			"1212",
			6,
		},
		{
			"1221",
			0,
		},
		{
			"123425",
			4,
		},
		{
			"123123",
			12,
		},
		{
			"12131415",
			4,
		},
	}

	for _, c := range cases {
		actual := InverseCaptcha(c.given)
		assert.Equal(t, c.expected, actual)
	}

}
