package expression

import "fmt"

// Float64Constant is a wrapper around float64
// supporting the Expression interface
type Float64Constant float64

// Calculate simply returns the constant's value
func (c Float64Constant) Calculate(v SymbolValues) (float64, error) {
	return float64(c), nil
}

// RequiresSymbols returns an empty string slice;
// no symbols are required to calculate a constant
func (c Float64Constant) RequiresSymbols() []string {
	return []string{}
}

func (c Float64Constant) String() string {
	return fmt.Sprintf("%f", c)
}
