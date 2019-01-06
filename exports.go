package expression

import "fmt"

// ParseExpression wraps the entire parse process
func ParseExpression(input string) (Expression, error) {
	tokens, err := tokenize([]rune(input))
	if err != nil {
		return nil, fmt.Errorf("Tokenizer error: %s", err.Error())
	}

	rpnTokens, err := infixToPostfix(tokens)
	if err != nil {
		return nil, fmt.Errorf("RPN conversion error: %s", err.Error())
	}

	return postfixToExpression(rpnTokens)
}
