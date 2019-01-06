package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	expression "github.com/fsufitch/algebra-expression"
)

var stdin = bufio.NewReader(os.Stdin)

func getInput(prompt string) string {
	fmt.Print(prompt + " ")
	input, _ := stdin.ReadString('\n')
	return input
}

func main() {
	input := getInput("Input expression:")
	expr, err := expression.ParseExpression(input)

	if err != nil {
		fmt.Println("Error during parsing: ", err)
		return
	}

	fmt.Println("Parsed as: ", expr.String())
	symbolValues := expression.SymbolValues{}

	if len(expr.RequiresSymbols()) > 0 {
		fmt.Println("Values required for evaluation...")
		for _, symbol := range expr.RequiresSymbols() {
			valueStr := strings.TrimSpace(getInput(symbol + " ="))
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				fmt.Println(valueStr)
				fmt.Println("invalid float, exiting")
				return
			}
			symbolValues[symbol] = value
		}
	}

	result, err := expr.Calculate(symbolValues)
	if err != nil {
		fmt.Println("Evaluation error: ", err)
		return
	}
	fmt.Println("Evaluation result: ", result)
	fmt.Println("Success!")
}
