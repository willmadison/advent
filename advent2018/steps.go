package advent2018

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"
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

	roots, numSteps := findRoots(stepHierarchy)

	visibleSteps := []step{}
	completedSteps := map[step]struct{}{}
	visible := map[step]struct{}{}

	var buf bytes.Buffer

	sort.Slice(roots, func(i, j int) bool {
		return roots[i] < roots[j]
	})

	root := roots[0]

	buf.WriteString(string(root))
	completedSteps[root] = struct{}{}

	nextSteps := stepHierarchy[root]
	nextSteps = append(nextSteps, roots[1:]...)

	sort.Slice(nextSteps, func(i, j int) bool {
		return nextSteps[i] < nextSteps[j]
	})

	visibleSteps = append(visibleSteps, nextSteps...)
	for _, s := range visibleSteps {
		visible[s] = struct{}{}
	}

	var nextStep step

	for len(completedSteps) < numSteps {
		nextStep = visibleSteps[0]
		visibleSteps = visibleSteps[1:]

		prerequisites := nextStep.findPrerequisites(stepHierarchy)

		if nextStep.isCompleteable(completedSteps, prerequisites) {
			buf.WriteString(string(nextStep))
			completedSteps[nextStep] = struct{}{}

			upcomingSteps := stepHierarchy[nextStep]

			for _, s := range upcomingSteps {
				if _, present := visible[s]; !present {
					visibleSteps = append(visibleSteps, s)
					visible[s] = struct{}{}
				}
			}

			sort.Slice(visibleSteps, func(i, j int) bool {
				return visibleSteps[i] < visibleSteps[j]
			})
		} else {
			visibleSteps = append(visibleSteps, nextStep)
		}
	}

	return buf.String()
}

type job struct {
	step          step
	prerequisites []step
}

type completion struct {
	step        step
	timeElapsed int
}

func DetermineInstructionSLA(instructions io.Reader, numHelpers, stepCompletionOverhead int) int {
	workers := numHelpers + 1

	var timeElapsed int

	jobs := make(chan job)
	completions := make(chan completion)

	quit := make(chan struct{})

	for i := 0; i < workers; i++ {
		go work(jobs, completions, stepCompletionOverhead, quit)
	}

	stepHierarchy := map[step][]step{}

	scanner := bufio.NewScanner(instructions)

	for scanner.Scan() {
		step, precedingStep := parseStep(scanner.Text())
		stepHierarchy[precedingStep] = append(stepHierarchy[precedingStep], step)
	}

	roots, numSteps := findRoots(stepHierarchy)

	visibleSteps := []step{}
	completedSteps := map[step]struct{}{}
	visible := map[step]struct{}{}

	sort.Slice(roots, func(i, j int) bool {
		return roots[i] < roots[j]
	})

	visibleSteps = append(visibleSteps, roots...)
	for _, s := range visibleSteps {
		visible[s] = struct{}{}
	}

	go func() {
		for {
			select {
			case completion := <-completions:
				fmt.Printf("completing %v\n", completion)
				completedSteps[completion.step] = struct{}{}
				upcomingSteps := stepHierarchy[completion.step]

				for _, s := range upcomingSteps {
					if _, present := visible[s]; !present {
						visibleSteps = append(visibleSteps, s)
						visible[s] = struct{}{}
					}
				}

				sort.Slice(visibleSteps, func(i, j int) bool {
					return visibleSteps[i] < visibleSteps[j]
				})
			case <-quit:
				return
			}
		}
	}()

	for len(completedSteps) < numSteps {
		if len(visibleSteps) == 0 {
			fmt.Println("waiting for visible steps...")
			time.Sleep(2 * time.Millisecond)
			continue
		}

		nextStep := visibleSteps[0]
		visibleSteps = visibleSteps[1:]

		prerequisites := nextStep.findPrerequisites(stepHierarchy)

		if nextStep.isCompleteable(completedSteps, prerequisites) {
			job := job{nextStep, prerequisites}
			fmt.Printf("sending job: %v\n", job)
			jobs <- job
			fmt.Printf("sent job: %v\n", job)
		} else {
			visibleSteps = append(visibleSteps, nextStep)
		}
	}

	close(quit)

	return timeElapsed
}

func work(jobs <-chan job, completions chan<- completion, overhead int, quit <-chan struct{}) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("receiving job: %v\n", job)
			timeElapsed := int(string(job.step)[0] - 'A' + 1)
			workTime := time.Duration(timeElapsed) * time.Millisecond
			time.Sleep(workTime)
			completion := completion{job.step, timeElapsed + overhead}
			completions <- completion
			fmt.Printf("completed job: %v in %v\n", job, workTime)
		case <-quit:
			return
		}
	}
}

var stepRegEx = regexp.MustCompile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$`)

func parseStep(rawStep string) (step, step) {
	matches := stepRegEx.FindStringSubmatch(rawStep)
	return step(matches[2]), step(matches[1])
}

func findRoots(tree map[step][]step) ([]step, int) {
	seen := map[step]struct{}{}

	for _, children := range tree {
		for _, child := range children {
			seen[child] = struct{}{}
		}
	}

	var roots []step

	for step := range tree {
		if _, present := seen[step]; !present {
			roots = append(roots, step)
		}
	}

	return roots, len(seen) + len(roots)
}
