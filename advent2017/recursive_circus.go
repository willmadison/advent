package advent2017

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
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

	var programs []Program

	for _, p := range programsByName {
		programs = append(programs, p)
	}

	sort.Slice(programs, func(i, j int) bool {
		return programs[i].Parent > programs[j].Parent
	})

	for _, p := range programs {
		if len(p.Supporting) > 0 {
			for s := range p.Supporting {
				program := programsByName[s]
				p.Subroutines = append(p.Subroutines, program)
			}
		}

		if p.Parent != "" {
			parent := programsByName[p.Parent]
			parent.Subroutines = append(parent.Subroutines, p)
			programsByName[p.Parent] = parent
		}
	}

	for _, p := range parents {
		program := programsByName[p]

		if program.Parent == "" {
			root = program
		}
	}

	return root
}

func FindImbalance(root Program) int {
	fmt.Println("root:", root)

	weightsBySubtree := map[string]int{}

	for _, t := range root.Subroutines {
		weightsBySubtree[t.Name] = t.TotalWeight()
	}

	fmt.Println("weightsBySubtree:", weightsBySubtree)

	return 0
}
