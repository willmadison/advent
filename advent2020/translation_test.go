package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTicketRules(t *testing.T) {
	given := strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	expectedRules := TicketRules(map[string][]Validator{
		"class": {RangeValidator{1, 3}, RangeValidator{5, 7}},
		"row":   {RangeValidator{6, 11}, RangeValidator{33, 44}},
		"seat":  {RangeValidator{13, 40}, RangeValidator{45, 50}},
	})

	expectedTickets := [][]int{
		{7, 1, 14},
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	actualRules, actualTickets := ParseTicketRules(given)

	assert.Equal(t, expectedRules, actualRules)
	assert.Equal(t, expectedTickets, actualTickets)
}

func TestFindScanErrorRate(t *testing.T) {

	rules := TicketRules(map[string][]Validator{
		"class": {RangeValidator{1, 3}, RangeValidator{5, 7}},
		"row":   {RangeValidator{6, 11}, RangeValidator{33, 44}},
		"seat":  {RangeValidator{13, 40}, RangeValidator{45, 50}},
	})

	tickets := [][]int{
		{7, 1, 14},
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	expectedErrorRate := 71

	expectedValidTickets := [][]int{
		{7, 3, 47},
	}

	actualErrorRate, actualValidTickets := rules.FindErrorScanRate(tickets[1:])

	assert.Equal(t, expectedErrorRate, actualErrorRate)
	assert.ElementsMatch(t, expectedValidTickets, actualValidTickets)
}

func TestDetermineFieldLocale(t *testing.T) {
	rules := TicketRules(map[string][]Validator{
		"class": {RangeValidator{0, 1}, RangeValidator{4, 19}},
		"row":   {RangeValidator{0, 5}, RangeValidator{8, 19}},
		"seat":  {RangeValidator{0, 13}, RangeValidator{16, 19}},
	})

	tickets := [][]int{
		{11, 12, 13},
		{3, 9, 18},
		{15, 1, 5},
		{5, 14, 9},
	}

	expectedFieldLocale := map[string]int{
		"class": 1,
		"row":   0,
		"seat":  2,
	}

	actualFieldLocale := rules.DetermineFieldLocale(tickets[1:])
	assert.Equal(t, expectedFieldLocale, actualFieldLocale)
}
