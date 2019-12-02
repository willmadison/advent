package main

import (
	"fmt"
	"os"

	"github.com/willmadison/advent"
	"github.com/willmadison/advent/advent2019"
)

func main() {
	input, err := advent.FetchTestInput(os.Getenv("DAY"), os.Getenv("SESSION_ID"))

	if err != nil {
		os.Exit(1)
	}

	defer input.Close()

	//	finalProgram := advent2019.RunIntCodeMachine(input, advent2019.Modifier(func(values []int) []int {
	//		values[1] = 12
	//		values[2] = 2
	//		return values
	//	}))
	//
	//	fmt.Println("program[0] =", finalProgram[0])

	noun, verb := advent2019.ReverseEngineerIntCodeMachine(input, 19690720)

	fmt.Println("100*noun + verb =", 100*noun+verb)
}
