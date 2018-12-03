package advent2018

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Calibrate(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var frequency int

	for scanner.Scan() {
		rawChange := scanner.Text()

		change, _ := strconv.Atoi(rawChange)
		frequency += change
	}

	return frequency
}

func CalibrateDuplication(in string) int {
	var duplicateFound bool

	var duplicate int

	seen := map[int]struct{}{}
	seen[0] = struct{}{}

	var frequency int

	for !duplicateFound {
		scanner := bufio.NewScanner(strings.NewReader(in))

		for scanner.Scan() {
			rawChange := scanner.Text()

			change, _ := strconv.Atoi(rawChange)
			frequency += change

			if _, present := seen[frequency]; !present {
				seen[frequency] = struct{}{}
			} else {
				duplicateFound = true
				duplicate = frequency
				break
			}
		}
	}

	return duplicate
}
