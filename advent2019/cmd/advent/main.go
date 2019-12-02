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

	totalFuel := advent2019.DeriveTotalFuelRequirementIncludingFuelMass(input)

	fmt.Println("total fuel need (including fuel mass)=", totalFuel)
}
