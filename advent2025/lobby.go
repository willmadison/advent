package advent2025

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/stacks"
)

func FindMaximumJoltages(r io.Reader, maxDigits int) ([]int64, error) {
	var joltages []int64

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		rawDigits := strings.Split(line, "")

		var digits []int

		for _, rawDigit := range rawDigits {
			digit, err := strconv.Atoi(rawDigit)
			if err != nil {
				return nil, err
			}

			digits = append(digits, digit)
		}

		joltage := determineMaximumJoltage(digits, maxDigits)

		joltages = append(joltages, joltage)
	}

	return joltages, nil
}

func determineMaximumJoltage(digits []int, maxDigits int) int64 {

	if maxDigits == 2 {
		return determineMaximumTwoDigitJoltage(digits)
	}

	return determineMaximimumNDigitJoltage(digits, maxDigits)
}

func determineMaximumTwoDigitJoltage(digits []int) int64 {
	bestTens := -1
	maxJoltage := -1

	for _, digit := range digits {
		if bestTens != -1 {
			joltage := bestTens*10 + digit

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}

		if digit > bestTens {
			bestTens = digit
		}
	}

	return int64(maxJoltage)
}

func determineMaximimumNDigitJoltage(digits []int, maxDigits int) int64 {
	numDigits := len(digits)

	if maxDigits <= 0 || maxDigits > numDigits {
		panic("invalid input")
	}

	removalsAllowed := numDigits - maxDigits

	stack := stacks.NewStack[int]()

	for _, digit := range digits {
		for removalsAllowed > 0 && stack.Size() > 0 {
			top, _ := stack.Peek()

			if top >= digit {
				break
			}

			stack.Pop()
			removalsAllowed--
		}

		stack.Push(digit)
	}

	for stack.Size() > maxDigits {
		stack.Pop()
	}

	joltageDigits := make([]int, stack.Size())
	for i := len(joltageDigits) - 1; i >= 0; i-- {
		v, _ := stack.Pop()
		joltageDigits[i] = v
	}

	var maxJoltage int64
	for _, d := range joltageDigits {
		maxJoltage = maxJoltage*10 + int64(d)
	}

	return maxJoltage
}
