package advent2020

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type Program []Instruction

type Version int

const (
	_ Version = iota
	Version1
	Version2
)

type Computer struct {
	Version Version
	Mask    string
	Memory  map[int]int
}

func (p Program) Run(c *Computer) {
	if c.Version == 0 {
		c.Version = Version1
	}

	for _, i := range p {
		i.Apply(c)
	}
}

type Op int

const (
	_ Op = iota
	SetMask
	Write
)

type Instruction struct {
	Operation       Op
	Location, Value int
	Mask            string
}

func (i Instruction) Apply(c *Computer) {
	switch i.Operation {
	case SetMask:
		c.Mask = i.Mask
	case Write:
		if c.Memory == nil {
			c.Memory = map[int]int{}
		}

		switch c.Version {
		case Version1:
			c.Memory[i.Location] = applyMask(c.Mask, i.Value)
		case Version2:
			locations := applyMemoryMask(c.Mask, i.Location)

			for _, l := range locations {
				c.Memory[l] = i.Value
			}
		}
	}
}

func ParseInitializationProgram(r io.Reader) Program {
	var instructions []Instruction

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawInstruction := scanner.Text()
		instruction := parseProgramInstruction(rawInstruction)
		instructions = append(instructions, instruction)
	}

	return Program(instructions)
}

func parseProgramInstruction(rawInstruction string) Instruction {
	var instruction Instruction

	if strings.HasPrefix(rawInstruction, "mask =") {
		instruction.Operation = SetMask
		parts := strings.Split(rawInstruction, " = ")
		instruction.Mask = parts[1]
	} else {
		instruction.Operation = Write
		parts := strings.Split(rawInstruction, " = ")

		rawLocation := strings.TrimPrefix(parts[0], "mem[")
		rawLocation = strings.TrimSuffix(rawLocation, "]")

		instruction.Value, _ = strconv.Atoi(parts[1])
		instruction.Location, _ = strconv.Atoi(rawLocation)
	}

	return instruction
}

func applyMask(mask string, value int) int {
	result := value

	for n, v := range mask {
		switch rune(v) {
		case '0':
			result = clearBit(result, 35-n)
		case '1':
			result = setBit(result, 35-n)
		}
	}

	return result
}

func applyMemoryMask(mask string, value int) []int {
	floatingBits := []int{}

	result := value

	for n, v := range mask {
		switch rune(v) {
		case '1':
			result = setBit(result, 35-n)
		case 'X':
			floatingBits = append(floatingBits, 35-n)
		}
	}

	var locations []int

	max := math.Pow(2, float64(len(floatingBits)))
	combinations := generateCombinations(0, int(max))

	for _, combo := range combinations {
		v := result

		for len(combo) < len(floatingBits) {
			combo = "0" + combo
		}

		for location, bit := range combo {
			floater := floatingBits[location]

			switch rune(bit) {
			case '0':
				v = clearBit(v, floater)
			case '1':
				v = setBit(v, floater)
			}
		}

		locations = append(locations, v)
	}

	return locations
}

func setBit(number, bit int) int {
	number |= 1 << bit
	return number
}

func clearBit(number, bit int) int {
	number &= ^(1 << bit)
	return number
}

func generateCombinations(min, max int) []string {
	var combos []string

	for i := min; i < max; i++ {
		combos = append(combos, strconv.FormatUint(uint64(i), 2))
	}

	return combos
}
