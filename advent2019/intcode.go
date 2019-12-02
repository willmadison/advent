package advent2019

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Modifier func([]int) []int

func ReverseEngineerIntCodeMachine(input io.Reader, desiredOutput int) (int, int) {
	program := parseProgram(input)
	var noun int
	var verb int

	for ; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			p := append([]int(nil), program...)
			result := run(p, Modifier(func(values []int) []int {
				values[1] = noun
				values[2] = verb
				return values
			}))

			if result[0] == desiredOutput {
				return noun, verb
			}
		}
	}

	fmt.Println("Uh oh, never able to find a suitor returning max noun and max verb...")

	return noun, verb
}

func RunIntCodeMachine(input io.Reader, modifiers ...Modifier) []int {
	program := parseProgram(input)

	run(program, modifiers...)

	return program
}

func parseProgram(r io.Reader) []int {
	scanner := bufio.NewScanner(r)

	var rawProgram string

	if scanner.Scan() {
		rawProgram = scanner.Text()
	}

	rawProgramParts := strings.Split(rawProgram, ",")

	var program []int

	for _, part := range rawProgramParts {
		i, _ := strconv.Atoi(part)
		program = append(program, i)
	}

	return program
}

type opCode int

var add opCode = 1
var mult opCode = 2
var halt opCode = 99

func run(program []int, modifiers ...Modifier) []int {
	for _, modifier := range modifiers {
		program = modifier(program)
	}

	var pc int

	code := opCode(program[pc])

Begin:
	for code != halt {
		switch code {
		case add:
			x, y, result := program[program[pc+1]], program[program[pc+2]], program[pc+3]
			program[result] = x + y
		case mult:
			a, b, result := program[program[pc+1]], program[program[pc+2]], program[pc+3]
			program[result] = a * b
		case halt:
			break Begin
		}

		pc += 4

		if pc < len(program) {
			code = opCode(program[pc])
		}
	}

	return program
}
