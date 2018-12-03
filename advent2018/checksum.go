package advent2018

import (
	"bufio"
	"io"
)

func Checksum(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var doubles, triples int

	for scanner.Scan() {
		id := scanner.Text()

		counts := map[rune]int{}

		for _, c := range id {
			counts[c]++
		}

		var hasDouble, hasTriple bool

		for _, count := range counts {
			switch count {
			case 2:
				hasDouble = true
			case 3:
				hasTriple = true
			}
		}

		if hasDouble {
			doubles++
		}

		if hasTriple {
			triples++
		}
	}

	return doubles * triples
}

func CommonBoxIds(r io.Reader) string {
	var boxIds []string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		boxIds = append(boxIds, scanner.Text())
	}

	a, b := findLevenshteinPair(boxIds, 1)

	if a != "" {
		var prefixEnd int

		for a[prefixEnd] == b[prefixEnd] {
			prefixEnd++
		}

		return a[0:prefixEnd] + a[prefixEnd+1:]
	}

	return a
}

func findLevenshteinPair(values []string, desiredDistance int) (string, string) {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			distance := calculateLevenshteinDistance(values[i], values[j])

			if distance == desiredDistance {
				return values[i], values[j]
			}
		}
	}

	return "", ""
}

type pair struct {
	x, y int
}

func calculateLevenshteinDistance(a string, b string) int {
	if a == b {
		return 0
	}

	cache := map[pair]int{}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			switch {
			case i == 0:
				cache[pair{i, j}] = j
			case j == 0:
				cache[pair{i, j}] = i
			default:
				cache[pair{i, j}] = min(
					cache[pair{i - 1, j - 1}]+costOfSubstitution(rune(a[i-1]), rune(b[j-1])),
					cache[pair{i - 1, j}]+1,
					cache[pair{i, j - 1}]+1)
			}
		}
	}

	return cache[pair{len(a), len(b)}]
}

func costOfSubstitution(a rune, b rune) int {
	if a == b {
		return 0
	} else {
		return 1
	}
}

func min(numbers ...int) int {
	var min = numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}

	return min
}
