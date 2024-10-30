package X1671_MinimumNumberOfRemovalsToMakeMountainArray

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{1, 3, 1}

	result := minimumMountainRemovals(nums)
	fmt.Println("Result Minimum Number of Removals to Make Mountain Array: " + strconv.Itoa(result))
}

func minimumMountainRemovals(nums []int) int {
	lengthNums := len(nums)
	longestIncreasingSequence := make([]int, lengthNums)
	longestDecreasingSequence := make([]int, lengthNums)

	for i := range longestIncreasingSequence {
		longestIncreasingSequence[i] = 1
		longestDecreasingSequence[i] = 1
	}

	for i := 0; i < lengthNums; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				longestIncreasingSequence[i] = max(longestIncreasingSequence[i], longestIncreasingSequence[j]+1)
			}
		}
	}

	for i := lengthNums - 1; i >= 0; i-- {
		for j := lengthNums - 1; j > i; j-- {
			if nums[i] > nums[j] {
				longestDecreasingSequence[i] = max(longestDecreasingSequence[i], longestDecreasingSequence[j]+1)
			}
		}
	}

	maxMountainLength := 0

	for i := 1; i < lengthNums-1; i++ {
		if longestIncreasingSequence[i] > 1 && longestDecreasingSequence[i] > 1 {
			maxMountainLength = max(maxMountainLength, longestIncreasingSequence[i]+longestDecreasingSequence[i]-1)
		}
	}

	return lengthNums - maxMountainLength
}
