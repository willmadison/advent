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

	floor := advent2020.ParseAllDirections(response)

	for _, direction := range floor.AllDirections {
		floor.Follow(direction)
	}

	fmt.Println(floor.GetBlackCount())

	for i := 0; i < 100; i++ {
		floor.Rotate()
	}

	fmt.Println(floor.GetBlackCount())
}
