package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassphraseValidation(t *testing.T) {
	cases := []struct {
		given Passphrase
		valid bool
	}{
		{
			"abcde fghij",
			true,
		},
		{
			"abcde xyz ecdab",
			false,
		},
		{
			"a ab abc abd abf abj",
			true,
		},
		{
			"iiii oiii ooii oooi oooo",
			true,
		},
		{
			"oiii ioii iioi iiio",
			false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.valid, c.given.IsValid())
	}
}

func TestCountValidPassphrases(t *testing.T) {
	assert.Equal(t, 2, CountValidPassphrases(`aa bb cc dd aaa
aa bb cc dd aa
aa bb cc dd ee`))
}
