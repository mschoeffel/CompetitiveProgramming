package X0242_ValidAnagram

import (
	"fmt"
	"strconv"
)

func Main() {
	s := ""
	t := ""

	output := isAnagram(s, t)
	fmt.Println("Result Is Anagram: " + strconv.FormatBool(output))

}

func isAnagram(s string, t string) bool {
	occur := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		sRune := rune(s[i])
		if _, ok := occur[sRune]; !ok {
			occur[sRune] = 0
		}
		occur[sRune]++
	}

	for i := 0; i < len(t); i++ {
		tRune := rune(t[i])
		if _, ok := occur[tRune]; !ok {
			return false
		} else {
			occur[tRune]--
			if occur[tRune] == 0 {
				delete(occur, tRune)
			}
		}
	}
	return len(occur) == 0
}

func isAnagramCompact(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	occur := make([]int, 26, 26)

	for i := range s {
		occur[s[i]-'a']++
		occur[t[i]-'a']--
	}

	for i := range occur {
		if occur[i] != 0 {
			return false
		}
	}
	return true
}
