package advent2023

import (
	"bufio"
	"io"
	"strings"
)

type instruction int

const (
	right instruction = iota
	left
)

var instructionsByLabel = map[rune]instruction{
	'R': right,
	'L': left,
}

type element struct {
	name        string
	left, right string
}

func DetermineTripLength(r io.Reader) int {
	instructions, elementsByName := parse(r)

	var steps int

	var currentInstruction int

	var found bool

	current := elementsByName["AAA"]

	for !found {
		instruction := instructions[currentInstruction%len(instructions)]

		switch instruction {
		case right:
			current = elementsByName[current.right]
		case left:
			current = elementsByName[current.left]
		}

		steps++

		found = current.name == "ZZZ"
		currentInstruction++
	}

	return steps
}

func DetermineGhostlyTripLength(r io.Reader) int {
	instructions, elementsByName := parse(r)

	currentElements := findElementsEndingWith("A", elementsByName)

	var allSteps []int

	for _, e := range currentElements {
		var steps int

		var currentInstruction int

		var finished bool

		for !finished {
			instruction := instructions[currentInstruction%len(instructions)]

			switch instruction {
			case right:
				e = elementsByName[e.right]
			case left:
				e = elementsByName[e.left]
			}

			steps++

			finished = strings.HasSuffix(e.name, "Z")
			currentInstruction++
		}

		allSteps = append(allSteps, steps)
	}

	return lcm(allSteps[0], allSteps[1], allSteps[2:]...)
}

func findElementsEndingWith(suffix string, elementsByName map[string]element) []element {
	var elements []element

	for name, element := range elementsByName {
		if strings.HasSuffix(name, suffix) {
			elements = append(elements, element)
		}
	}

	return elements
}

func parse(r io.Reader) ([]instruction, map[string]element) {
	scanner := bufio.NewScanner(r)

	var lineNumber int

	var instructions []instruction

	elementsByName := map[string]element{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if lineNumber == 0 {
			instructions = parseInstructions(line)
			lineNumber++
		} else {
			e := parseElement(line)
			elementsByName[e.name] = e
		}
	}

	return instructions, elementsByName
}

func parseInstructions(value string) []instruction {
	var instructions []instruction

	for _, c := range value {
		instructions = append(instructions, instructionsByLabel[c])
	}

	return instructions
}

func parseElement(value string) element {
	elementParts := strings.Split(value, " = ")

	name := elementParts[0]

	rawLeftRightParts := strings.TrimPrefix(elementParts[1], "(")
	rawLeftRightParts = strings.TrimSuffix(rawLeftRightParts, ")")
	leftRightParts := strings.Split(rawLeftRightParts, ", ")

	return element{name: name, left: leftRightParts[0], right: leftRightParts[1]}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, values ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(values); i++ {
		result = lcm(result, values[i])
	}

	return result
}
