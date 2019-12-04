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

	// var origin advent.Point
	// nearest := advent2019.FindNearestIntersection(input)

	//fmt.Println("distance of nearest intersection =", origin.ManhattanDistance(nearest))

	fmt.Println("Sum of shortest paths:", advent2019.FindMinimalTotalSteps(input))
}
