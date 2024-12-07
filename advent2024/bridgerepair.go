package advent2024

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type PartialEquation struct {
	ExpectedValue int
	Operands      []int
}

func (e PartialEquation) OperatorAlternatives(operatorsConsidered ...Operator) [][]Operator {
	if len(operatorsConsidered) == 0 {
		operatorsConsidered = []Operator{Plus, Times}
	}

	numBits := len(e.Operands) - 1

	start := 0

	base := len(operatorsConsidered)

	end := int(math.Pow(float64(base), float64(numBits)) - 1)

	var alternatives [][]Operator

	for i := start; i <= end; i++ {
		var alternative []Operator

		for digit := numBits - 1; digit >= 0; digit-- {
			digitValue := getDigit(i, base, digit)
			alternative = append(alternative, operatorsConsidered[digitValue])
		}

		alternatives = append(alternatives, alternative)
	}

	return alternatives
}

func getDigit(value, base, digit int) int {
	asString := strconv.FormatInt(int64(value), base)

	for len(asString) <= digit {
		asString = "0" + asString
	}

	needle, _ := strconv.Atoi(string(asString[len(asString)-digit-1]))

	return needle
}

func (p PartialEquation) CouldBeMadeTrue(ops ...Operator) bool {
	alternatives := p.OperatorAlternatives(ops...)

	for _, operators := range alternatives {
		if p.ExpectedValue == p.evaluate(operators) {
			return true
		}
	}

	return false
}

func (p PartialEquation) evaluate(operators []Operator) int {
	answer := p.Operands[0]

	for i := 1; i < len(p.Operands); i++ {
		operator := operators[i-1]
		switch operator {
		case Plus:
			answer += p.Operands[i]
		case Times:
			answer *= p.Operands[i]
		case Concat:
			var sb strings.Builder

			sb.WriteString(strconv.Itoa(answer))
			sb.WriteString(strconv.Itoa(p.Operands[i]))

			temp := sb.String()

			answer, _ = strconv.Atoi(temp)
		}
	}

	return answer
}

type Operator int

const (
	Plus Operator = iota
	Times
	Concat
)

var AllOperators = []Operator{Plus, Times, Concat}

func ParseCalibrationEquations(r io.Reader) ([]PartialEquation, error) {
	var equations []PartialEquation

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		rawEquationParts := strings.Fields(line)

		equation := PartialEquation{}

		rawExpectedValue := rawEquationParts[0]
		rawExpectedValue = strings.TrimRight(rawExpectedValue, ":")

		equation.ExpectedValue, _ = strconv.Atoi(rawExpectedValue)

		rawOperands := rawEquationParts[1:]

		for _, rawOperand := range rawOperands {
			operand, _ := strconv.Atoi(rawOperand)
			equation.Operands = append(equation.Operands, operand)
		}

		equations = append(equations, equation)
	}

	return equations, nil
}
