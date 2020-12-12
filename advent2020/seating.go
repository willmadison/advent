package advent2020

import (
	"bufio"
	"io"
	"sync"

	"github.com/willmadison/advent/internal/location"
)

type NeighboringStrategy func(location.Coordinate, map[location.Coordinate]Seat) []Seat

type SeatingArrangement []Seat

var NearestNeighbor = NeighboringStrategy(func(locale location.Coordinate, seatsByLocation map[location.Coordinate]Seat) []Seat {
	var neighbors []Seat

	neighboringLocations := []location.Coordinate{
		{locale.Row - 1, locale.Col - 1},
		{locale.Row - 1, locale.Col},
		{locale.Row - 1, locale.Col + 1},
		{locale.Row, locale.Col - 1},
		{locale.Row, locale.Col + 1},
		{locale.Row + 1, locale.Col - 1},
		{locale.Row + 1, locale.Col},
		{locale.Row + 1, locale.Col + 1},
	}

	for _, l := range neighboringLocations {
		if seat, present := seatsByLocation[l]; present {
			neighbors = append(neighbors, seat)
		}
	}

	return neighbors
})

var FirstVisible = NeighboringStrategy(func(locale location.Coordinate, seatsByLocation map[location.Coordinate]Seat) []Seat {
	var maxRow, maxCol int

	for l := range seatsByLocation {
		if l.Row > maxRow {
			maxRow = l.Row
		}

		if l.Col > maxCol {
			maxCol = l.Col
		}
	}

	type modifier func(location.Coordinate) location.Coordinate

	upwardLeftDiagonal := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row - 1, l.Col - 1}
	})
	upward := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row - 1, l.Col}
	})
	upwardRightDiagonal := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row - 1, l.Col + 1}
	})
	left := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row, l.Col - 1}
	})
	right := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row, l.Col + 1}
	})
	downwardLeftDiagonal := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row + 1, l.Col - 1}
	})
	down := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row + 1, l.Col}
	})
	downwardRightDiagonal := modifier(func(l location.Coordinate) location.Coordinate {
		return location.Coordinate{l.Row + 1, l.Col + 1}
	})

	modifiers := []modifier{
		upwardLeftDiagonal,
		upward,
		upwardRightDiagonal,
		left,
		right,
		downwardLeftDiagonal,
		down,
		downwardRightDiagonal,
	}

	visibleNeighbors := make(chan Seat, len(modifiers))
	var wg sync.WaitGroup

	for _, m := range modifiers {
		wg.Add(1)
		go func(origin location.Coordinate, mod modifier, out chan<- Seat, waitGroup *sync.WaitGroup) {
			current := location.Coordinate{origin.Row, origin.Col}

			for {
				current = mod(current)

				if seat, present := seatsByLocation[current]; present {
					out <- seat
					break
				}

				if current.Row < 0 || current.Col < 0 || current.Row > maxRow || current.Col > maxCol {
					break
				}
			}

			waitGroup.Done()
		}(locale, m, visibleNeighbors, &wg)
	}

	wg.Wait()
	close(visibleNeighbors)

	collected := map[location.Coordinate]struct{}{}

	neighbors := []Seat{}

	for neighbor := range visibleNeighbors {
		if _, seen := collected[neighbor.Location]; !seen {
			neighbors = append(neighbors, neighbor)
			collected[neighbor.Location] = struct{}{}
		}
	}

	return neighbors
})

func (s SeatingArrangement) RunSeatingCycle(occupancyThreshold int, strategies ...NeighboringStrategy) SeatingArrangement {
	strategy := NearestNeighbor

	if len(strategies) != 0 {
		strategy = strategies[0]
	}

	seatsByLocation := map[location.Coordinate]Seat{}
	occupiedSeats := map[location.Coordinate]Seat{}
	emptySeats := map[location.Coordinate]Seat{}

	shouldBeOccupied := []location.Coordinate{}
	shouldBeEmptied := []location.Coordinate{}

	for _, seat := range s {
		seatsByLocation[seat.Location] = seat

		if seat.Status == Occupied {
			occupiedSeats[seat.Location] = seat
		} else {
			emptySeats[seat.Location] = seat
		}
	}

	for _, seat := range s {
		neighbors := seat.GetNeighbors(strategy, seatsByLocation)

		var occupiedNeighbors int

		for _, n := range neighbors {
			if n.Status == Occupied {
				occupiedNeighbors++
			}
		}

		if occupiedNeighbors == 0 {
			shouldBeOccupied = append(shouldBeOccupied, seat.Location)
		} else if occupiedNeighbors >= occupancyThreshold && seat.Status == Occupied {
			shouldBeEmptied = append(shouldBeEmptied, seat.Location)
		}
	}

	for _, l := range shouldBeOccupied {
		seat := seatsByLocation[l]
		seat.Status = Occupied
		seatsByLocation[l] = seat
	}

	for _, l := range shouldBeEmptied {
		seat := seatsByLocation[l]
		seat.Status = Empty
		seatsByLocation[l] = seat
	}

	var resultantSeats []Seat

	for _, seat := range seatsByLocation {
		resultantSeats = append(resultantSeats, seat)
	}

	return SeatingArrangement(resultantSeats)
}

func (s SeatingArrangement) Equals(other SeatingArrangement) bool {
	if len(s) != len(other) {
		return false
	}

	otherSeatsByLocation := map[location.Coordinate]Seat{}

	for _, otherSeat := range other {
		otherSeatsByLocation[otherSeat.Location] = otherSeat
	}

	for _, mySeat := range s {
		if seat, present := otherSeatsByLocation[mySeat.Location]; !present || seat.Status != mySeat.Status {
			return false
		}
	}

	return true
}

func (s SeatingArrangement) SeatsByState(state State) []Seat {
	seatsByState := map[State][]Seat{}

	for _, seat := range s {
		seatsByState[seat.Status] = append(seatsByState[seat.Status], seat)
	}

	if _, present := seatsByState[state]; !present {
		return []Seat{}
	}

	return seatsByState[state]
}

type State int

const (
	_ State = iota
	Occupied
	Empty
)

type Seat struct {
	Location location.Coordinate
	Status   State
}

func (s Seat) GetNeighbors(strategy NeighboringStrategy, seatsByLocation map[location.Coordinate]Seat) []Seat {
	return strategy(s.Location, seatsByLocation)
}

func ParseSeatingArrangement(r io.Reader) SeatingArrangement {
	var seats []Seat

	var row int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawRow := scanner.Text()

		for col, v := range rawRow {
			value := rune(v)

			if value == '.' {
				continue
			}

			var status State

			switch value {
			case 'L':
				status = Empty
			case '#':
				status = Occupied
			}

			seat := Seat{
				location.Coordinate{row, col},
				status,
			}

			seats = append(seats, seat)
		}

		row++
	}

	return SeatingArrangement(seats)
}
