package advent2017

import (
	"bufio"
	"strconv"
	"strings"
)

func Checksum(spreadsheet string) int {
	var checksum int

	scanner := bufio.NewScanner(strings.NewReader(spreadsheet))

	for scanner.Scan() {
		row := scanner.Text()

		min, max := determineMinMax(row)

		checksum += max - min
	}

	return checksum
}

func determineMinMax(row string) (int, int) {
	var min, max int

	numbers := strings.Fields(row)

	first, _ := strconv.Atoi(numbers[0])

	min, max = first, first

	numbers = numbers[1:]

	for _, number := range numbers {
		n, _ := strconv.Atoi(number)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max
}
