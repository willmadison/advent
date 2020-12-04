package main

import (
	"fmt"
	"os"
	"time"

	"github.com/willmadison/advent/advent2020"
	"github.com/willmadison/advent/internal/location"
	"github.com/willmadison/advent/internal/problems"
)

func main() {
	response, err := problems.Fetch(time.Now().Year(), os.Getenv("DAY"), os.Getenv("SESSION_ID"))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		os.Exit(1)
	}

	defer response.Close()

	tm := advent2020.NewTrajectoryMap(response)
	slope := location.Slope{1, 3}

	fmt.Println(advent2020.CountEncounteredTrees(tm, slope))

	slopes := []location.Slope{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	treesEncountered := []int{}

	for _, slope := range slopes {
		trees := advent2020.CountEncounteredTrees(tm, slope)
		treesEncountered = append(treesEncountered, trees)
	}

	product := 1

	for _, numTrees := range treesEncountered {
		product *= numTrees
	}

	fmt.Println(product)
}
