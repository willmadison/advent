package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveLoopNumber(t *testing.T) {
	cases := []struct {
		given    int
		expected int
	}{
		{
			5764801,
			8,
		},
		{
			17807724,
			11,
		},
	}
	var e Encryptor

	for _, tc := range cases {
		e.DeriveLoopNumber(tc.given)
		assert.Equal(t, tc.expected, e.LoopNumber)
	}
}

func TestDetermineEncryptionKey(t *testing.T) {
	given := strings.NewReader(`5764801
17807724`)

	expectedEncryptionKey := 14897079

	encryptionKey := DetermineEncryptionKey(given)

	assert.Equal(t, expectedEncryptionKey, encryptionKey)
}
