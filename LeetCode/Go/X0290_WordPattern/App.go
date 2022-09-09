package X0290_WordPattern

import (
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	pattern := "abba"
	s := "dog cat cat dog"

	output := wordPattern(pattern, s)
	fmt.Println("Result Word Pattern: " + strconv.FormatBool(output))

}

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	m := make(map[rune]string)
	set := make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		if val, ok := m[rune(pattern[i])]; !ok {
			if _, okSet := set[words[i]]; okSet {
				return false
			}
			m[rune(pattern[i])] = words[i]
			set[words[i]] = true
		} else if val != words[i] {
			return false
		}
	}
	return true
}
