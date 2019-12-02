package advent

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchTestInput(day, sessionID string) (io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%s/input", time.Now().Year(), day), nil)

	session := &http.Cookie{
		Name:   "session",
		Value:  sessionID,
		Domain: "adventofcode.com",
	}
	req.AddCookie(session)

	start := time.Now()

	response, err := http.DefaultClient.Do(req)

	fmt.Println("input fetched in", time.Since(start))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", day, err)
		return nil, err
	}

	return response.Body, nil
}
