package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/willmadison/advent/advent2024"
	"github.com/willmadison/advent/internal/problems"
)

func main() {
	year := time.Now().Year()

	if os.Getenv("YEAR") != "" {
		year, _ = strconv.Atoi(os.Getenv("YEAR"))
	}

	response, err := problems.Fetch(year, os.Getenv("DAY"), os.Getenv("SESSION_ID"))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		os.Exit(1)
	}

	defer response.Close()

	start := time.Now()

	instructions, _ := advent2024.FindInstructions(response)

	var answer int

	doableInstructions := advent2024.FindDoables(instructions)

	for _, instruction := range doableInstructions {
		answer += instruction.Operands[0] * instruction.Operands[1]
	}

	color.Green("=================")
	color.Green("%v", answer)
	color.Green("=================")

	fmt.Println("solved in", time.Since(start))
}
