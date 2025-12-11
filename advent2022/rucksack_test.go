package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindTotalPriorityFromRucksacks(t *testing.T) {
	given := strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	expectedTotalPriority := 157

	totalPriority, err := advent2022.FindTotalPriorityFromRucksacks(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedTotalPriority, totalPriority)
}

func TestFindTotalPriorityFromBadgesInRucksacks(t *testing.T) {
	given := strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	expectedTotalPriority := 70

	totalPriority, err := advent2022.FindTotalPriorityFromBadgesInRucksacks(given)

	assert.Nil(t, err)

	assert.Equal(t, expectedTotalPriority, totalPriority)
}
