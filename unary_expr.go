package expression

import (
	"fmt"
	"math"
)

type unaryExpression struct {
	term     Expression
	calcFunc func(float64) (float64, error)
	format   string
}

func (exp unaryExpression) Calculate(vals SymbolValues) (float64, error) {
	termValue, err := exp.term.Calculate(vals)
	if err != nil {
		return 0, err
	}
	return exp.calcFunc(termValue)
}

func (exp unaryExpression) RequiresSymbols() []string {
	return exp.term.RequiresSymbols()
}

func (exp unaryExpression) String() string {
	return fmt.Sprintf(exp.format, exp.term.String())
}

// Negative creates an expression negating another expression
func Negative(term Expression) Expression {
	return unaryExpression{
		term:     term,
		calcFunc: func(x float64) (float64, error) { return -x, nil },
		format:   "-%s",
	}
}

// SquareRoot creates an expression calculating the square root of another expression
func SquareRoot(term Expression) Expression {
	return unaryExpression{
		term: term,
		calcFunc: func(x float64) (float64, error) {
			if x < 0 {
				return 0, fmt.Errorf("Sqrt of negative number: %f", x)
			}
			return math.Sqrt(x), nil
		},
		format: "sqrt(%s)",
	}
}

// Sine creates an expression calculating the sine of another expression (using radians)
func Sine(term Expression) Expression {
	return unaryExpression{
		term:     term,
		calcFunc: func(x float64) (float64, error) { return math.Sin(x), nil },
		format:   "sin(%s)",
	}
}

// Cosine creates an expression calculating the cosine of another expression (using radians)
func Cosine(term Expression) Expression {
	return unaryExpression{
		term:     term,
		calcFunc: func(x float64) (float64, error) { return math.Cos(x), nil },
		format:   "cos(%s)",
	}
}

// Tangent creates an expression calculating the tangent of another expression (using radians)
func Tangent(term Expression) Expression {
	return unaryExpression{
		term:     term,
		calcFunc: func(x float64) (float64, error) { return math.Tan(x), nil },
		format:   "tan(%s)",
	}
}
