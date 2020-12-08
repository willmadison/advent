package advent2020

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Operation string

const (
	NoOp       Operation = "nop"
	Accumulate Operation = "acc"
	Jump       Operation = "jmp"
)

type Command struct {
	op       Operation
	argument int
}

func toCommand(instruction string) Command {
	instructionParts := strings.Split(instruction, " ")
	argument, _ := strconv.Atoi(instructionParts[1])
	return Command{Operation(instructionParts[0]), argument}
}

func ParseProgram(r io.Reader) []Command {
	scanner := bufio.NewScanner(r)

	var commands []Command

	for scanner.Scan() {
		commands = append(commands, toCommand(scanner.Text()))
	}

	return commands
}

func DetermineAccumulatorValueBeforeLoop(commands []Command) int {
	var accumulator int

	commandsExecuted := map[int]struct{}{}

	var pc int

	for pc < len(commands) {
		command := commands[pc]

		if _, executed := commandsExecuted[pc]; executed {
			break
		}

		commandsExecuted[pc] = struct{}{}

		switch command.op {
		case Jump:
			pc += command.argument
		case Accumulate:
			accumulator += command.argument
			fallthrough
		default:
			pc++
		}
	}

	return accumulator
}

func PatchProgram(commands []Command) []Command {
	jumpCommands := map[int]struct{}{}
	noOpCommands := map[int]struct{}{}

	for i, c := range commands {
		switch c.op {
		case Jump:
			jumpCommands[i] = struct{}{}
		case NoOp:
			noOpCommands[i] = struct{}{}
		}
	}

	for jumpCommand := range jumpCommands {
		newCommands := make([]Command, len(commands))
		copy(newCommands, commands)

		newCommands[jumpCommand].op = NoOp

		if !hasCycle(newCommands) {
			return newCommands
		}
	}

	for noOpCommand := range noOpCommands {
		newCommands := make([]Command, len(commands))
		copy(newCommands, commands)

		newCommands[noOpCommand].op = Jump

		if !hasCycle(newCommands) {
			return newCommands
		}
	}

	return nil
}

func hasCycle(commands []Command) bool {
	commandsExecuted := map[int]struct{}{}

	var pc int

	for pc < len(commands) {
		command := commands[pc]

		if _, executed := commandsExecuted[pc]; executed {
			return true
		}

		commandsExecuted[pc] = struct{}{}

		switch command.op {
		case Jump:
			pc += command.argument
		default:
			pc++
		}
	}

	return false
}

func Run(commands []Command) int {
	var accumulator, pc int

	for pc < len(commands) {
		command := commands[pc]

		switch command.op {
		case Jump:
			pc += command.argument
		case Accumulate:
			accumulator += command.argument
			fallthrough
		default:
			pc++
		}
	}

	return accumulator
}
