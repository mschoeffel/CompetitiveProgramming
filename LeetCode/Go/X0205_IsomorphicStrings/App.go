package X0205_IsomorphicStrings

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "egg"
	t := "add"

	result := isIsomorphic(s, t)
	fmt.Println("Result Isomorphic Strings: " + strconv.FormatBool(result))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	replaceS := make(map[rune]rune)
	replaceT := make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		rs := rune(s[i])
		rt := rune(t[i])
		if r, ok := replaceS[rs]; ok {
			if r != rt {
				return false
			}
		} else if r, ok := replaceT[rt]; ok {
			if r != rs {
				return false
			}
		} else {
			replaceS[rs] = rt
			replaceT[rt] = rs
		}
	}
	return true
}
