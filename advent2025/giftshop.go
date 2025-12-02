package advent2025

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/intervals"
)

type ValidityStrategy int

const (
	ExactlyTwice ValidityStrategy = iota
	AtLeastTwice
)

func FindInvalidIdentifiers(r io.Reader, strategy ValidityStrategy) ([]int64, error) {
	var invalidIdentifiers []int64

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		rawRanges := strings.Split(line, ",")

		for _, rawRange := range rawRanges {

			rng := intervals.ParseRange(rawRange)

			for id := range rng.Iter() {
				if isInvalidIdentifier(id, strategy) {
					invalidIdentifiers = append(invalidIdentifiers, id)
				}
			}

		}
	}

	return invalidIdentifiers, nil
}

func isInvalidIdentifier(id int64, strategy ValidityStrategy) bool {
	identifier := strconv.FormatInt(id, 10)

	switch strategy {
	case ExactlyTwice:
		return isExactlyTwice(identifier)
	case AtLeastTwice:
		return isAtLeastTwice(identifier)
	default:
		return false
	}
}

func isExactlyTwice(identifier string) bool {
	n := len(identifier)

	if n%2 != 0 {
		return false
	}

	midpoint := n / 2

	left, right := identifier[:midpoint], identifier[midpoint:]

	return left == right
}

func isAtLeastTwice(identifier string) bool {
	if len(identifier) <= 1 {
		return false
	}

	doubled := identifier + identifier
	doubled = doubled[1 : len(doubled)-1]

	return strings.Contains(doubled, identifier)
}
