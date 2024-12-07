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

	equations, _ := advent2024.ParseCalibrationEquations(response)

	var answer int

	for _, e := range equations {
		if e.CouldBeMadeTrue() {
			answer += e.ExpectedValue
		}
	}

	color.Green("=================")
	color.Green("%v", answer)
	color.Green("=================")

	fmt.Println("solved in", time.Since(start))
}
