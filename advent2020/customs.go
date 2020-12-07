package advent2020

import (
	"bufio"
	"io"
)

type Group struct {
	Affirmatives map[int]map[rune]struct{}
}

func ParseGroups(r io.Reader) []Group {
	var groups []Group

	scanner := bufio.NewScanner(r)

	var groupMember int
	group := Group{
		Affirmatives: map[int]map[rune]struct{}{},
	}

	for scanner.Scan() {
		affirmatives := scanner.Text()

		if len(affirmatives) == 0 {
			groups = append(groups, group)
			group = Group{
				Affirmatives: map[int]map[rune]struct{}{},
			}
			groupMember = 0
			continue
		}

		if _, present := group.Affirmatives[groupMember]; !present {
			group.Affirmatives[groupMember] = map[rune]struct{}{}
		}

		for _, question := range affirmatives {
			group.Affirmatives[groupMember][rune(question)] = struct{}{}
		}

		groupMember++
	}

	groups = append(groups, group)

	return groups
}

func CountAffirmatives(groups []Group) int {
	var count int

	for _, group := range groups {
		yesesByQuestion := map[rune]int{}

		for _, affirmatives := range group.Affirmatives {
			for question := range affirmatives {
				yesesByQuestion[question]++
			}
		}

		for _, yeses := range yesesByQuestion {
			if yeses == len(group.Affirmatives) {
				count++
			}
		}
	}

	return count
}
