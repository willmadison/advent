package advent2018

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
)

type step string

func (s step) isCompleteable(completed map[step]struct{}, prerequisites []step) bool {
	for _, prereq := range prerequisites {
		if _, present := completed[prereq]; !present {
			return false
		}
	}

	return true
}

func (s step) findPrerequisites(tree map[step][]step) []step {
	prerequisites := []step{}

	for step, children := range tree {
		if step == s {
			continue
		}

		for _, child := range children {
			if child == s {
				prerequisites = append(prerequisites, step)
			}
		}
	}

	return prerequisites
}

func DetermineStepOrder(instructions io.Reader) string {
	stepHierarchy := map[step][]step{}

	scanner := bufio.NewScanner(instructions)

	for scanner.Scan() {
		step, precedingStep := parseStep(scanner.Text())
		stepHierarchy[precedingStep] = append(stepHierarchy[precedingStep], step)
	}

	root, _ := findRoot(stepHierarchy)

	fmt.Println("root=", root)
	alphas := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for _, l := range alphas {
		fmt.Println("stepHierarchy[", l, "] =", stepHierarchy[step(l)])
	}

	visibleSteps := []step{}
	completedSteps := map[step]struct{}{}
	visible := map[step]struct{}{}

	var buf bytes.Buffer

	buf.WriteString(string(root))
	completedSteps[root] = struct{}{}

	nextSteps := stepHierarchy[root]
	sort.Slice(nextSteps, func(i, j int) bool {
		return nextSteps[i] < nextSteps[j]
	})

	visibleSteps = append(visibleSteps, nextSteps...)
	for _, s := range visibleSteps {
		visible[s] = struct{}{}
	}

	fmt.Println("visibleSteps=", visibleSteps)

	//var nextStep step

	//	for len(completedSteps) < numSteps {
	//		nextStep = visibleSteps[0]
	//
	//		fmt.Println("nextStep =", nextStep)
	//
	//		visibleSteps = visibleSteps[1:]
	//		fmt.Println("visibleSteps=", visibleSteps)
	//
	//		prerequisites := nextStep.findPrerequisites(stepHierarchy)
	//
	//		if nextStep.isCompleteable(completedSteps, prerequisites) {
	//			buf.WriteString(nextStep.name)
	//			completedSteps[nextStep] = struct{}{}
	//
	//			upcomingSteps := stepHierarchy[nextStep]
	//
	//			for _, s := range upcomingSteps {
	//				if _, present := visible[s]; !present {
	//					visibleSteps = append(visibleSteps, s)
	//				}
	//			}
	//
	//			sort.Slice(visibleSteps, func(i, j int) bool {
	//				return visibleSteps[i].name < visibleSteps[j].name
	//			})
	//		} else {
	//			visibleSteps = append(visibleSteps, nextStep)
	//		}
	//	}

	return buf.String()
}

var stepRegEx = regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)

func parseStep(rawStep string) (step, step) {
	matches := stepRegEx.FindStringSubmatch(rawStep)
	return step(matches[2]), step(matches[1])
}

func findRoot(tree map[step][]step) (step, int) {
	seen := map[step]struct{}{}

	for _, children := range tree {
		for _, child := range children {
			seen[child] = struct{}{}
		}
	}

	var root step

	for step := range tree {
		if _, present := seen[step]; !present {
			root = step
			break
		}
	}

	return root, len(seen) + 1
}
