package expression

import (
	"fmt"
	"strconv"
	"unicode"
)

type tokenType uint64

const (
	operator tokenType = iota
	parenthesis
	number
	text
)

type token struct {
	Text     string
	Type     tokenType
	HardStop bool
}

var (
	plus        = token{"+", operator, true}
	minus       = token{"-", operator, true}
	minusUnary  = token{"u-", operator, true}
	asterisk    = token{"*", operator, true}
	slash       = token{"/", operator, true}
	caret       = token{"^", operator, true}
	sqrt        = token{"sqrt", operator, false}
	log         = token{"log", operator, false}
	sin         = token{"sin", operator, false}
	cos         = token{"cos", operator, false}
	tan         = token{"tan", operator, false}
	openParens  = token{"(", parenthesis, true}
	closeParens = token{")", parenthesis, true}
)

var reservedStrings = map[string]token{}

func init() {
	reservedTokens := []token{
		plus, minus, asterisk, slash, caret, sqrt, log, openParens, closeParens, sin, cos, tan,
	}
	for _, t := range reservedTokens {
		reservedStrings[t.Text] = t
	}
}

func tokenize(input []rune) ([]token, error) {
	input = append(input, ' ') // Add a space to flush buffer at the end
	tokens := []token{}
	buf := []rune{}
	for i := 0; i < len(input); i++ {
		nextRune := input[i]
		nextBuf := append(buf, nextRune)

		if t, ok := reservedStrings[string(nextBuf)]; ok && t.HardStop {
			// Process "hard stop" tokens
			if t == minus && (len(tokens) == 0 || (len(tokens) > 0 && (tokens[len(tokens)-1].Type == operator || tokens[len(tokens)-1] == openParens))) {
				// Tricky case for unary minus sign
				tokens = append(tokens, minusUnary)
			} else {
				tokens = append(tokens, t)
			}

			buf = []rune{}
			continue
		}

		if _, err := strconv.ParseFloat(string(buf), 64); err == nil {
			// We're building a number
			if _, err := strconv.ParseFloat(string(nextBuf), 64); err == nil {
				// Continues to be a valid number
				buf = nextBuf
			} else {
				// Not valid anymore, save old buf, reprocess current rune
				tokens = append(tokens, token{string(buf), number, false})
				buf = []rune{}
				i--
			}
			continue
		}

		if isValidVariable(buf) {
			if isValidVariable(nextBuf) {
				buf = nextBuf
				continue
			}
			// End of a valid variable
			if t, ok := reservedStrings[string(buf)]; ok {
				tokens = append(tokens, t)
			} else {
				tokens = append(tokens, token{string(buf), text, false})
			}
			buf = []rune{}
			i--
			continue
		}

		if len(buf) == 0 {
			if isValidVariable(nextBuf) {
				// Starting a variable
				buf = nextBuf
				continue
			}
			if _, err := strconv.ParseFloat(string(nextBuf), 64); err == nil {
				// Starting a number
				buf = nextBuf
				continue
			}
			if unicode.IsSpace(nextRune) {
				continue
			}
			return nil, fmt.Errorf("Unexpected rune on empty buffer: %s", string(nextRune))
		}

		return nil, fmt.Errorf("Invalid syntax: %s", string(nextBuf))
	}

	return tokens, nil
}

func isValidVariable(name []rune) bool {
	if len(name) < 1 {
		return false
	}
	if !unicode.IsLetter(name[0]) {
		return false
	}
	for i := 1; i < len(name); i++ {
		if !(unicode.IsLetter(name[i]) || unicode.IsDigit(name[i])) {
			return false
		}
	}
	return true
}
