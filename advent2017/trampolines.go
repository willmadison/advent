package advent2017

import (
	"bufio"
	"strings"
	"strconv"
)

func parseJumpInstructions(rawInstructions string) []int {
	var instructions []int

	scanner := bufio.NewScanner(strings.NewReader(rawInstructions))

	for scanner.Scan() {
		rawInstruction := scanner.Text()

		instruction, _ := strconv.Atoi(rawInstruction)
		instructions = append(instructions, instruction)
	}

	return instructions
}

func JumpIterations(instructions []int) int {
	var lastLocation, currentLocation, iterations int

	for currentLocation < len(instructions) {
		lastLocation = currentLocation

		offset := instructions[currentLocation]

		currentLocation += offset

		switch {
		case offset >= 3:
			instructions[lastLocation]--
		default:
			instructions[lastLocation]++
		}

		iterations++
	}

	return iterations
}

func DeriveJumpIterations(instructions string) int {
	return JumpIterations(parseJumpInstructions(instructions))
}
