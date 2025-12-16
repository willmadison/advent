package advent2022

import (
	"bufio"
	"fmt"
	"io"
)

func FindFirstMarkerIndex(r io.Reader, markerLength int) (int, error) {
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return 0, fmt.Errorf("no input found")
	}

	line := scanner.Text()
	window := make([]rune, 0, markerLength)

	for i, char := range line {
		window = append(window, char)
		if len(window) > markerLength {
			window = window[1:]
		}

		if len(window) == markerLength {
			seen := make(map[rune]bool)
			unique := true
			for _, r := range window {
				if seen[r] {
					unique = false
					break
				}
				seen[r] = true
			}
			if unique {
				return i + 1, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("no marker found")
}
