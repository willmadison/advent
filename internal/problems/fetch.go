package problems

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetch(year int, day, sessionID string) (io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%s/input", year, day), nil)

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
		return nil, err
	}

	return response.Body, nil
}
