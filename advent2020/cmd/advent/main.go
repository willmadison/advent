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

	program := advent2020.ParseInitializationProgram(response)

	var c advent2020.Computer
	c.Version = advent2020.Version2

	program.Run(&c)

	var memoryTotal int

	for _, v := range c.Memory {
		memoryTotal += v
	}

	fmt.Println(memoryTotal)
}
