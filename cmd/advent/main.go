package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/willmadison/advent/advent2020"
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

	answer := advent2020.RepairReport(response)
	fmt.Println(answer)
}
