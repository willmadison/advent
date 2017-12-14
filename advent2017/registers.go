package advent2017

import (
	"strings"
	"strconv"
)

type Operation int

const (
	_                   = iota
	Increment Operation = iota
	Decrement
)

type Instruction struct {
	register  string
	operation Operation
	operand   int
	condition Condition
}

type Comparator int

const (
	_   = iota
	GT  = iota
	LT
	EQ
	GTE
	LTE
)

type Condition struct {
	register   string
	comparator Comparator
	value      int
}

var zeroCondition Condition

func (c Condition) IsZero() bool {
	return c == zeroCondition
}

func ParseInstruction(rawInstruction string) (Instruction, error) {
	parts := strings.Fields(rawInstruction)
	instruction := Instruction{
		register:  parts[0],
		operation: asOperation(parts[1]),
	}

	operand, err := strconv.Atoi(parts[2])
	if err != nil {
		return Instruction{}, err
	}

	instruction.operand = operand

	if len(parts) > 3 {
		instruction.condition, err = asCondition(parts[4:])
		if err != nil {
			return Instruction{}, err
		}
	}

	return instruction, nil
}

func asOperation(op string) Operation {
	var operation Operation

	switch op {
	case "inc":
		operation = Increment
	case "dec":
		operation = Decrement
	}

	return operation
}

func asCondition(parts []string) (Condition, error) {
	condition := Condition{
		register:   parts[0],
		comparator: asComparator(parts[1]),
	}

	value, err := strconv.Atoi(parts[2])
	if err != nil {
		return Condition{}, err
	}

	condition.value = value

	return condition, nil
}

func asComparator(comp string) Comparator {
	var comparator Comparator

	switch comp {
	case ">":
		comparator = GT
	case "<":
		comparator = LT
	case "==":
		comparator = EQ
	case "<=":
		comparator = LTE
	case ">=":
		comparator = GTE
	}

	return comparator
}
