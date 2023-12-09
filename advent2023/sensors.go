package advent2023

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Reading []int

func (r Reading) extrapolateNextValues() (int, int) {
	countsByValue := map[int]int{}

	for _, v := range r {
		countsByValue[v]++
	}

	if len(countsByValue) == 1 {
		return r[0], r[0]
	} else {
		var deltas Reading

		for i := 1; i < len(r); i++ {
			delta := r[i] - r[i-1]
			deltas = append(deltas, delta)
		}

		prior, next := deltas.extrapolateNextValues()

		return r[0] - prior, r[len(r)-1] + next
	}
}

func ExtrapolateSenorReadings(r io.Reader) []Reading {
	scanner := bufio.NewScanner(r)

	var readings []Reading

	for scanner.Scan() {
		rawReadings := strings.Fields(scanner.Text())

		var reading Reading

		for _, rawReading := range rawReadings {
			i, _ := strconv.Atoi(rawReading)
			reading = append(reading, i)
		}

		readings = append(readings, reading)
	}

	var extrapolatedReadings []Reading

	for _, reading := range readings {
		prior, next := reading.extrapolateNextValues()
		r := []int{prior}
		r = append(r, reading...)
		r = append(r, next)
		extrapolatedReadings = append(extrapolatedReadings, r)
	}

	return extrapolatedReadings
}
