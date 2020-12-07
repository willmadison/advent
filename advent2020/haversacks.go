package advent2020

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type BagRules map[string]map[string]int

type Queue interface {
	Enqueue(string)
	Dequeue() (string, error)
	Size() int
}

type stringQueue struct {
	data []string
	size int
}

func NewStringQueue() Queue {
	return &stringQueue{data: []string{}}
}

func (s *stringQueue) Enqueue(value string) {
	s.data = append(s.data, value)
	s.size++
}

func (s *stringQueue) Dequeue() (string, error) {
	if s.size > 0 {
		value := s.data[0]
		s.size--
		s.data = s.data[1:]

		return value, nil
	}

	return "", errors.New("no such element")
}

func (s *stringQueue) Peek() (string, error) {
	if s.size > 0 {
		value := s.data[0]
		return value, nil
	}

	return "", errors.New("no such element")
}

func (s stringQueue) Size() int {
	return s.size
}

func (b BagRules) FindAncestorsOf(bag string) []string {
	q := NewStringQueue()
	q.Enqueue(bag)

	uniqueAncestors := map[string]struct{}{}

	visited := map[string]struct{}{}

	for q.Size() > 0 {
		c, _ := q.Dequeue()

		if _, seen := visited[c]; seen {
			continue
		}

		visited[c] = struct{}{}

		for bag, descendants := range b {
			if _, present := descendants[c]; present {
				uniqueAncestors[bag] = struct{}{}
				q.Enqueue(bag)
			}
		}
	}

	ancestors := []string{}

	for ancestor := range uniqueAncestors {
		ancestors = append(ancestors, ancestor)
	}

	return ancestors
}

func (b BagRules) TotalDescendantsOf(ancestor string) int {
	if len(b[ancestor]) == 0 {
		return 0
	}

	var totalDescendants int

	for descendant, count := range b[ancestor] {
		totalDescendants += count + count*b.TotalDescendantsOf(descendant)
	}

	return totalDescendants
}

func ParseBagRules(r io.Reader) BagRules {
	rules := BagRules{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		parent, descendants := parseRule(scanner.Text())
		rules[parent] = descendants
	}

	return rules
}

func parseRule(rawRule string) (string, map[string]int) {
	rawRule = strings.TrimRight(rawRule, ".")
	ruleParts := strings.Split(rawRule, "contain")

	parent := strings.TrimSuffix(ruleParts[0], " bags ")

	descendantParts := strings.Split(ruleParts[1], ",")

	descendants := map[string]int{}

	bagRegex := regexp.MustCompile(`^(\d+)\s([a-z\s]+)\sbags?$`)

	for _, descendant := range descendantParts {
		descendant = strings.TrimSpace(descendant)
		if descendant == "no other bags" {
			break
		}

		matches := bagRegex.FindStringSubmatch(descendant)

		amount, _ := strconv.Atoi(matches[1])
		color := matches[2]

		descendants[color] = amount
	}

	return parent, descendants
}
