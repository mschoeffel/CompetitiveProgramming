package X020_ValidParentheses

import (
	"fmt"
	"strconv"
)

func Main() {
	parentheses := "({()[{}]})()"

	result := isValid(parentheses)
	fmt.Println("Result Valid Parentheses: " + strconv.FormatBool(result))

}

func isValid(s string) bool {
	var stack []rune

	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		switch c {
		case '(':
			stack = append(stack, ')')
			break
		case '[':
			stack = append(stack, ']')
			break
		case '{':
			stack = append(stack, '}')
			break
		default:
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
