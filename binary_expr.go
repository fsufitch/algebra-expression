package expression

import (
	"fmt"
	"math"
)

type binaryExpression struct {
	first    Expression
	second   Expression
	calcFunc func(float64, float64) (float64, error)
	format   string
}

func (exp binaryExpression) Calculate(vals SymbolValues) (float64, error) {
	var err error
	firstValue, err := exp.first.Calculate(vals)
	if err != nil {
		return 0, err
	}
	secondValue, err := exp.second.Calculate(vals)
	if err != nil {
		return 0, err
	}
	return exp.calcFunc(firstValue, secondValue)
}

func (exp binaryExpression) RequiresSymbols() []string {
	return mergeSymbols(exp.first.RequiresSymbols(), exp.second.RequiresSymbols())
}

func (exp binaryExpression) String() string {
	return fmt.Sprintf(exp.format, exp.first.String(), exp.second.String())
}

// Addition creates an addition expression with the given expressions as its terms
func Addition(first, second Expression) Expression {
	return binaryExpression{
		first:    first,
		second:   second,
		calcFunc: func(x, y float64) (float64, error) { return x + y, nil },
		format:   "(%s + %s)",
	}
}

// Subtraction creates a subtraction expression with the given expressions as its terms
func Subtraction(first, second Expression) Expression {
	return binaryExpression{
		first:    first,
		second:   second,
		calcFunc: func(x, y float64) (float64, error) { return x - y, nil },
		format:   "(%s - %s)",
	}
}

// Multiplication creates a multiplication expression with the given expressions as its terms
func Multiplication(first, second Expression) Expression {
	return binaryExpression{
		first:    first,
		second:   second,
		calcFunc: func(x, y float64) (float64, error) { return x * y, nil },
		format:   "(%s * %s)",
	}
}

// Division creates a division expression with the given expressions as its terms
func Division(first, second Expression) Expression {
	return binaryExpression{
		first:  first,
		second: second,
		calcFunc: func(x, y float64) (float64, error) {
			if result, ok := safeDivision(x, y); ok {
				return result, nil
			}
			return 0, fmt.Errorf("Divide by zero error: %f / %f", x, y)
		},
		format: "(%s / %s)",
	}
}

func safeDivision(x, y float64) (float64, bool) {
	defer func() {
		recover()
	}()
	return x / y, true
}

// Exponentiation creates an exponentiation expression with the given expressions as its terms
func Exponentiation(first, second Expression) Expression {
	return binaryExpression{
		first:    first,
		second:   second,
		calcFunc: func(x, y float64) (float64, error) { return math.Pow(x, y), nil },
		format:   "(%s ^ %s)",
	}
}

// Logarithm creates a logarithm expression with the given expressions as its terms
func Logarithm(exponent, base Expression) Expression {
	return binaryExpression{
		first:  exponent,
		second: base,
		calcFunc: func(x float64, y float64) (float64, error) {
			if result, ok := safeDivision(math.Log(x), math.Log(y)); ok {
				return result, nil
			}
			return 0, fmt.Errorf("Divide by zero error while calculating: %f log %f", x, y)
		},
		format: "(%s log %s)",
	}
}
