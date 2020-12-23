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

	cups := advent2020.ParseCups(response)

	// Part B.
	cups.AddAdditionalCups(1000000 - cups.Cups.Len())

	for i := 0; i < 10_000_000; i++ {
		cups.Move()
	}

	first := cups.CupsByID[1]

	a := first.Move(1).Value.(int)
	b := first.Move(2).Value.(int)

	fmt.Println(a * b)
}
