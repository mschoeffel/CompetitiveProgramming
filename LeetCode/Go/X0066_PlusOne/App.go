package X0066_PlusOne

import (
	"encoding/json"
	"fmt"
)

func Main() {
	digits := []int{1, 2, 3}

	result := plusOne(digits)
	resultString, _ := json.Marshal(result)
	fmt.Println("Result Plus One: " + string(resultString))

}

func plusOne(digits []int) []int {
	overflow := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if overflow > 0 {
			digits[i] += overflow
			overflow = 0
			if digits[i] >= 10 {
				digits[i] -= 10
				overflow++
			}
		}
	}
	if overflow > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
