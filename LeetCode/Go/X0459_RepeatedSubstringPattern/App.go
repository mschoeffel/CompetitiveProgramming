package X0459_RepeatedSubstringPattern

import (
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	s := ""

	result := repeatedSubstringPattern(s)
	fmt.Println("Result Repeated Substring Pattern: " + strconv.FormatBool(result))

}

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 1; i <= length/2; i++ {
		if length%i == 0 && strings.Repeat(s[:i], length/i) == s {
			return true
		}
	}
	return false
}
