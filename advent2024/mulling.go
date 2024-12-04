package advent2024

import (
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Operation string

type Instruction struct {
	Op       Operation
	Operands []int
}

func FindInstructions(r io.Reader) ([]Instruction, error) {
	rawMemory, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	memory := string(rawMemory)

	re := regexp.MustCompile(`(mul\(\d+\,\d+\)|do\(\)|don\'t\(\))`)

	matches := re.FindAllString(memory, -1)

	var instructions []Instruction

	for _, match := range matches {
		instruction := toInstruction(match)

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func toInstruction(value string) Instruction {
	values := strings.Split(value, "(")
	rawOperands := values[1]

	rawOperands = strings.TrimRight(rawOperands, ")")

	rawIndividualOperands := strings.Split(rawOperands, ",")

	var operands []int

	for _, v := range rawIndividualOperands {
		o, _ := strconv.Atoi(v)
		operands = append(operands, o)
	}

	instruction := Instruction{
		Op:       Operation(values[0]),
		Operands: operands,
	}

	return instruction
}

func FindDoables(instructions []Instruction) []Instruction {
	var doableInstructions []Instruction

	doable := true

	for _, i := range instructions {
		switch {
		case i.Op == Operation("do"):
			doable = true
		case i.Op == Operation("don't"):
			doable = false
		default:
			if doable {
				doableInstructions = append(doableInstructions, i)
			}
		}
	}

	return doableInstructions
}
