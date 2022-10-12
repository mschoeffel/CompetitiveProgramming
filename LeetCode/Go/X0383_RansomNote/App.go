package X0383_RansomNote

import (
	"fmt"
	"strconv"
)

func Main() {
	note := "a"
	mag := "b"

	result := canConstruct(note, mag)
	fmt.Println("Result Ransom Note: " + strconv.FormatBool(result))

}

func canConstruct(ransomNote string, magazine string) bool {
	occur := make(map[rune]int)

	for i := 0; i < len(magazine); i++ {
		sRune := rune(magazine[i])
		if _, ok := occur[sRune]; !ok {
			occur[sRune] = 0
		}
		occur[sRune]++
	}

	for i := 0; i < len(ransomNote); i++ {
		tRune := rune(ransomNote[i])
		if _, ok := occur[tRune]; !ok {
			return false
		} else {
			occur[tRune]--
			if occur[tRune] < 0 {
				return false
			}
		}
	}
	return true
}
