package X0014_LongestCommonPrefix

import (
	"fmt"
)

func Main() {
	words := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(words)
	fmt.Println("Result Longest Common Prefix: " + result)
	result = longestCommonPrefixSearchByFirst(words)
	fmt.Println("Result Longest Common Prefix SearchByFirst: " + result)

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefixTemp := ""
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] == strs[i][j] {
				prefixTemp += string(prefix[j])
			} else {
				break
			}
		}
		prefix = prefixTemp
	}

	return prefix
}

func longestCommonPrefixSearchByFirst(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return prefix
			}
		}

		prefix += string(strs[0][i])
	}

	return prefix
}
