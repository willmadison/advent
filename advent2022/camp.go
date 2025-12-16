package advent2022

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/willmadison/advent/internal/intervals"
)

func FindFullyContainedCleanupRangePairs(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	fullyContainedPairs := 0

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		if len(ranges) != 2 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}

		a := intervals.ParseRange(ranges[0])
		b := intervals.ParseRange(ranges[1])

		if a.ContainsRange(b) || b.ContainsRange(a) {
			fullyContainedPairs++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return fullyContainedPairs, nil
}

func FindOverlappingCleanupRangePairs(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	overlappingPairs := 0

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		if len(ranges) != 2 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}

		a := intervals.ParseRange(ranges[0])
		b := intervals.ParseRange(ranges[1])

		if a.Overlaps(b) || b.Overlaps(a) {
			overlappingPairs++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return overlappingPairs, nil
}
