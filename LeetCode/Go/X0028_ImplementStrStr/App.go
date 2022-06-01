package X0028_ImplementStrStr

import (
	"fmt"
	"strconv"
)

func Main() {
	haystack := "hello"
	needle := "ll"

	result := strStr(haystack, needle)
	fmt.Println("Result strStr: " + strconv.Itoa(result))

}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	equal := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			equal++
			for j := 1; j < len(needle) && (i+j) < len(haystack); j++ {
				if haystack[i+j] == needle[j] {
					equal++
				}
			}
			if equal == len(needle) {
				return i
			}
			equal = 0
		}
	}
	return -1

}
