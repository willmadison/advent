package main

import (
	"fmt"
	"os"
	"time"

	"github.com/willmadison/advent/advent2020"
	"github.com/willmadison/advent/internal/problems"
)

func main() {
	response, err := problems.Fetch(time.Now().Year(), os.Getenv("DAY"), os.Getenv("SESSION_ID"))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		os.Exit(1)
	}

	defer response.Close()

	var lastSeatingArrangement advent2020.SeatingArrangement

	seatingArrangement := advent2020.ParseSeatingArrangement(response)

	for !lastSeatingArrangement.Equals(seatingArrangement) {
		seatingArrangement, lastSeatingArrangement = seatingArrangement.RunSeatingCycle(5, advent2020.FirstVisible), seatingArrangement
	}

	fmt.Println(len(seatingArrangement.SeatsByState(advent2020.Occupied)))
}
