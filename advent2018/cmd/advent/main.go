package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/willmadison/advent/advent2018"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2018/day/%s/input", os.Getenv("DAY")), nil)

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

	content, _ := ioutil.ReadAll(response.Body)

	start = time.Now()
	reduced := advent2018.OptimalReduction(string(bytes.TrimSpace(content)))
	fmt.Println("reduction completed in", time.Since(start))

	fmt.Println(reduced)
	fmt.Println(len(reduced))
}
