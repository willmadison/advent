package advent2022

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/containers"
)

type Crate rune

type MoveInstruction struct {
	Amount    int
	FromStack int
	ToStack   int
}

type MoveStrategy int

const (
	MoveStrategyOneByOne MoveStrategy = iota
	MoveStrategyMultipleAtOnce
)

func FindTerminalTopCrates(r io.Reader, strategy MoveStrategy) (string, error) {
	scanner := bufio.NewScanner(r)

	var crateLines []string
	var moves []MoveInstruction

	parsingCrates := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingCrates = false
			continue
		}

		if parsingCrates {
			crateLines = append(crateLines, line)
		} else if strings.HasPrefix(line, "move") {
			var amount, from, to int
			fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
			moves = append(moves, MoveInstruction{
				Amount:    amount,
				FromStack: from,
				ToStack:   to,
			})
		}
	}

	if len(crateLines) == 0 {
		return "", fmt.Errorf("no crate lines found")
	}

	stackNumberLine := crateLines[len(crateLines)-1]
	numStacks := 0
	for _, char := range stackNumberLine {
		if char >= '1' && char <= '9' {
			stackNum, _ := strconv.Atoi(string(char))
			if stackNum > numStacks {
				numStacks = stackNum
			}
		}
	}

	crateStacks := make([]containers.Stack[Crate], numStacks)
	for i := 0; i < numStacks; i++ {
		crateStacks[i] = containers.NewStack[Crate]()
	}

	for i := len(crateLines) - 2; i >= 0; i-- {
		line := crateLines[i]

		for stackIndex := 0; stackIndex < numStacks; stackIndex++ {
			charPos := 1 + (stackIndex * 4)
			if charPos < len(line) {
				char := line[charPos]
				if char != ' ' {
					crateStacks[stackIndex].Push(Crate(char))
				}
			}
		}
	}

	for _, move := range moves {
		if strategy == MoveStrategyOneByOne {
			for i := 0; i < move.Amount; i++ {
				if crateStacks[move.FromStack-1].Size() > 0 {
					crate, _ := crateStacks[move.FromStack-1].Pop()
					crateStacks[move.ToStack-1].Push(crate)
				}
			}
		} else {
			var tempCrates []Crate
			for i := 0; i < move.Amount; i++ {
				if crateStacks[move.FromStack-1].Size() > 0 {
					crate, _ := crateStacks[move.FromStack-1].Pop()
					tempCrates = append(tempCrates, crate)
				}
			}
			for i := len(tempCrates) - 1; i >= 0; i-- {
				crateStacks[move.ToStack-1].Push(tempCrates[i])
			}
		}
	}

	var result strings.Builder
	for i := 0; i < numStacks; i++ {
		if crateStacks[i].Size() > 0 {
			crate, _ := crateStacks[i].Peek()
			result.WriteRune(rune(crate))
		}
	}

	return result.String(), nil
}
