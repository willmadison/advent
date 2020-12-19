package advent2020

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
)

type Expression []Token

type Token interface {
	Value() string
}

type Operand int

func (o Operand) Value() string {
	return strconv.Itoa(int(o))
}

type Operator string

func (o Operator) Value() string {
	return string(o)
}

func (o Operator) Precedence() int {
	if Operator("(") == o || Operator(")") == o {
		return math.MaxInt64
	}

	if os.Getenv("PART_B") != "" && o == Operator("+") {
		return 2
	}

	return 1
}

func (o Operator) Apply(a, b Operand) int {
	switch o {
	case Operator("+"):
		return int(a) + int(b)
	case Operator("*"):
		return int(a) * int(b)
	}

	return -1
}

type TokenStack interface {
	Push(Token)
	Pop() (Token, error)
	Peek() (Token, error)
	Size() int
}

type TokenQueue interface {
	Enqueue(Token)
	Dequeue() (Token, error)
	Peek() (Token, error)
	Size() int
}

type tokenStack struct {
	data []Token
	size int
}

type tokenQueue struct {
	data []Token
	size int
}

func NewTokenStack() TokenStack {
	return &tokenStack{data: []Token{}}
}

func NewTokenQueue() TokenQueue {
	return &tokenQueue{data: []Token{}}
}

func (t *tokenStack) Push(value Token) {
	t.data = append(t.data, value)
	t.size++
}

func (t *tokenStack) Pop() (Token, error) {
	if t.size > 0 {
		value := t.data[t.size-1]
		t.size--
		t.data = t.data[:t.size]
		return value, nil
	}

	return Operator(""), errors.New("No Such Element")
}

func (t *tokenStack) Peek() (Token, error) {
	if t.size > 0 {
		value := t.data[t.size-1]
		return value, nil
	}

	return Operator(""), errors.New("No Such Element")
}

func (t tokenStack) Size() int {
	return t.size
}

func (t *tokenQueue) Enqueue(value Token) {
	t.data = append(t.data, value)
	t.size++
}

func (t *tokenQueue) Dequeue() (Token, error) {
	if t.size > 0 {
		value := t.data[0]
		t.size--
		t.data = t.data[1:]
		return value, nil
	}

	return Operator(""), errors.New("No Such Element")
}

func (t *tokenQueue) Peek() (Token, error) {
	if t.size > 0 {
		value := t.data[0]
		return value, nil
	}

	return Operator(""), errors.New("No Such Element")
}

func (t tokenQueue) Size() int {
	return t.size
}

func ParseExpressions(r io.Reader) []Expression {
	var expressions []Expression

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		expressions = append(expressions, ParseExpression(scanner.Text()))
	}

	return expressions
}

func ParseExpression(s string) Expression {
	return postfix(extractTokens(s))
}

func extractTokens(s string) []Token {
	var allTokens []Token

	var buf bytes.Buffer

	for _, v := range s {
		if v == ' ' && buf.Len() == 0 {
			continue
		}

		if v == ' ' && buf.Len() > 0 {
			value := buf.String()
			allTokens = append(allTokens, tokenize(value))
			buf.Reset()
			continue
		}

		if v == '(' {
			allTokens = append(allTokens, tokenize(string(v)))
			continue
		}

		if v == ')' {
			if buf.Len() > 0 {
				value := buf.String()
				allTokens = append(allTokens, tokenize(value))
				buf.Reset()
			}
			allTokens = append(allTokens, tokenize(string(v)))
			continue
		}

		buf.WriteRune(v)
	}

	if buf.Len() > 0 {
		value := buf.String()
		allTokens = append(allTokens, tokenize(value))
		buf.Reset()
	}

	return allTokens
}

func postfix(tokens []Token) []Token {
	output := NewTokenQueue()
	operators := NewTokenStack()

	for _, token := range tokens {
		switch t := token.(type) {
		case Operator:
			top, err := operators.Peek()

			if err == nil {
				topOperator := top.(Operator)

				if t.Precedence() <= topOperator.Precedence() && topOperator != Operator("(") {
					output.Enqueue(topOperator)
					operators.Pop()
				}

				if t == Operator(")") {
					for topOperator != Operator("(") {
						operator, _ := operators.Pop()
						output.Enqueue(operator)
						op, err := operators.Peek()

						if err == nil {
							topOperator = op.(Operator)
						} else {
							break
						}
					}

					if topOperator == Operator("(") {
						operators.Pop()
					}

					continue
				}

				operators.Push(t)
			} else {
				operators.Push(t)
			}
		case Operand:
			output.Enqueue(t)
		}
	}

	for operators.Size() > 0 {
		o, _ := operators.Pop()

		if o != Operator("(") {
			output.Enqueue(o)
		}
	}

	var postfixed []Token

	for output.Size() > 0 {
		t, _ := output.Dequeue()
		postfixed = append(postfixed, t)
	}

	return postfixed
}

func tokenize(v string) Token {
	switch v {
	case "+", "*", "(", ")":
		return Operator(v)
	default:
		i, _ := strconv.Atoi(v)
		return Operand(i)
	}
}

func EvaluateExpression(e Expression) int {
	stack := NewTokenStack()

	for _, token := range e {
		switch t := token.(type) {
		case Operator:
			operandA, _ := stack.Pop()
			operandB, _ := stack.Pop()
			a := operandA.(Operand)
			b := operandB.(Operand)

			stack.Push(Operand(t.Apply(a, b)))
		case Operand:
			stack.Push(t)
		}
	}

	result, _ := stack.Pop()
	return int(result.(Operand))
}
