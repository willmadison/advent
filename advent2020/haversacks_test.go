package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBagRules(t *testing.T) {
	given := strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	expected := BagRules{
		"faded blue":   {},
		"dotted black": {},
		"vibrant plum": {
			"faded blue":   5,
			"dotted black": 6,
		},
		"dark olive": {
			"faded blue":   3,
			"dotted black": 4,
		},
		"shiny gold": {
			"dark olive":   1,
			"vibrant plum": 2,
		},
		"muted yellow": {
			"shiny gold": 2,
			"faded blue": 9,
		},
		"bright white": {
			"shiny gold": 1,
		},
		"dark orange": {
			"bright white": 3,
			"muted yellow": 4,
		},
		"light red": {
			"bright white": 1,
			"muted yellow": 2,
		},
	}

	actual := ParseBagRules(given)
	assert.Equal(t, expected, actual)
}

func TestFindAncestorsOf(t *testing.T) {
	given := BagRules{
		"faded blue":   {},
		"dotted black": {},
		"vibrant plum": {
			"faded blue":   5,
			"dotted black": 6,
		},
		"dark olive": {
			"faded blue":   3,
			"dotted black": 4,
		},
		"shiny gold": {
			"dark olive":   1,
			"vibrant plum": 2,
		},
		"muted yellow": {
			"shiny gold": 2,
			"faded blue": 9,
		},
		"bright white": {
			"shiny gold": 1,
		},
		"dark orange": {
			"bright white": 3,
			"muted yellow": 4,
		},
		"light red": {
			"bright white": 1,
			"muted yellow": 2,
		},
	}

	expected := []string{
		"bright white",
		"muted yellow",
		"light red",
		"dark orange",
	}

	actual := given.FindAncestorsOf("shiny gold")
	assert.ElementsMatch(t, expected, actual)
}

func TestTotalDescendantsOf(t *testing.T) {
	given := BagRules{
		"faded blue":   {},
		"dotted black": {},
		"vibrant plum": {
			"faded blue":   5,
			"dotted black": 6,
		},
		"dark olive": {
			"faded blue":   3,
			"dotted black": 4,
		},
		"shiny gold": {
			"dark olive":   1,
			"vibrant plum": 2,
		},
		"muted yellow": {
			"shiny gold": 2,
			"faded blue": 9,
		},
		"bright white": {
			"shiny gold": 1,
		},
		"dark orange": {
			"bright white": 3,
			"muted yellow": 4,
		},
		"light red": {
			"bright white": 1,
			"muted yellow": 2,
		},
	}

	expected := 32

	actual := given.TotalDescendantsOf("shiny gold")
	assert.Equal(t, expected, actual)
}
