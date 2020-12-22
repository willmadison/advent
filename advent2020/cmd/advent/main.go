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

	imageTileset := advent2020.ParseImageTiles(response)

	product := 1

	edgeTileIDs := imageTileset.FindCorners()

	for _, tileID := range edgeTileIDs {
		product *= tileID
	}

	fmt.Println(product)
}
