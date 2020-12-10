package advent2020

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

func FindFirstEncodingError(r io.Reader, preambleLength, lookbackLength int) (int, []int) {
	values := parseValues(r)

	i := preambleLength

	for i < len(values) {
		previous := values[i-lookbackLength : i]
		value := values[i]

		if !hasTwoSum(value, previous) {
			return value, values
		}

		i++
	}

	return -1, nil
}

func parseValues(r io.Reader) []int {
	scanner := bufio.NewScanner(r)

	var values []int

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		values = append(values, v)
	}

	return values
}

func hasTwoSum(value int, values []int) bool {
	uniqueValues := map[int]struct{}{}

	for _, v := range values {
		uniqueValues[v] = struct{}{}
	}

	for _, v := range values {
		otherAddend := value - v
		if _, present := uniqueValues[otherAddend]; present && otherAddend != v {
			return true
		}
	}

	return false
}

func FindEncryptionWeakness(values []int, target int) (int, int) {
	sublist := findSublistSum(values, target)

	if len(sublist) == 0 {
		return -1, -1
	}

	return min(sublist), max(sublist)
}

func findSublistSum(values []int, target int) []int {
	var start, end int

	for start < len(values) {
		var sum int

		current := start
		for sum < target {
			sum += values[current]
			current++
		}

		if sum == target {
			end = current - 1
			break
		}

		start++
	}

	return values[start:end]
}

func min(values []int) int {
	minimum := math.MaxInt64

	for _, v := range values {
		if v < minimum {
			minimum = v
		}
	}

	return minimum
}

func max(values []int) int {
	maximum := math.MinInt64

	for _, v := range values {
		if v > maximum {
			maximum = v
		}
	}

	return maximum
}
