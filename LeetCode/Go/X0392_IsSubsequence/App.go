package X0392_IsSubsequence

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "abc"
	t := "ahbgdc"

	result := isSubsequence(s, t)
	fmt.Println("Result Is Subsequence: " + strconv.FormatBool(result))

}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	currentIdx := 0
	for _, v := range t {
		if v == rune(s[currentIdx]) {
			currentIdx++
			if currentIdx >= len(s) {
				return true
			}
		}
	}
	return false
}
