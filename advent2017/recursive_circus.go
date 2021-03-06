package advent2017

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"fmt"
	"errors"
)

type Program struct {
	Name        string
	Parent      string
	Weight      int
	Supporting  map[string]struct{}
	Subroutines []Program
}

func (p Program) TotalWeight() int {
	var subRoutineWeight int

	for _, r := range p.Subroutines {
		subRoutineWeight += r.TotalWeight()
	}

	return p.Weight + subRoutineWeight
}

func (p Program) String() string {
	return fmt.Sprintf("%s (%d) [Parent: %s]", p.Name, p.Weight, p.Parent)
}

type Queue interface {
	Enqueue(Program)
	Dequeue() (Program, error)
	Peek() (Program, error)
}

type programQueue struct {
	data []Program
	size int
}

func NewProgramQueue() Queue {
	return &programQueue{data: []Program{}}
}

func (p *programQueue) Enqueue(value Program) {
	p.data = append(p.data, value)
	p.size += 1
}

func (p *programQueue) Dequeue() (Program, error) {
	if p.size > 0 {
		value := p.data[0]
		p.size -= 1
		p.data = p.data[1:]

		return value, nil
	}

	return Program{}, errors.New("no such element")
}

func (p *programQueue) Peek() (Program, error) {
	if p.size > 0 {
		value := p.data[0]
		return value, nil
	}

	return Program{}, errors.New("no such element")
}

func (p Program) Traverse(visit func(Program)) {
	queue := NewProgramQueue()

	queue.Enqueue(p)

	var current Program
	var err error

	for err == nil {
		current, err = queue.Dequeue()
		if err != nil {
			break
		}

		visit(current)

		for _, s := range current.Subroutines {
			queue.Enqueue(s)
		}
	}
}

func ParseProgramTowerDataPoint(data string) Program {
	var program Program
	parts := strings.SplitN(data, " -> ", 2)

	program.Name, program.Weight = deriveNameAndWeight(parts[0])

	if len(parts) > 1 {
		program.Supporting = make(map[string]struct{})

		for _, p := range strings.Split(parts[1], ", ") {
			program.Supporting[p] = struct{}{}
		}
	}

	return program
}

func deriveNameAndWeight(nameWeight string) (string, int) {
	parts := strings.Split(nameWeight, " ")

	name := parts[0]
	weight, _ := strconv.Atoi(strings.Trim(parts[1], "()"))

	return name, weight
}

func FindRootOfCallTree(r io.Reader) Program {
	var parents []string

	programsByName := map[string]Program{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		data := scanner.Text()
		p := ParseProgramTowerDataPoint(data)
		programsByName[p.Name] = p

		if len(p.Supporting) > 0 {
			parents = append(parents, p.Name)
		}
	}

	var root Program

	for _, p := range parents {
		parent := programsByName[p]

		for s := range parent.Supporting {
			program := programsByName[s]
			program.Parent = parent.Name
			programsByName[s] = program
		}
	}

	for _, p := range parents {
		program := programsByName[p]

		if program.Parent == "" {
			root = program
		}
	}

	return populateSubroutines(root, programsByName)
}

func populateSubroutines(root Program, programsByName map[string]Program) Program {
	if len(root.Supporting) == 0 {
		return root
	} else {
		for s := range root.Supporting {
			program := programsByName[s]
			root.Subroutines = append(root.Subroutines, populateSubroutines(program, programsByName))
		}

		return root
	}
}

func FindImbalance(root Program) int {
	subtreesByWeight := map[int][]Program{}
	weights := map[int]struct{}{}

	for _, t := range root.Subroutines {
		weight := t.TotalWeight()
		subtreesByWeight[weight] = append(subtreesByWeight[weight], t)
		weights[weight] = struct{}{}
	}

	var uniqueWeights []int

	for weight := range weights {
		uniqueWeights = append(uniqueWeights, weight)
	}

	var commonWeight, outlierWeight int

	var outlier Program

	for weight, subtrees := range subtreesByWeight {
		if len(subtrees) == 1 {
			outlierWeight = weight
			outlier = subtrees[0]
		} else if len(subtrees) > 1 {
			commonWeight = weight
		}
	}

	return doFindImbalance(outlier, outlierWeight-commonWeight)
}

func doFindImbalance(root Program, offset int) int {
	subtreesByWeight := map[int][]Program{}
	weights := map[int]struct{}{}

	for _, t := range root.Subroutines {
		weight := t.TotalWeight()
		subtreesByWeight[weight] = append(subtreesByWeight[weight], t)
		weights[weight] = struct{}{}
	}

	var uniqueWeights []int

	for weight := range weights {
		uniqueWeights = append(uniqueWeights, weight)
	}

	if len(uniqueWeights) == 1 {
		if offset > 0 {
			return root.Weight - offset
		} else {
			return root.Weight + offset
		}
	} else {
		var commonWeight, outlierWeight int

		var outlier Program

		for weight, subtrees := range subtreesByWeight {
			if len(subtrees) == 1 {
				outlierWeight = weight
				outlier = subtrees[0]
			} else if len(subtrees) > 1 {
				commonWeight = weight
			}
		}

		return doFindImbalance(outlier, outlierWeight-commonWeight)
	}
}
