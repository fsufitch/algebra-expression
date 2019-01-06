package expression

import (
	"fmt"
	"strconv"
)

type expressionStack struct {
	arr []Expression
}

func (s *expressionStack) Push(e Expression) {
	if s.arr == nil {
		s.arr = []Expression{}
	}
	s.arr = append(s.arr, e)
}

func (s *expressionStack) Pop() *Expression {
	if s.arr == nil || len(s.arr) == 0 {
		return nil
	}
	t := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return &t
}

func (s *expressionStack) Len() int {
	return len(s.arr)
}
func postfixToExpression(tokens []token) (Expression, error) {
	stack := expressionStack{}
	for _, t := range tokens {
		switch t.Type {
		case number:
			floatValue, err := strconv.ParseFloat(t.Text, 64)
			if err != nil {
				return nil, fmt.Errorf("Builder error: cannot convert to float: %s", t.Text)
			}
			stack.Push(Float64Constant(floatValue))
		case text:
			stack.Push(Variable(t.Text))
		case operator:
			expr, err := operatorToExpression(t, &stack)
			if err != nil {
				return nil, err
			}
			stack.Push(expr)
		default:
			return nil, fmt.Errorf("Builder error: invalid stack element: %s", t.Text)
		}
	}

	if stack.Len() != 1 {
		return nil, fmt.Errorf("Builder error: Invalid final stack length")
	}

	return *stack.Pop(), nil
}

func operatorToExpression(op token, exprStack *expressionStack) (Expression, error) {
	operandRHS := exprStack.Pop()
	operandLHS := exprStack.Pop()
	switch op {
	case sqrt, minusUnary, sin, cos, tan:
		if operandRHS == nil {
			return nil, fmt.Errorf("Builder error: missing operand for unary %s", op.Text)
		}
		if operandLHS != nil {
			exprStack.Push(*operandLHS)
		}
		switch op {
		case sqrt:
			return SquareRoot(*operandRHS), nil
		case minusUnary:
			return Negative(*operandRHS), nil
		case sin:
			return Sine(*operandRHS), nil
		case cos:
			return Cosine(*operandRHS), nil
		case tan:
			return Tangent(*operandRHS), nil
		}
	case plus, minus, asterisk, slash, caret, log:
		if operandRHS == nil {
			return nil, fmt.Errorf("Builder error: missing RHS operand for binary %s", op.Text)
		}
		if operandLHS == nil {
			return nil, fmt.Errorf("Builder error: missing LHS operand for binary %s", op.Text)
		}
		switch op {
		case plus:
			return Addition(*operandLHS, *operandRHS), nil
		case minus:
			return Subtraction(*operandLHS, *operandRHS), nil

		case asterisk:
			return Multiplication(*operandLHS, *operandRHS), nil
		case slash:
			return Division(*operandLHS, *operandRHS), nil
		case caret:
			return Exponentiation(*operandLHS, *operandRHS), nil
		case log:
			return Logarithm(*operandLHS, *operandRHS), nil
		}
	}

	return nil, fmt.Errorf("Builder error: invalid state for operator %s", op.Text)
}
