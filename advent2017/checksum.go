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

		divisor, dividend := determineDivisorDividend(row)

		checksum += dividend / divisor
	}

	return checksum
}

func determineDivisorDividend(row string) (int, int) {
	var divisor, dividend, i, j int

	var numbers []int

	for _, n := range strings.Fields(row) {
		i, _ := strconv.Atoi(n)
		numbers = append(numbers, i)
	}

	for i, divisor = range numbers {
		for j, dividend = range numbers {
			if i == j {
				continue
			}

			if dividend % divisor == 0 {
				return divisor, dividend
			}
		}
	}

	return divisor, dividend
}
