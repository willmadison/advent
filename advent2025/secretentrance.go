package advent2025

import (
	"bufio"
	"io"
	"strconv"
)

type Dial struct {
	Value int
}

func (d *Dial) Turn(direction rune, ticks int) int {
	var numPassesThroughZero int

	var delta int

	switch direction {
	case 'L':
		delta = -1
	case 'R':
		delta = 1
	}

	iterations := ticks

	for iterations > 0 {
		d.Value += delta

		if d.Value < 0 {
			d.Value = 99
		}

		if d.Value > 99 {
			d.Value = 0
		}

		if d.Value == 0 {
			numPassesThroughZero++
		}

		iterations--
	}

	return numPassesThroughZero
}

func NewDial() *Dial {
	return &Dial{Value: 50}
}

func CrackPassword(r io.Reader) (int, error) {
	var password int

	scanner := bufio.NewScanner(r)

	dial := NewDial()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		magnitude, err := strconv.Atoi(line[1:])

		if err != nil {
			return 0, err
		}

		dial.Turn(direction, magnitude)

		if dial.Value == 0 {
			password++
		}
	}

	return password, nil
}

func CrackPasswordV2(r io.Reader) (int, error) {
	var password int

	scanner := bufio.NewScanner(r)

	dial := NewDial()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		magnitude, err := strconv.Atoi(line[1:])

		if err != nil {
			return 0, err
		}

		var numPassesThroughZero int

		numPassesThroughZero += dial.Turn(direction, magnitude)

		password += numPassesThroughZero
	}

	return password, nil
}
