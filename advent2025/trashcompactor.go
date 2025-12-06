package advent2025

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type MathOperandVariant int

const (
	Human MathOperandVariant = iota
	Cephalopod
)

type Expression struct {
	Operator rune
	Operands []int
}

func CheckWorksheet(input io.Reader, variant MathOperandVariant) ([]int64, error) {
	scanner := bufio.NewScanner(input)

	terms := [][]int{}
	operators := []rune{}
	rawLines := []string{}

	var operatorLine string

	for scanner.Scan() {
		line := scanner.Text()

		rawValues := strings.Fields(line)

		if unicode.IsDigit(rune(rawValues[0][0])) {
			rawLines = append(rawLines, line)

			var subterms []int

			for _, rawValue := range rawValues {
				value, _ := strconv.Atoi(rawValue)
				subterms = append(subterms, value)
			}

			terms = append(terms, subterms)
		} else {
			operatorLine = line

			for _, rawValue := range rawValues {
				operators = append(operators, rune(rawValue[0]))
			}
		}
	}

	var subtotals []int64

	if variant == Cephalopod {
		expressions := reinterpretCephalopodTerms(rawLines, operatorLine)

		for _, expression := range expressions {
			var total int64

			switch expression.Operator {
			case '+':
				for _, operand := range expression.Operands {
					total += int64(operand)
				}
			case '*':
				total = 1
				for _, operand := range expression.Operands {
					total *= int64(operand)
				}
			}

			subtotals = append(subtotals, total)
		}
	} else {
		subtotals = make([]int64, len(terms[0]))

		for col, operator := range operators {
			var total int64

			switch operator {
			case '+':
				for row := 0; row < len(terms); row++ {
					total += int64(terms[row][col])
				}
			case '*':
				total = 1
				for row := 0; row < len(terms); row++ {
					total *= int64(terms[row][col])
				}
			}

			subtotals[col] = total
		}
	}

	return subtotals, nil
}

func reinterpretCephalopodTerms(rawLines []string, operatorLine string) []Expression {
	maxLength := len(operatorLine)

	for _, line := range rawLines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	for i, line := range rawLines {
		if len(line) < maxLength {
			rawLines[i] = padRight(line, maxLength)
		}
	}

	if len(operatorLine) < maxLength {
		operatorLine = padRight(operatorLine, maxLength)
	}

	columnValues := make([][]int, maxLength)

	for c := 0; c < maxLength; c++ {
		inNumber := false
		val := 0

		for _, line := range rawLines {
			ch := line[c]

			if ch >= '0' && ch <= '9' {
				inNumber = true
				val = val*10 + int(ch-'0')
			} else {
				if inNumber {
					columnValues[c] = append(columnValues[c], val)
					val = 0
					inNumber = false
				}
			}
		}
		if inNumber {
			columnValues[c] = append(columnValues[c], val)
		}
	}

	var expressions []Expression

	for i := 0; i < maxLength; {
		r := rune(operatorLine[i])

		if r == ' ' {
			i++
			continue
		}

		end := maxLength
		for j := i + 1; j < maxLength; j++ {
			if operatorLine[j] != ' ' {
				end = j
				break
			}
		}

		expression := Expression{Operator: r}

		for c := i; c < end; c++ {
			expression.Operands = append(expression.Operands, columnValues[c]...)
		}

		expressions = append(expressions, expression)
		i = end
	}

	return expressions
}

func padRight(s string, n int) string {
	if len(s) >= n {
		return s
	}
	b := make([]byte, n)
	copy(b, s)
	for i := len(s); i < n; i++ {
		b[i] = ' '
	}
	return string(b)
}
