package X0485_MaxConsecutiveOnes

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{1, 1, 0, 1, 1, 1}

	result := findMaxConsecutiveOnes(nums)
	fmt.Println("Result Max Consecutive Ones: " + strconv.Itoa(result))
}

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0

	for _, num := range nums {
		if num == 1 {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = 0
		}
	}

	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}
