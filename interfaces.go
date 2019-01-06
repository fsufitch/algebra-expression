package expression

// SymbolValues is a map of symbols to mapped values.
// It can be thought of as a variable value map.
type SymbolValues map[string]float64

// Expression is an interface representing anything
// that can be evaluated to a value, be it a number,
// variable, or larger/complex expression.
type Expression interface {
	Calculate(SymbolValues) (float64, error)
	RequiresSymbols() []string
	String() string
}
