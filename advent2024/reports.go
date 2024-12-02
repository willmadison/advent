package advent2024

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Report []int

func (r Report) IsSafe() bool {
	reversed := make(Report, len(r))
	copy(reversed, r)

	slices.Reverse(reversed)

	if !slices.IsSorted(r) && !slices.IsSorted(reversed) {
		return false
	}

	minDifference := 1
	maxDifference := 3

	for i := 0; i <= len(r)-2; i++ {
		difference := abs(r[i] - r[i+1])

		if difference < minDifference || difference > maxDifference {
			return false
		}
	}

	return true
}

func (r Report) IsSafeWithTolerance() bool {
	if r.IsSafe() {
		return true
	}

	for i := range r {
		partial := slices.Delete(slices.Clone(r), i, i+1)

		if partial.IsSafe() {
			return true
		}
	}

	return false
}

func ParseReports(r io.Reader) ([]Report, error) {
	var reports []Report

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		var report Report

		for _, v := range values {
			i, _ := strconv.Atoi(v)
			report = append(report, i)
		}

		reports = append(reports, report)
	}

	return reports, nil
}
