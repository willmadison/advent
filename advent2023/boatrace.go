package advent2023

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

type RaceStrategy struct {
	Winners        int
	RecordDistance int
	TimeAlloted    int
}

type Winner struct {
	ChargingTime     int
	Rate             int
	DistanceTraveled int
}

type Kerning int

const (
	Bad Kerning = iota
	Good
)

func (r *RaceStrategy) DetermineWinners() {
	var minChargingTime int

	for t := 0; t < r.RecordDistance; t++ {
		rate := t
		runtime := r.TimeAlloted - t
		d := rate * runtime

		if d > r.RecordDistance {
			minChargingTime = t
			break
		}
	}

	r.Winners = r.TimeAlloted - 2*minChargingTime + 1
}

func FindWinningRaceStrategies(r io.Reader, kernings ...Kerning) []*RaceStrategy {
	scanner := bufio.NewScanner(r)

	times := []int{}
	distances := []int{}

	var kerning Kerning

	if len(kernings) == 0 {
		kerning = Bad
	} else {
		kerning = kernings[0]
	}

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "Time:"):
			line = strings.TrimPrefix(line, "Time:")
			rawTimes := strings.Fields(line)

			switch kerning {
			case Bad:
				for _, rawTime := range rawTimes {
					t, _ := strconv.Atoi(rawTime)
					times = append(times, t)
				}
			case Good:
				var buf bytes.Buffer

				for _, rawTime := range rawTimes {
					buf.WriteString(rawTime)
				}

				t, _ := strconv.Atoi(buf.String())
				times = append(times, t)
			}
		case strings.HasPrefix(line, "Distance:"):
			line = strings.TrimPrefix(line, "Distance:")
			rawDistances := strings.Fields(line)

			switch kerning {
			case Bad:
				for _, rawDistance := range rawDistances {
					d, _ := strconv.Atoi(rawDistance)
					distances = append(distances, d)
				}
			case Good:
				var buf bytes.Buffer

				for _, rawDistance := range rawDistances {
					buf.WriteString(rawDistance)
				}

				d, _ := strconv.Atoi(buf.String())
				distances = append(distances, d)
			}
		}
	}

	strategies := []*RaceStrategy{}

	for i, time := range times {
		strategies = append(strategies, &RaceStrategy{
			RecordDistance: distances[i],
			TimeAlloted:    time,
		})
	}

	for _, strategy := range strategies {
		strategy.DetermineWinners()
	}

	return strategies
}
