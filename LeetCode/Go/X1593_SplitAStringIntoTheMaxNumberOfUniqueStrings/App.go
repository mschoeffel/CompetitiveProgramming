package X1593_SplitAStringIntoTheMaxNumberOfUniqueStrings

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "ababccc"

	result := maxUniqueSplit(s)
	fmt.Println("Result Split a String Into the Max Number of Unique Strings: " + strconv.Itoa(result))
}

func maxUniqueSplit(s string) int {
	if len(s) == 0 {
		return 0
	}

	return maxUniqueSplitHelper(s, make(map[string]bool))
}

func maxUniqueSplitHelper(s string, substrings map[string]bool) int {
	if len(s) == 0 {
		return 0
	}

	result := 0

	for i := 1; i <= len(s); i++ {
		substring := s[:i]

		if _, ok := substrings[substring]; !ok {
			substrings[substring] = true
			result = max(result, 1+maxUniqueSplitHelper(s[i:], substrings))
			delete(substrings, substring)
		}
	}

	return result
}
