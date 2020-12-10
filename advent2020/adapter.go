package advent2020

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func FindJoltDifferences(r io.Reader) map[int]int {
	jolts := parseJoltRatings(r)

	sort.Ints(jolts)

	countsByJoltDifference := map[int]int{
		3: 1,
	}
	used := map[int]struct{}{}

	var last int

	for len(used) < len(jolts) {
		nextIndex := findNextAdapter(last, jolts, used)
		next := jolts[nextIndex]

		countsByJoltDifference[next-last]++
		used[nextIndex] = struct{}{}

		last = next
	}

	return countsByJoltDifference
}

func parseJoltRatings(r io.Reader) []int {
	scanner := bufio.NewScanner(r)

	var jolts []int

	for scanner.Scan() {
		jolt, _ := strconv.Atoi(scanner.Text())
		jolts = append(jolts, jolt)
	}
	return jolts
}

func findNextAdapter(lastAdapter int, jolts []int, used map[int]struct{}) int {
	for i, jolt := range jolts {
		if _, consumed := used[i]; consumed {
			continue
		}

		diff := jolt - lastAdapter

		if diff >= 1 && diff <= 3 {
			return i
		}
	}

	return -1
}

func CountDistinctPossibleArrangements(r io.Reader) int {
	jolts := parseJoltRatings(r)

	sort.Ints(jolts)

	positionsByJolt := map[int]int{}

	for i, jolt := range jolts {
		positionsByJolt[jolt] = i
	}

	ways := make([]int, len(jolts))

	ways[len(jolts)-1] = 1

	for i := len(jolts) - 2; i >= 0; i-- {
		sum := 0

		for diff := 1; diff <= 3; diff++ {
			if position, present := positionsByJolt[jolts[i]+diff]; present {
				sum += ways[position]
			}
		}

		ways[i] = sum
	}

	var totalWays int

	for i := 1; i <= 3; i++ {
		if position, present := positionsByJolt[i]; present {
			totalWays += ways[position]
		}
	}

	return totalWays
}
