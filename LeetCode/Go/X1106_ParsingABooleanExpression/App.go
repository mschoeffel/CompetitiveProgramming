package X1106_ParsingABooleanExpression

import (
	"fmt"
	"strconv"
)

func Main() {
	expression := "&(|(f))"

	result := parseBoolExpr(expression)
	fmt.Println("Result ParsingABooleanExpression: " + strconv.FormatBool(result))
}

func parseBoolExpr(expression string) bool {
	stack := []rune{}

	for _, character := range expression {
		if character != ')' && character != ',' {
			stack = append(stack, character)
		} else if character == ')' {
			currentSubexpression := []bool{}

			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				t := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				currentSubexpression = append(currentSubexpression, t == 't')
			}

			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				t := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				value := currentSubexpression[0]

				if t == '&' {
					for _, b := range currentSubexpression {
						value = value && b
					}
				} else if t == '|' {
					for _, b := range currentSubexpression {
						value = value || b
					}
				} else {
					value = !value
				}

				if value {
					stack = append(stack, 't')
				} else {
					stack = append(stack, 'f')
				}
			}
		}
	}
	return stack[len(stack)-1] == 't'
}
