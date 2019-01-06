package expression

import "fmt"

// Variable is a string wrapper implementing the
// Expression interface, for evaluating a variable's value
type Variable string

// Calculate finds the variable in the given symbol values
// and returns the value; if not found, returns an error
func (v Variable) Calculate(vals SymbolValues) (float64, error) {
	if value, ok := vals[string(v)]; ok {
		return value, nil
	}
	return 0, fmt.Errorf("no value found for symbol: %s", v)
}

// RequiresSymbols returns a slice containing this variable's symbol
func (v Variable) RequiresSymbols() []string {
	return []string{string(v)}
}

func (v Variable) String() string {
	return string(v)
}

func mergeSymbols(symbols1, symbols2 []string) []string {
	merged := append([]string{}, symbols1...)

	knownSymbols := map[string]bool{}
	for _, s := range symbols1 {
		knownSymbols[s] = true
	}

	for _, s := range symbols2 {
		if _, ok := knownSymbols[s]; !ok {
			knownSymbols[s] = true
			merged = append(merged, s)
		}
	}

	return merged
}
