package advent2022

import (
	"bufio"
	"errors"
	"io"
)

type item rune

func (i item) priority() int {
	if i >= 'a' && i <= 'z' {
		return int(i-'a') + 1
	} else if i >= 'A' && i <= 'Z' {
		return int(i-'A') + 27
	}

	return 0
}

type rucksack struct {
	compartment1 map[item]struct{}
	compartment2 map[item]struct{}
}

func (r rucksack) findCommonItemInCompartments() (item, error) {
	for it := range r.compartment1 {
		if _, exists := r.compartment2[it]; exists {
			return it, nil
		}
	}

	return 0, errors.New("no common item found in compartments")
}

func FindTotalPriorityFromRucksacks(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var totalPriority int

	for scanner.Scan() {
		line := scanner.Text()

		rucksack, err := parseRucksack(line)
		if err != nil {
			return 0, err
		}

		commonItem, err := rucksack.findCommonItemInCompartments()
		if err != nil {
			return 0, err
		}

		priority := commonItem.priority()

		totalPriority += priority
	}

	return totalPriority, nil
}

func FindTotalPriorityFromBadgesInRucksacks(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var rucksacks []rucksack

	for scanner.Scan() {
		line := scanner.Text()

		rucksack, err := parseRucksack(line)
		if err != nil {
			return 0, err
		}

		rucksacks = append(rucksacks, rucksack)
	}

	var totalPriority int

	for i := 0; i < len(rucksacks); i += 3 {
		if i+2 >= len(rucksacks) {
			break
		}

		badgeItem, err := findCommonItemInRucksackGroup(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		if err != nil {
			return 0, err
		}

		priority := badgeItem.priority()

		totalPriority += priority
	}

	return totalPriority, nil
}

func findCommonItemInRucksackGroup(rucksack1, rucksack2, rucksack3 rucksack) (item, error) {
	itemCounts := make(map[item]int)

	updateItemCounts := func(r rucksack) {
		for it := range r.compartment1 {
			itemCounts[it]++
		}
		for it := range r.compartment2 {
			itemCounts[it]++
		}
	}

	updateItemCounts(rucksack1)
	updateItemCounts(rucksack2)
	updateItemCounts(rucksack3)

	for it, count := range itemCounts {
		if count == 3 {
			return it, nil
		}
	}

	return 0, errors.New("no common item found in rucksack group")
}

func parseRucksack(line string) (rucksack, error) {
	mid := len(line) / 2

	compartment1 := make(map[item]struct{})
	compartment2 := make(map[item]struct{})

	for i, char := range line {
		it := item(char)

		if i < mid {
			compartment1[it] = struct{}{}
		} else {
			compartment2[it] = struct{}{}
		}
	}

	return rucksack{
		compartment1: compartment1,
		compartment2: compartment2,
	}, nil
}
