package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/willmadison/advent/advent2018"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%s/input", time.Now().Year(), os.Getenv("DAY")), nil)

	session := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("SESSION_ID"),
		Domain: "adventofcode.com",
	}
	req.AddCookie(session)

	start := time.Now()

	response, err := http.DefaultClient.Do(req)

	fmt.Println("input fetched in", time.Since(start))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		os.Exit(1)
	}

	defer response.Body.Close()

	start = time.Now()
	a := advent2018.FindRegionAreaMinimizedByConstraint(response.Body, 10000)
	fmt.Println("area =", a)
	fmt.Println("area found in", time.Since(start))
}
