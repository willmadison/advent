package advent2025

import (
	"bufio"
	"io"
	"strconv"

	"github.com/willmadison/advent/internal/intervals"
)

func FindFreshIngredients(r io.Reader) ([]int, error) {
	var freshIngredients []int

	rangeGroup, ingredientIds, err := parseRangeGroupsAndIds(r)
	if err != nil {
		return nil, err
	}

	for _, id := range ingredientIds {
		if rangeGroup.Contains(int64(id)) {
			freshIngredients = append(freshIngredients, id)
		}
	}

	return freshIngredients, nil
}

func parseRangeGroupsAndIds(r io.Reader) (intervals.RangeGroup, []int, error) {
	scanner := bufio.NewScanner(r)

	parsingRanges := true

	var rangeGroup intervals.RangeGroup
	var ingredientIds []int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			rng := intervals.ParseRange(line)
			rangeGroup = append(rangeGroup, rng)
		} else {
			ingredientId, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}

			ingredientIds = append(ingredientIds, ingredientId)
		}
	}

	return rangeGroup, ingredientIds, nil
}

func EnumerateFreshIngredientIds(r io.Reader) (int, error) {
	rangeGroup, _, err := parseRangeGroupsAndIds(r)
	if err != nil {
		return 0, err
	}

	return rangeGroup.Size(), nil
}
