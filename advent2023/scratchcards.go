package advent2023

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Scratchcard struct {
	Number         int
	Numbers        map[int]struct{}
	WinningNumbers map[int]struct{}
	Value          int
	NumMatches     int
}

func NewScratchCard(number int, numbers, winningNumbers map[int]struct{}) Scratchcard {
	scratchcard := Scratchcard{
		Number:         number,
		Numbers:        numbers,
		WinningNumbers: winningNumbers,
	}

	var value int
	var numMatches int

	for winningNumber := range winningNumbers {
		if _, present := scratchcard.Numbers[winningNumber]; present {
			numMatches++
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}

	scratchcard.Value = value
	scratchcard.NumMatches = numMatches

	return scratchcard
}

func FindWinningScratchcards(r io.Reader) ([]Scratchcard, map[int]int) {
	scanner := bufio.NewScanner(r)

	var cardNumber int
	var scratchcards []Scratchcard

	for scanner.Scan() {
		cardNumber++

		rawScratcherParts := strings.Split(scanner.Text(), " | ")
		rawWinningNumberParts := strings.Split(rawScratcherParts[0], ": ")

		winningNumbers := parseNumbers(rawWinningNumberParts[1])
		numbers := parseNumbers(rawScratcherParts[1])

		scratchcards = append(scratchcards, NewScratchCard(cardNumber, numbers, winningNumbers))
	}

	var winners []Scratchcard

	for _, scratchcard := range scratchcards {
		if scratchcard.Value > 0 {
			winners = append(winners, scratchcard)
		}
	}

	countsByCardNumber := map[int]int{}

	for _, scratchcard := range scratchcards {
		countsByCardNumber[scratchcard.Number]++
	}

	for _, card := range winners {
		multiplier := countsByCardNumber[card.Number]

		for i := card.Number + 1; i < card.Number+card.NumMatches+1; i++ {
			countsByCardNumber[i] += multiplier
		}
	}

	return winners, countsByCardNumber
}

func parseNumbers(spaceDelimitedNumbers string) map[int]struct{} {
	uniqueNumbers := map[int]struct{}{}

	for _, rawNumber := range strings.Fields(spaceDelimitedNumbers) {
		n, _ := strconv.Atoi(rawNumber)
		uniqueNumbers[n] = struct{}{}
	}

	return uniqueNumbers
}
