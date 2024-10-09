package X0005_LongestPalindromicSubstring

import (
	"fmt"
)

func Main() {
	example1 := "babad"

	result := longestPalindrome(example1)
	fmt.Println("Result Longest Palindromic Substring: " + result)
}

func longestPalindrome(s string) string {
	var start int
	var end int

	for i := 0; i < len(s); i++ {
		len1 := searchFromLetter(s, i, i)
		len2 := searchFromLetter(s, i, i+1)
		length := max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func searchFromLetter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
