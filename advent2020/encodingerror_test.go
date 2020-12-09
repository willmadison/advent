package advent2020

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindFirstEncodingError(t *testing.T) {
	given := strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)
	expected := 127

	preambleLength := 5
	actual, _ := FindFirstEncodingError(given, preambleLength, preambleLength)

	assert.Equal(t, expected, actual)
}

func TestFindEncryptionWeakness(t *testing.T) {
	given := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	target := 127
	expectedMin, expectedMax := 15, 47

	actualMin, actualMax := FindEncryptionWeakness(given, target)

	assert.Equal(t, expectedMin, actualMin)
	assert.Equal(t, expectedMax, actualMax)
}
