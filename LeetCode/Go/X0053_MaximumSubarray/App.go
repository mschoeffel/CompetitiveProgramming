package X0053_MaximumSubarray

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{1, 3, 5, 6}

	result := maxSubArray(nums)
	fmt.Println("Result Maximum Subarray: " + strconv.Itoa(result))

}

func maxSubArray(nums []int) int {
	maxCount := nums[0]
	lastSum := nums[0]

	for i := 1; i < len(nums); i++ {

		if lastSum >= 0 {
			lastSum += nums[i]
		} else {
			lastSum = nums[i]
		}

		if maxCount < lastSum {
			maxCount = lastSum
		}
	}

	return maxCount
}
