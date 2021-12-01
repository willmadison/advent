package advent2021

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

type Strategy uint

const (
	IndividualReadingStrategy Strategy = iota
	SlidingWindowStrategy
)

func SonarDepthIncreaseRate(r io.Reader, strategies ...Strategy) int {
	var strategy Strategy

	if len(strategies) == 0 {
		strategy = IndividualReadingStrategy
	} else {
		strategy = strategies[0]
	}

	readings := parseSonarReadings(r, strategy)

	var increaseRate int

	lastReading := math.MaxInt

	for _, reading := range readings {
		if reading-lastReading > 0 {
			increaseRate++
		}

		lastReading = reading
	}

	return increaseRate
}

func parseSonarReadings(r io.Reader, strategy Strategy) []int {
	scanner := bufio.NewScanner(r)

	var individualReadings []int

	for scanner.Scan() {
		reading, _ := strconv.Atoi(scanner.Text())
		individualReadings = append(individualReadings, reading)
	}

	if strategy == SlidingWindowStrategy {
		var readings []int
		windowSize := 3

		for i := 0; i < len(individualReadings); i++ {
			readings = append(readings, sum(individualReadings[i:i+windowSize]))
		}

		return readings
	}

	return individualReadings
}

func sum(nums []int) int {
	var sum int

	for _, v := range nums {
		sum += v
	}

	return sum
}
