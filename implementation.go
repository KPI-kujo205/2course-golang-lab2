package lab2

import (
	"errors"
	"regexp"
	"strings"
)

// PrefixToPostfix function takes an input string representing a prefix expression and converts it to a postfix expression.
// The function splits the input string into tokens and then iterates through those tokens in reverse order.
// If the token is an operand, it is added to the stack.
// If the token is an operator, it pops two operands from the stack, applies the operator to those operands, and pushes the result back onto the stack.
// At the end of the process, only one element remains on the stack - the result of the postfix expression.
func PrefixToPostfix(input string) (string, error) {
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return "", errors.New("input string contains only whitespace characters")
	}

	var stack []string
	tokens := strings.Fields(input)

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if !isNumber(token) && !isOperator(token[0]) {
			return "", errors.New("invalid character")
		}

		if isOperator(token[0]) {
			if len(stack) < 2 {
				return "", errors.New("not enough operands for operator")
			}
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			postfix := operand1 + " " + operand2 + " " + token
			stack = append(stack, postfix)
		} else {
			stack = append(stack, token)
		}
	}

	return stack[0], nil
}

func isNumber(input string) bool {
	pattern := `^[+-]?\d*\.?\d+$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(input)
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^'
}
