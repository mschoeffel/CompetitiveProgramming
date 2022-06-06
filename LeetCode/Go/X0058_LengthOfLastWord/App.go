package X0058_LengthOfLastWord

import (
	"fmt"
	"strconv"
)

func Main() {
	words := "Hello World"

	result := lengthOfLastWord(words)
	fmt.Println("Result Length of last Word: " + strconv.Itoa(result))

}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
