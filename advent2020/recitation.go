package advent2020

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Stack interface {
	Push(i int)
	Pop() (int, error)
	Peek() (int, error)
	Size() int
}

type intStack struct {
	data []int
	size int
}

func NewIntStack() Stack {
	return &intStack{data: []int{}}
}

func (s *intStack) Push(value int) {
	s.data = append(s.data, value)
	s.size++
}

func (s *intStack) Pop() (int, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		s.size -= 1
		s.data = s.data[:s.size]
		return value, nil
	}

	return 0, errors.New("No Such Element")
}

func (s *intStack) Peek() (int, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		return value, nil
	}

	return 0, errors.New("No Such Element")
}

func (s intStack) Size() int {
	return s.size
}

func FindNthSpokenNumber(r io.Reader, n int) int {
	turnsBySpokenNumber := map[int][]int{}

	stack := NewIntStack()

	startingNumbers := parseStartingNumbers(r)

	turn := 1

	for _, number := range startingNumbers {
		turnsBySpokenNumber[number] = append(turnsBySpokenNumber[number], turn)
		stack.Push(number)
		turn++
	}

	var spokenNumber int

	for turn <= n {
		lastSpoken, _ := stack.Peek()
		if turnsSpokenOn, previouslySpoken := turnsBySpokenNumber[lastSpoken]; previouslySpoken && len(turnsSpokenOn) == 1 {
			spokenNumber = 0
		} else {
			spokenNumber = turnsSpokenOn[len(turnsSpokenOn)-1] - turnsSpokenOn[len(turnsSpokenOn)-2]
		}

		turnsBySpokenNumber[spokenNumber] = append(turnsBySpokenNumber[spokenNumber], turn)
		stack.Push(spokenNumber)
		turn++
	}

	return spokenNumber
}

func parseStartingNumbers(r io.Reader) []int {
	var numbers []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rawRow := scanner.Text()

		rawNumbers := strings.Split(rawRow, ",")

		for _, raw := range rawNumbers {
			number, _ := strconv.Atoi(raw)
			numbers = append(numbers, number)
		}
	}

	return numbers
}
