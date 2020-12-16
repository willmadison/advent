package advent2020

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Validator interface {
	Validate(int) bool
}

type RangeValidator struct {
	LowerBound, UpperBound int
}

func (r RangeValidator) Validate(i int) bool {
	return i >= r.LowerBound && i <= r.UpperBound
}

type TicketRules map[string][]Validator

func (t TicketRules) FindErrorScanRate(tickets [][]int) (int, [][]int) {
	var invalids []int
	invalidValueTickets := map[int]struct{}{}

	for index, ticket := range tickets {
		for _, value := range ticket {
			var atLeastOneValidPlace bool

			for _, validators := range t {
				for _, validator := range validators {
					atLeastOneValidPlace = validator.Validate(value) || atLeastOneValidPlace
				}
			}

			if !atLeastOneValidPlace {
				invalids = append(invalids, value)
				invalidValueTickets[index] = struct{}{}
			}
		}
	}

	var errorRate int

	for _, invalid := range invalids {
		errorRate += invalid
	}

	var validTickets [][]int

	for index, ticket := range tickets {
		if _, present := invalidValueTickets[index]; !present {
			validTickets = append(validTickets, ticket)
		}
	}

	return errorRate, validTickets
}

func (t TicketRules) DetermineFieldLocale(tickets [][]int) map[string]int {
	validationResults := map[string][]int{}

	ticketLength := len(tickets[0])

	for field := range t {
		validationResults[field] = make([]int, ticketLength)
	}

	for _, ticket := range tickets {
		for ticketPosition, value := range ticket {
			for field, validators := range t {
				for _, validator := range validators {
					if validator.Validate(value) {
						validationResults[field][ticketPosition]++
						break
					}
				}
			}
		}
	}

	fieldLocale := map[string]int{}

	availableSlots := map[int]struct{}{}

	for i := 0; i < ticketLength; i++ {
		availableSlots[i] = struct{}{}
	}

	for len(fieldLocale) < len(validationResults) {
		for field := range validationResults {
			if _, placed := fieldLocale[field]; !placed {
				positions := place(field, validationResults, availableSlots, len(tickets))

				if len(positions) == 1 {
					position := positions[0]
					fieldLocale[field] = position
					delete(availableSlots, position)
				}
			}
		}
	}

	return fieldLocale
}

func place(field string, results map[string][]int, availableSlots map[int]struct{}, numTickets int) []int {
	result := results[field]

	possibilities := map[int]struct{}{}

	for slot := range availableSlots {
		if result[slot] == numTickets {
			possibilities[slot] = struct{}{}
		}
	}

	if len(possibilities) > 1 {
		toRemove := map[int]struct{}{}

		for possibility := range possibilities {
			worksInAllOthers := true

			for f, res := range results {
				if f == field {
					continue
				}

				worksInAllOthers = worksInAllOthers && res[possibility] == numTickets
			}

			if worksInAllOthers {
				toRemove[possibility] = struct{}{}
			}
		}

		for c := range toRemove {
			delete(possibilities, c)
		}
	}

	var positions []int

	for possibility := range possibilities {
		positions = append(positions, possibility)
	}

	return positions
}

func ParseTicketRules(r io.Reader) (TicketRules, [][]int) {
	rules := map[string][]Validator{}
	tickets := [][]int{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		value := scanner.Text()

		if value == "" || value == "your ticket:" || value == "nearby tickets:" {
			continue
		}

		if strings.Contains(value, ":") && len(tickets) == 0 {
			field, validators := parseTicketRule(value)
			rules[field] = validators
		} else {
			var ticketValues []int

			rawTicketValues := strings.Split(value, ",")

			for _, v := range rawTicketValues {
				val, _ := strconv.Atoi(v)
				ticketValues = append(ticketValues, val)
			}

			tickets = append(tickets, ticketValues)
		}
	}

	return TicketRules(rules), tickets
}

func parseTicketRule(value string) (string, []Validator) {
	parts := strings.Split(value, ": ")

	field := parts[0]

	var validators []Validator

	rawConstraintParts := strings.Split(parts[1], " or ")

	for _, part := range rawConstraintParts {
		rawBounds := strings.Split(part, "-")

		lowerBound, _ := strconv.Atoi(rawBounds[0])
		upperBound, _ := strconv.Atoi(rawBounds[1])

		validators = append(validators, RangeValidator{lowerBound, upperBound})
	}

	return field, validators
}
