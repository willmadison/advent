package advent2020

import (
	"bufio"
	"io"
	"strconv"
)

func RepairReport(r io.Reader) int {
	expenses := map[int]struct{}{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawExpense := scanner.Text()

		expense, _ := strconv.Atoi(rawExpense)

		expenses[expense] = struct{}{}
	}

	for expense := range expenses {
		complement := 2020 - expense
		if _, present := expenses[complement]; present {
			return complement * expense
		}
	}

	return 0
}

func RepairReportTriplet(r io.Reader) int {
	expenses := map[int]struct{}{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawExpense := scanner.Text()
		expense, _ := strconv.Atoi(rawExpense)
		expenses[expense] = struct{}{}
	}

	for expense := range expenses {
		complement := 2020 - expense
		addends := findTargetAddends(complement, expenses)
		if len(addends) > 0 {
			return expense * addends[0] * addends[1]
		}
	}

	return 0
}

func findTargetAddends(targetSum int, allAddends map[int]struct{}) []int {
	var addends []int

	for addend := range allAddends {
		complement := targetSum - addend
		if _, present := allAddends[complement]; present {
			addends = append(addends, addend, complement)
			break
		}
	}

	return addends
}
