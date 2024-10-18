package X2044_CountNumberOfMaximumBitwiseORSubsets

import (
	"fmt"
	"strconv"
)

func Main() {
	example1 := []int{3, 1}

	result := countMaxOrSubsets(example1)
	fmt.Println("Result Count Number of Maximum Bitwise OR Subsets: " + strconv.Itoa(result))
}

func countMaxOrSubsets(nums []int) int {
	maxOR := 0

	for _, num := range nums {
		maxOR |= num
	}

	count := 0
	backtrack(nums, 0, 0, maxOR, &count)

	return count
}

func backtrack(nums []int, index int, currentOR int, maxOR int, count *int) {
	if currentOR == maxOR {
		*count++
	}

	for i := index; i < len(nums); i++ {
		backtrack(nums, i+1, currentOR|nums[i], maxOR, count)
	}
}
