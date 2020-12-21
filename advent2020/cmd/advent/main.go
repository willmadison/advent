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

	ruleset, messages := advent2020.ParseRulesAndMessages(response)

	ruleset[8] = advent2020.Rule{
		ID: 8,
		Subrules: [][]int{
			{42},
			{42, 8},
		},
	}

	ruleset[11] = advent2020.Rule{
		ID: 11,
		Subrules: [][]int{
			{42, 31},
			{42, 11, 31},
		},
	}

	fmt.Println(ruleset.FindMatches(0, messages))
}
