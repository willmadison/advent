package advent2025_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindInvalidIdentifiers(t *testing.T) {
	given := strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)

	expectedSum := int64(1227775554)

	identifiers, err := advent2025.FindInvalidIdentifiers(given, advent2025.ExactlyTwice)
	fmt.Println("identifiers:", identifiers)

	assert.Nil(t, err)

	var sum int64

	for _, id := range identifiers {
		sum += id
	}

	assert.Equal(t, expectedSum, sum)

	given = strings.NewReader(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)

	expectedSum = int64(4174379265)

	identifiers, err = advent2025.FindInvalidIdentifiers(given, advent2025.AtLeastTwice)
	fmt.Println("identifiers:", identifiers)

	assert.Nil(t, err)

	sum = int64(0)

	for _, id := range identifiers {
		sum += id
	}

	assert.Equal(t, expectedSum, sum)
}
