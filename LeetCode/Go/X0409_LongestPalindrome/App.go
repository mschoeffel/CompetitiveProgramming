package X0409_LongestPalindrome

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "abccccdd"

	result := longestPalindrome(s)
	fmt.Println("Result Longest Palindrome: " + strconv.Itoa(result))

}

func longestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	occur := make(map[rune]int)
	longest := 0
	one := false

	for _, v := range s {
		occur[v]++
	}

	for _, v := range occur {
		if v >= 2 {
			doubles := (v / 2) * 2
			longest += doubles
			v -= doubles
		}
		if v == 1 && !one {
			one = true
			longest += 1
		}
	}

	return longest
}
