package advent2020

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePasswordDatabase(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected []PasswordDBEntry
	}{
		{
			strings.NewReader(`1-3 a: abcde`),
			[]PasswordDBEntry{
				{
					Policy:   PasswordPolicy{'a', 1, 3},
					Password: `abcde`,
				},
			},
		},
	}

	for _, tc := range cases {
		actual := parsePasswordDBEntries(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCountValidPasswords(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`),
			2,
		},
	}

	for _, tc := range cases {
		actual := CountValidPasswords(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestCountValidPasswordsUpdatedPolicy(t *testing.T) {
	cases := []struct {
		given    io.Reader
		expected int
	}{
		{
			strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`),
			1,
		},
	}

	for _, tc := range cases {
		actual := CountValidPasswordsUpdatedPolicy(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}
