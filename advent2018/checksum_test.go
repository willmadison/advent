package advent2018

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestChecksum(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`),
			12,
		},
	}

	for _, tc := range cases {
		actual := Checksum(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCommonBoxIds(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected string
	}{
		{
			strings.NewReader(`abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`),
			"fgij",
		},
	}

	for _, tc := range cases {
		actual := CommonBoxIds(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}