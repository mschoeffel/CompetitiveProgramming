package X0022_GenerateParentheses

import (
	"encoding/json"
	"fmt"
)

func Main() {
	example1 := 2

	result, _ := json.Marshal(generateParenthesis(example1))
	fmt.Println("Result Generate Parentheses: " + string(result))
}

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, str string, open int, close int, max int) {
	if len(str) == max*2 {
		*result = append(*result, str)
		return
	}
	if open < max {
		backtrack(result, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, str+")", open, close+1, max)
	}
}
