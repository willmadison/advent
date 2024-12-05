package advent2024

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Ruleset map[int]map[int]struct{}

func (r Ruleset) GetPrecedents(page int) []int {
	precedents := []int{}

	if parents, present := r[page]; present {
		for parent := range parents {
			precedents = append(precedents, parent)
		}
	}

	return precedents
}

func (r Ruleset) Compare(pageA, pageB int) int {
	aPrecedents := r.GetPrecedents(pageA)

	lookup := map[int]struct{}{}

	for _, precedent := range aPrecedents {
		lookup[precedent] = struct{}{}
	}

	if _, present := lookup[pageB]; present {
		return 1
	}

	bPrecedents := r.GetPrecedents(pageB)

	lookup = map[int]struct{}{}

	for _, precedent := range bPrecedents {
		lookup[precedent] = struct{}{}
	}

	if _, present := lookup[pageA]; present {
		return -1
	}

	return 0
}

func ParseRulesAndPages(r io.Reader) (Ruleset, [][]int, error) {
	parsingRules := true

	scanner := bufio.NewScanner(r)

	ruleset := map[int]map[int]struct{}{}
	pagelists := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rawPageNumbers := strings.Split(line, "|")
			rawParentPage := rawPageNumbers[0]
			rawChildPage := rawPageNumbers[1]

			parent, _ := strconv.Atoi(rawParentPage)
			child, _ := strconv.Atoi(rawChildPage)

			if _, present := ruleset[child]; !present {
				ruleset[child] = map[int]struct{}{}
			}

			ruleset[child][parent] = struct{}{}
		} else {
			rawPageNumbers := strings.Split(line, ",")

			pages := []int{}

			for _, n := range rawPageNumbers {
				page, _ := strconv.Atoi(n)
				pages = append(pages, page)
			}

			pagelists = append(pagelists, pages)
		}
	}

	return ruleset, pagelists, nil
}

func CanPrint(ruleset Ruleset, pages []int) bool {
	printedPages := map[int]struct{}{}
	requestedPages := map[int]struct{}{}

	for _, page := range pages {
		requestedPages[page] = struct{}{}
	}

	for _, page := range pages {
		precedants := ruleset.GetPrecedents(page)

		for _, preprecedant := range precedants {
			if _, present := requestedPages[preprecedant]; present {
				if _, printed := printedPages[preprecedant]; !printed {
					return false
				}
			}
		}

		printedPages[page] = struct{}{}
	}

	return true
}

func CorrectPrintRequest(ruleset Ruleset, pages []int) []int {
	slices.SortStableFunc(pages, func(a, b int) int {
		return ruleset.Compare(a, b)
	})

	return pages
}
