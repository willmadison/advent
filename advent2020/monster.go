package advent2020

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Rule struct {
	ID       int
	Letter   string
	Subrules [][]int
}

type Ruleset map[int]Rule

func (rs Ruleset) FindMatches(ruleNumber int, values []string) int {
	var matches int

	for _, v := range values {
		results := rs.Match(ruleNumber, v)

		for _, result := range results {
			if result == len(v) {
				matches++
				break
			}
		}
	}

	return matches
}

func (rs Ruleset) Match(ruleNumber int, value string) []int {
	r := rs[ruleNumber]

	if len(r.Subrules) == 0 {
		if len(value) < len(r.Letter) {
			return nil
		}

		if value[:len(r.Letter)] == r.Letter {
			return []int{len(r.Letter)}
		}
	}

	var matchedCharacters []int

	for _, subrule := range r.Subrules {
		potentialMatches := []int{0}

		for _, rule := range subrule {
			var newPotentialMatches []int

			for _, match := range potentialMatches {
				matches := rs.Match(rule, value[match:])

				if len(matches) == 0 {
					continue
				}

				for _, m := range matches {
					newPotentialMatches = append(newPotentialMatches, m+match)
				}
			}

			potentialMatches = newPotentialMatches
		}

		matchedCharacters = append(matchedCharacters, potentialMatches...)
	}

	return matchedCharacters

}

func ParseRulesAndMessages(r io.Reader) (Ruleset, []string) {
	messages := []string{}
	rulesByID := map[int]Rule{}

	parsingRules := true

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		t := scanner.Text()

		if t == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rule := parseMessageRule(t)
			rulesByID[rule.ID] = rule
		} else {
			messages = append(messages, t)
		}
	}

	return Ruleset(rulesByID), messages
}

func parseMessageRule(value string) Rule {
	rule := Rule{}

	ruleParts := strings.Split(value, ": ")

	rule.ID, _ = strconv.Atoi(ruleParts[0])
	rule.Subrules = [][]int{}

	if strings.Contains(ruleParts[1], `"`) {
		rule.Letter = strings.Trim(ruleParts[1], `"`)
	} else {
		ruleIDParts := strings.Split(ruleParts[1], " | ")

		for _, part := range ruleIDParts {
			ids := []int{}
			idParts := strings.Split(part, " ")

			for _, rawID := range idParts {
				id, _ := strconv.Atoi(rawID)
				ids = append(ids, id)
			}

			rule.Subrules = append(rule.Subrules, ids)
		}
	}

	return rule
}
