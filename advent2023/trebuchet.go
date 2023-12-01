package advent2023

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func ProcessCalibrationDocument(r io.Reader) int {
	var sum int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		value := extractCalibrationValue(scanner.Text())
		sum += value
	}

	return sum
}

func extractCalibrationValue(line string) int {
	var digits []rune

	replacementValuesByDigitWord := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for location, character := range line {
		if unicode.IsDigit(character) {
			digits = append(digits, character)
		}

		for word, replacement := range replacementValuesByDigitWord {
			if strings.HasPrefix(line[location:], word) {
				digits = append(digits, replacement)
			}
		}
	}

	var buf bytes.Buffer

	buf.WriteRune(digits[0])
	buf.WriteRune(digits[len(digits)-1])

	value, _ := strconv.Atoi(buf.String())

	return value
}
