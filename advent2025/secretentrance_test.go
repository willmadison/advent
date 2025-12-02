package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestCrackPassword(t *testing.T) {
	given := strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)

	expected := 3

	pw, err := advent2025.CrackPassword(given)

	assert.Nil(t, err)
	assert.Equal(t, expected, pw)
}

func TestCrackPasswordV2(t *testing.T) {
	given := strings.NewReader(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`)

	expected := 6

	pw, err := advent2025.CrackPasswordV2(given)

	assert.Nil(t, err)
	assert.Equal(t, expected, pw)
}
