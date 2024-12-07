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

func (e PartialEquation) OperatorAlternatives() [][]Operator {
	numBits := len(e.Operands) - 1

	start := 0
	end := int(math.Pow(2, float64(numBits)) - 1)

	var alternatives [][]Operator

	for i := start; i <= end; i++ {
		var alternative []Operator

		for bit := numBits - 1; bit >= 0; bit-- {
			bitSet := i&(1<<bit) == 0

			var operator Operator

			if bitSet {
				operator = Times
			} else {
				operator = Plus
			}

			alternative = append(alternative, operator)
		}

		alternatives = append(alternatives, alternative)
	}

	return alternatives
}

func (p PartialEquation) CouldBeMadeTrue() bool {
	alternatives := p.OperatorAlternatives()

	for _, operators := range alternatives {
		if p.ExpectedValue == p.evalute(operators) {
			return true
		}
	}

	return false
}

func (p PartialEquation) evalute(operators []Operator) int {
	answer := p.Operands[0]

	for i := 1; i < len(p.Operands); i++ {
		operator := operators[i-1]
		switch operator {
		case Plus:
			answer += p.Operands[i]
		case Times:
			answer *= p.Operands[i]
		}
	}

	return answer
}

type Operator int

const (
	None Operator = -1
	Plus Operator = iota
	Times
)

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
