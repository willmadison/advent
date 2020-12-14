package advent2020

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type Itinerary struct {
	EarliestDeparture int
	Busses            []Bus
}

type Bus struct {
	ID, ArrivalOffset int
}

func ParseShuttleItinerary(r io.Reader) Itinerary {
	scanner := bufio.NewScanner(r)

	var itinerary Itinerary

	var lineNumber int

	for scanner.Scan() {
		line := scanner.Text()

		if lineNumber == 0 {
			itinerary.EarliestDeparture, _ = strconv.Atoi(line)
		} else {
			rawBusIds := strings.Split(line, ",")

			var arrivalOffset int

			for _, rawBusID := range rawBusIds {
				if rawBusID == "x" {
					arrivalOffset++
					continue
				}

				busID, _ := strconv.Atoi(rawBusID)
				bus := Bus{busID, arrivalOffset}
				arrivalOffset++
				itinerary.Busses = append(itinerary.Busses, bus)
			}
		}

		lineNumber++
	}

	return itinerary
}

type arrival struct {
	busID, arrivalTime int
}

func FindEarliestBus(i Itinerary) (int, int) {
	bussesByWaitTime := map[int]int{}

	smallestWaitTime := math.MaxInt64

	for _, bus := range i.Busses {
		waitTime := bus.ID - (i.EarliestDeparture % bus.ID)

		if waitTime < smallestWaitTime {
			smallestWaitTime = waitTime
		}

		bussesByWaitTime[waitTime] = bus.ID
	}

	return bussesByWaitTime[smallestWaitTime], smallestWaitTime

}

func FindEarliestTimestampWithDepartureCadence(i Itinerary) int {
	earliestArrival := 0
	runningProduct := 1

	for _, bus := range i.Busses {
		for (earliestArrival+bus.ArrivalOffset)%bus.ID != 0 {
			earliestArrival += runningProduct
		}
		runningProduct *= bus.ID
	}

	return earliestArrival
}
