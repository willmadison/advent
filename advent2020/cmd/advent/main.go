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

	var ship advent2020.Ship
	instructions := advent2020.ParseNavigationInstructions(response)
	ship.NavigateByWaypoint(instructions)

	var origin location.Point

	fmt.Println(origin.ManhattanDistance(ship.Location))
}
