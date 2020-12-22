package main

import (
	"fmt"
	"os"
	"strings"
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

	menu := advent2020.ParseFoodListing(response)

	nonAllergenics, allergicIngredients := menu.FindNonAllergenicIngredients()

	var sum int

	for _, occurrences := range menu.CountOccurrencesFor(nonAllergenics) {
		sum += occurrences
	}

	fmt.Println(sum)
	fmt.Println(strings.Join(allergicIngredients, ","))
}
