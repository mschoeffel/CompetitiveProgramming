package X2490_CircularSentence

import (
	"fmt"
	"strconv"
)

func Main() {
	sentence := "leetcode exercises sound delightful"

	result := isCircularSentence(sentence)
	fmt.Println("Result Circular Sentence: " + strconv.FormatBool(result))
}

func isCircularSentence(sentence string) bool {
	if len(sentence) == 1 {
		return true
	}

	runeArray := []rune(sentence)

	if runeArray[0] != runeArray[len(runeArray)-1] {
		return false
	}

	for i := 1; i < len(runeArray)-1; i++ {
		if runeArray[i] == ' ' {
			previousRune := runeArray[i-1]
			nextRune := runeArray[i+1]

			if previousRune != nextRune {
				return false
			}
		}
	}
	return true
}
