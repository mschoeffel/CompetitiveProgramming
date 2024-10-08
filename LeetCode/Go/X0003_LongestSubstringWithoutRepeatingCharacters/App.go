package X0003_LongestSubstringWithoutRepeatingCharacters

import (
	"fmt"
	"strconv"
)

func Main() {
	example1 := "abcabcbb"

	result := lengthOfLongestSubstring(example1)
	fmt.Println("Result Longest Substring Without Repeating Characters: " + strconv.Itoa(result))
}

func lengthOfLongestSubstring(s string) int {
	var result int
	var start int
	var charMap = make(map[rune]int)

	for i, char := range s {
		if index, ok := charMap[char]; ok {
			if index >= start {
				start = index + 1
			}
		}
		charMap[char] = i
		if i-start+1 > result {
			result = i - start + 1
		}
	}

	return result
}
