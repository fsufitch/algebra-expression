package expression

import (
	"errors"
	"fmt"
)

type tokenStack struct {
	arr []token
}

func (s *tokenStack) Push(t token) {
	if s.arr == nil {
		s.arr = []token{}
	}
	s.arr = append(s.arr, t)
}

func (s *tokenStack) Pop() *token {
	if s.arr == nil || len(s.arr) == 0 {
		return nil
	}
	t := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return &t
}

func (s *tokenStack) Len() int {
	return len(s.arr)
}

var operatorPriorities = map[token]int{
	openParens: 10,
	minusUnary: 9,
	sqrt:       9,
	caret:      8,
	log:        8,
	asterisk:   7,
	slash:      7,
	minus:      6,
	plus:       6,
}

func infixToPostfix(tokens []token) ([]token, error) {
	// Thanks to http://csis.pace.edu/~wolf/CS122/infix-postfix.htm
	output := []token{}
	stack := tokenStack{}

	for _, t := range tokens {
		if t.Type == text || t.Type == number {
			output = append(output, t)
			continue
		}

		if t == openParens {
			stack.Push(t)
			continue
		}

		if t == closeParens {
			for true {
				popped := stack.Pop()
				if popped == nil {
					return nil, errors.New("Parsing error: close parens without open parens")
				}
				if *popped == openParens {
					break
				}
				output = append(output, *popped)
			}
			continue
		}

		if _, ok := operatorPriorities[t]; !ok {
			return nil, fmt.Errorf("Parsing error: unrecognized operator %s", t.Text)
		}

		for true {
			stackTop := stack.Pop()
			if stackTop == nil {
				stack.Push(t)
				break
			}
			if *stackTop == openParens {
				stack.Push(*stackTop)
				stack.Push(t)
				break
			}
			if comparePrecedence(t, *stackTop) > 0 {
				stack.Push(*stackTop)
				stack.Push(t)
				break
			}
			if comparePrecedence(t, *stackTop) == 0 {
				output = append(output, *stackTop)
				stack.Push(t)
				break
			}
			output = append(output, *stackTop)
			continue
		}
	}

	for stack.Len() > 0 {
		stackTop := stack.Pop()
		if stackTop == nil {
			return nil, errors.New("Parsing error: nil operator on stack")
		}
		if *stackTop == openParens {
			return nil, errors.New("Parsing error: unclosed open parens")
		}
		output = append(output, *stackTop)
	}

	return output, nil
}

func comparePrecedence(operator1, operator2 token) int {
	return operatorPriorities[operator1] - operatorPriorities[operator2]
}
