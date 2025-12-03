package stacks

import (
	"cmp"
	"errors"
	"fmt"
)

type Stack[C cmp.Ordered] interface {
	Push(value C)
	Pop() (C, error)
	Peek() (C, error)
	Size() int
}

type stack[C cmp.Ordered] struct {
	data []C
	size int
}

func NewStack[C cmp.Ordered]() Stack[C] {
	return &stack[C]{data: []C{}}
}

func (s *stack[C]) Push(value C) {
	s.data = append(s.data, value)
	s.size += 1
}

func (s *stack[C]) Pop() (C, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		s.size -= 1
		s.data = s.data[:s.size]
		return value, nil
	}

	var zero C

	return zero, errors.New("no such element")
}

func (s *stack[C]) Peek() (C, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		return value, nil
	}

	var zero C

	return zero, errors.New("no such element")
}

func (s *stack[C]) String() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *stack[C]) Size() int {
	return s.size
}
