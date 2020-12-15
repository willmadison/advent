package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNthSpokenNumber(t *testing.T) {
	given := strings.NewReader(`0,3,6`)

	expected := 436

	actual := FindNthSpokenNumber(given, 2020)

	assert.Equal(t, expected, actual)
}
