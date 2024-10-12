package X0520_DetectCapital

import (
	"fmt"
	"strconv"
)

func Main() {
	example1 := "USA"

	result := detectCapitalUse(example1)
	fmt.Println("Result Detect Capital: " + strconv.FormatBool(result))
}

func detectCapitalUse(word string) bool {
	concurrentUpperCase := 0

	for i, char := range word {
		if char >= 'A' && char <= 'Z' {
			if concurrentUpperCase == 0 && i > 0 {
				return false
			}
			concurrentUpperCase += 1
		}
		if char >= 'a' && char <= 'z' {
			if concurrentUpperCase > 1 {
				return false
			}
			concurrentUpperCase = 0
		}
	}
	return true
}
